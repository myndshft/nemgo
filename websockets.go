// Copyright 2018 Myndshft Technologies, Inc.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nemgo

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/net/websocket"
)

const eol = byte(10)

// TODO create structs for each of the response types to use in Body

// StreamMessage returns the Command, Headers, and Body of a stomp message
type StreamMessage struct {
	Command string
	Headers struct {
		ContentLength string `json:"content-length,omitempty"`
		ContentType   string `json:"content-type,omitempty"`
		Destination   string `json:"destination,omitempty"`
		Receipt       string `json:"receipt,omitempty"`
		Subscription  string `json:"subscription,omitempty"`
		MessageID     string `json:"message-id,omitempty"`
		Version       string `json:"version,omitempty"`
		HeartBeat     string `json:"heart-beat,omitempty"`
	}
	Body interface{}
}

func buildSubscribe(destination string) (string, *big.Int, error) {
	var b strings.Builder
	num, err := rand.Int(rand.Reader, big.NewInt(999999999999))
	if err != nil {
		return b.String(), num, err
	}
	b.WriteString("SUBSCRIBE\r\n")
	fmt.Fprintf(&b, "id:%d\n", num)
	fmt.Fprintf(&b, "destination:%s\n", destination)
	b.WriteString("ack:auto\n")
	b.WriteString("\n\x00")
	return b.String(), num, nil
}

func (c *Client) changeURLPort() {
	split := strings.Split(c.url.Host, ":")
	host, port := split[0], "7778"
	c.url.Host = strings.Join([]string{host, port}, ":")
}

func cleanCommand(line []byte) string {
	command := strings.Trim(string(line), "\r\n\x00")
	return command
}

// stompParser will take a STOMP formatted bytearray and return it in a
// nicely formatted StreamMessage for your enjoyment
// Check out the STOMP protocol here: https://stomp.github.io/index.html
func stompParser(msg []byte) (StreamMessage, error) {
	buf := bytes.NewBuffer(msg)
	msgPart := 0
	var sm StreamMessage
	headers := make(map[string]string)
	for {
		line, err := buf.ReadBytes(eol)
		if err == io.EOF {
			return sm, nil
		}
		if err != nil {
			return sm, errors.Wrap(err, "ReadBytes was not able to understand the message")
		}
		// This means message body is coming next
		if bytes.Equal(line, []byte{10}) {
			msgPart = -1
			continue
		}
		// This means there is no body message and we are done
		if bytes.Equal(line, []byte{0}) {
			return sm, nil
		}
		if msgPart == 0 {
			sm.Command = cleanCommand(line)
			msgPart++
			continue
		}
		if msgPart == -1 {
			switch sm.Command {
			case "SEND", "MESSAGE", "ERROR":
				var b interface{}
				if err = json.Unmarshal(line, &b); err != nil {
					return sm, errors.Wrap(err, "Unable to unmarshal body")
				}
				sm.Body = b
				return sm, nil
			default:
				return sm, nil
			}
		}
		header := strings.Trim(string(line), "\r\n\x00")
		spLine := strings.Split(header, ":")
		headers[spLine[0]] = spLine[1]
		bHeaders, err := json.Marshal(headers)
		if err != nil {
			return sm, err
		}
		if err = json.Unmarshal(bHeaders, &sm.Headers); err != nil {
			return sm, err
		}
	}
}

func (c *Client) stompConnect() (*websocket.Conn, error) {
	c.url.Scheme = "ws"
	c.url.Path = "/w/messages/websocket"
	// websockets use port 7778 not 7890
	c.changeURLPort()
	conn, err := websocket.Dial(c.url.String(), "", "http://localhost")
	if err != nil {
		return nil, err
	}
	host := strings.Split(c.url.Host, ":")[0]
	var b strings.Builder
	b.WriteString("CONNECT\r\n")
	b.WriteString("accept-version:1.2\n")
	fmt.Fprintf(&b, "host:%s\n", host)
	b.WriteString("\n\x00")
	if err = websocket.Message.Send(conn, b.String()); err != nil {
		return nil, err
	}
	var msg []byte
	if err = websocket.Message.Receive(conn, &msg); err != nil {
		return nil, err
	}
	parsedResp, err := stompParser(msg)
	if err != nil {
		return nil, err
	}
	if parsedResp.Command != "CONNECTED" {
		return nil, err
	}
	return conn, nil
}

