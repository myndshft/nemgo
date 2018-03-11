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
	"encoding/json"
	"net/http"
)

// TransactionMetadataPair is a set of metadata and transaction details
// about a specific transaction
type TransactionMetadataPair struct {
	Meta        TransactionMetadata
	Transaction Transaction
}

// TransactionMetadata contains metadata about a transaction
type TransactionMetadata struct {
	ID     int
	Height int
	// TODO(tyler): This need custom unmarshal
	Hash hash
}

// Transaction contains information about a transaction
type Transaction struct {
	TimeStamp int
	Amount    int
	Signature string
	Fee       int
	Recipient string
	Type      int
	Deadline  int
	Message   message
	Version   int
	Signer    string
}

type hash struct {
	Data string
}

type message struct {
	Payload string
	Type    int
}

// IncomingTransactions withh list all current pending transactions
// for a given address. This method is likely to be used in conjunction
// with the StreamingUnconfirmedTX method to get additional details about
// the transactions.
func (c Client) IncomingTransactions(address string) ([]TransactionMetadataPair, error) {
	var data struct{ Data []TransactionMetadataPair }
	c.url.Path = "/account/transfers/incoming"
	req, err := c.buildReq(map[string]string{"address": address}, nil, http.MethodGet)
	if err != nil {
		return data.Data, err
	}
	body, err := c.request(req)
	if err != nil {
		return data.Data, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return data.Data, err
	}
	return data.Data, nil
}

// OutgoingTransactions will list all current pending outgoing
// transactions for a given address. This method is likely to be used
// in the conjunction with StreamingUnconfirmedTX method to get additional
// details about the transactions.
func (c Client) OutgoingTransactions(address string) ([]TransactionMetadataPair, error) {
	var data struct{ Data []TransactionMetadataPair }
	c.url.Path = "/account/transfers/outgoing"
	req, err := c.buildReq(map[string]string{"address": address}, nil, http.MethodGet)
	if err != nil {
		return data.Data, err
	}
	body, err := c.request(req)
	if err != nil {
		return data.Data, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return data.Data, err
	}
	return data.Data, nil
}

// AllTransactions will list the most recent transactions either incoming
// or outgoing
func (c Client) AllTransactions(address string) ([]TransactionMetadataPair, error) {
	var data struct{ Data []TransactionMetadataPair }
	c.url.Path = "/account/transfers/all"
	req, err := c.buildReq(map[string]string{"address": address}, nil, http.MethodGet)
	if err != nil {
		return data.Data, err
	}
	body, err := c.request(req)
	if err != nil {
		return data.Data, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return data.Data, err
	}
	return data.Data, nil
}
