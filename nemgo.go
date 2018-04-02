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
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	// Testnet is a helper const to connect to the testnet
	Testnet = byte(0x98)
	// Mainnet is a helper const to connect to the mainnet
	Mainnet = byte(0x68)
)

// Client is used to interact with a NIS
type Client struct {
	network byte
	url     url.URL
	request func(*http.Request) ([]byte, error)
}

// Option can be passed into New() to enable additional configuration
// of the returned Client
type Option func(*Client)

// TODO(tyler): Implement a logger
// TODO(tyler): Create additional configuration options

// WithNIS allows the user to create a Client connected to a NIS
// of their choosing
// The network param must be a single byte representing the network
// which the host NIS is a member of
// mainnet = byte(0x68)
// testnet = byte(0x98)
func WithNIS(host string, network Network) Option {
	return func(c *Client) {
		c.url = url.URL{Scheme: "http", Host: host}
		c.network = network
	}
}

// New will return a Client object ready to be used
// defaulting to the NEM mainnet
func New(opts ...Option) *Client {
	c := &Client{
		network: Mainnet,
		url:     url.URL{Scheme: "http", Host: "209.126.98.204:7890"},
		request: sendReq}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c Client) buildReq(params map[string]string, payload []byte, method string) (*http.Request, error) {
	if params != nil {
		q := c.url.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		c.url.RawQuery = q.Encode()
	}
	req, err := http.NewRequest(method, c.url.String(), bytes.NewBuffer(payload))
	if err != nil {
		return req, err
	}
	return req, nil
}

func sendReq(req *http.Request) ([]byte, error) {
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