func subscribe(conn *websocket.Conn, msg string, out chan StreamMessage, subID *big.Int) (chan StreamMessage, error) {
	if err := websocket.Message.Send(conn, msg); err != nil {
		return nil, err
	}
	var e error
	go func() {
		var resp []byte
		for {
			if err := websocket.Message.Receive(conn, &resp); err == io.EOF {
				e = errors.Wrap(err, "The server has no more things to say")
				break
			} else if err != nil {
				e = errors.Wrap(err, "Error occurred while trying to receive message")
			}
			parsedResp, err := stompParser(resp)
			if err != nil {
				e = errors.Wrap(err, "Unable to parse message")
			}
			if parsedResp.Headers.Subscription == subID.String() {
				out <- parsedResp
			}
		}
		close(out)
	}()
	return out, e
}

// SubscribeErrors will return a channel subscribed to error messages
func (c Client) SubscribeErrors() (chan StreamMessage, error) {
	conn, err := c.stompConnect()
	if err != nil {
		return nil, err
	}
	subMsg, subID, err := buildSubscribe("/errors")
	if err != nil {
		return nil, err
	}
	out, err := subscribe(conn, subMsg, make(chan StreamMessage), subID)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscribeHeight will return a channel subscribed to block heights
func (c Client) SubscribeHeight() (chan StreamMessage, error) {
	conn, err := c.stompConnect()
	if err != nil {
		return nil, err
	}
	subMsg, subID, err := buildSubscribe("/blocks/new")
	if err != nil {
		return nil, err
	}
	out, err := subscribe(conn, subMsg, make(chan StreamMessage), subID)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscribeUnconfirmedTX will take an account address and subscribe to all
// unconfirmed transactions at that address
func (c Client) SubscribeUnconfirmedTX(address string) (chan StreamMessage, error) {
	conn, err := c.stompConnect()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/unconfirmed/%s", address)
	subMsg, subID, err := buildSubscribe(path)
	if err != nil {
		return nil, err
	}
	out, err := subscribe(conn, subMsg, make(chan StreamMessage), subID)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscribeConfirmedTX will take an account address and subscribe to all
// confirmed transactions at that address
func (c Client) SubscribeConfirmedTX(address string) (chan StreamMessage, error) {
	conn, err := c.stompConnect()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/transactions/%s", address)
	subMsg, subID, err := buildSubscribe(path)
	if err != nil {
		return nil, err
	}
	out, err := subscribe(conn, subMsg, make(chan StreamMessage), subID)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscribeRecentTX will take an account address and subscribe to all
// recent transactions at that address
func (c Client) SubscribeRecentTX(address string) (chan StreamMessage, error) {
	conn, err := c.stompConnect()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/recenttransactions/%s", address)
	subMsg, subID, err := buildSubscribe(path)
	if err != nil {
		return nil, err
	}
	out, err := subscribe(conn, subMsg, make(chan StreamMessage), subID)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscribeData will take an account address and subscribe to all
// account data changes at that address
func (c Client) SubscribeData(address string) (chan StreamMessage, error) {
	conn, err := c.stompConnect()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/account/%s", address)
	subMsg, subID, err := buildSubscribe(path)
	if err != nil {
		return nil, err
	}
	out, err := subscribe(conn, subMsg, make(chan StreamMessage), subID)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscribeMoasaicData will take an account address and subscribe to all
// mosaic definition changes for that address
func (c Client) SubscribeMoasaicData(address string) (chan StreamMessage, error) {
	conn, err := c.stompConnect()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/account/mosaic/owned/definition/%s", address)
	subMsg, subID, err := buildSubscribe(path)
	if err != nil {
		return nil, err
	}
	out, err := subscribe(conn, subMsg, make(chan StreamMessage), subID)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscribeMosaics will take an account address and subscribe to all
// mosaic changes for that address
func (c Client) SubscribeMosaics(address string) (chan StreamMessage, error) {
	conn, err := c.stompConnect()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/account/mosaic/owned/%s", address)
	subMsg, subID, err := buildSubscribe(path)
	if err != nil {
		return nil, err
	}
	out, err := subscribe(conn, subMsg, make(chan StreamMessage), subID)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscribeNamespaces will take an account address and subscribe to all
// namespace changes for that address
func (c Client) SubscribeNamespaces(address string) (chan StreamMessage, error) {
	conn, err := c.stompConnect()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/account/namespace/owned/%s", address)
	subMsg, subID, err := buildSubscribe(path)
	if err != nil {
		return nil, err
	}
	out, err := subscribe(conn, subMsg, make(chan StreamMessage), subID)
	if err != nil {
		return nil, err
	}
	return out, nil
}
