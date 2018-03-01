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
	"net/http"
	"net/url"
)

// RequestAnnounce is used to transfer the transaction data and the signature
// to NIS in order to initiate and broadcast a transaction.
type RequestAnnounce struct {
	// Data is the transaction data as string.
	// The string is created by first creating the corresponding byte array
	// and then converting the byte array to a hexadecimal string.
	Data string `json:"data"`
	// Signature is the signature for the transaction as a hexadecimal string.
	Signature string `json:"signature"`
}

// Announce will broadcast a transaction on the NEM network.
func Announce(sender Sender, u url.URL, reqAnn RequestAnnounce) (NemRequestResult, error) {
	u.Path = "/transaction/announce"
	payload, err := json.Marshal(reqAnn)
	if err != nil {
		return NemRequestResult{}, err
	}
	options := Options{
		URL:     u,
		Method:  http.MethodPost,
		Headers: JSON(payload),
		Body:    payload}
	senderOpts := NewDefaultSenderOptions(options)
	resp, err := sender.Send(senderOpts)
	if err != nil {
		return NemRequestResult{}, err
	}
	var data NemRequestResult
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return NemRequestResult{}, err
	}
	return data, nil

}

// ByHash gets a TransactionMetaDataPair object from the chain using it's hash.
func ByHash(sender Sender, u url.URL, txHash string) (TransactionMetaDataPair, error) {
	u.Path = "/transaction/get"
	q := u.Query()
	q.Set("hash", txHash)
	u.RawQuery = q.Encode()
	options := Options{
		URL:     u,
		Method:  http.MethodGet,
		Headers: URLEncoded}
	senderOpts := NewDefaultSenderOptions(options)
	resp, err := sender.Send(senderOpts)
	if err != nil {
		return TransactionMetaDataPair{}, err
	}
	var tmdp TransactionMetaDataPair
	err = json.Unmarshal(resp, &tmdp)
	if err != nil {
		return TransactionMetaDataPair{}, err
	}
	return tmdp, nil
}
