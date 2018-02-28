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

package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/myndshft/nem-go-sdk/src/model"
)

type MarketInfo struct {
	something int
}

// TODO determine the best way to handle this -- interface?

// Xem gets market information from Poloniex API
func Xem(sender Sender) (MarketInfo, error) {
	u := model.MarketInfo
	q := u.Query()
	q.Set("command", "returnTicker")
	u.RawQuery = q.Encode()
	options := Options{
		URL:    u,
		Method: http.MethodGet}
	senderOpts := NewDefaultSenderOptions(options)
	resp, err := sender.Send(senderOpts)
	if err != nil {
		return MarketInfo{}, err
	}
	fmt.Println(string(resp))
	var data MarketInfo
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return MarketInfo{}, err
	}
	return data, nil

}

// TODO determine the best way to handle this -- interface?

// BTC gets the BTC price from blockchain.info API
func BTC(sender Sender) (MarketInfo, error) {
	u := model.BTCPrice
	q := u.Query()
	q.Set("command", "returnTicker")
	u.RawQuery = q.Encode()
	options := Options{
		URL:    u,
		Method: http.MethodGet}
	senderOpts := NewDefaultSenderOptions(options)
	resp, err := sender.Send(senderOpts)
	if err != nil {
		return MarketInfo{}, err
	}
	fmt.Println(string(resp))
	var data MarketInfo
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return MarketInfo{}, err
	}
	return data, nil
}
