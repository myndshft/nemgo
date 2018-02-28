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

// NemRequestResult is typically used for requests that perform validation
// or return a status.
// Some requests such as announcing a new transaction return detailed
// information about the outcome of the request.
// In those cases the result of the request is returned in a special
// JSON object called NemRequestResult.
type NemRequestResult struct {
	// Type is dependent on the request which was answered.
	// The interpretation of the code field depends on the type.
	// Currently the following types are supported:
	// 	1: The result is a validation result.
	// 	2: The result is a heart beat result.
	// 	4: The result indicates a status.
	Type int `json:"type"`
	// Code meaning is dependent on the Type.
	//
	// For type 1 (validation result) only 0 and 1 mean there was no failure.
	// The following codes are the most frequent ones occurring:
	// 0:Neutral result. A typical example would be that a node validates
	// an incoming transaction and realizes that it already knows about the transaction.
	// In this case it is neither a success (meaning the node has a new transaction)
	// nor a failure (because the transaction itself is valid).
	// 1:Success result. A typical example would be that a node validates a new valid transaction.
	// 2:Unknown failure. The validation failed for unknown reasons.
	// 3:The entity that was validated has already past its deadline.
	// 4:The entity used a deadline which lies too far in the future.
	// 5:There was an account involved which had an insufficient balance to perform the operation.
	// 6:The message supplied with the transaction is too large.
	// 7:The hash of the entity which got validated is already in the database.
	// 8:The signature of the entity could not be validated.
	// 9:The entity used a timestamp that lies too far in the past.
	// 10:The entity used a timestamp that lies in the future which is not acceptable.
	// 11:The entity is unusable.
	// 12:The score of the remote block chain is inferior (although a superior score was promised).
	// 13:The remote block chain failed validation.
	// 14:There was a conflicting importance transfer detected.
	// 15:There were too many transaction in the supplied block.
	// 16:The block contains a transaction that was signed by the harvester.
	// 17:A previous importance transaction conflicts with a new transaction.
	// 18:An importance transfer activation was attempted while previous one is active.
	// 19:An importance transfer deactivation was attempted but is not active.

	// For type 2 the following codes are supported:
	// 1:Successful heart beat detected.

	// For type 3 the following codes are supported:
	// 0:Unknown status.
	// 1:NIS is stopped.
	// 2:NIS is starting.
	// 3:NIS is running.
	// 4:NIS is booting the local node (implies NIS is running).
	// 5:The local node is booted (implies NIS is running).
	// 6:The local node is synchronized (implies NIS is running and the local node is booted).
	// 7:There is no remote node available (implies NIS is running and the local node is booted).
	// 8:NIS is currently loading the block chain.
	Code int `json:"code"`
	// Message contains information about what the purpose of the request was.
	Message string `json:"message"`
	// TransactionHash is the JSON hash object of the transaction
	TransactionHash nemRequestResultData `json:"transactionHash,omitempty"`
	// InnerTransactionHash is the JSON hash object of the inner transaction or
	// null if the transaction is not a multisig transaction.
	InnerTransactionHash nemRequestResultData `json:"innerTransactionHash,omitempty"`
}

type nemRequestResultData struct {
	Data string `json:"data"`
}

// Heartbeat determines if NIS is up and responsive
func Heartbeat(u url.URL) (NemRequestResult, error) {
	u.Path = "/heartbeat"
	options := Options{
		URL:    u,
		Method: http.MethodGet}
	resp, err := Send(options)
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
