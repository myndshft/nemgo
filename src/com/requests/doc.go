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

/*
Package requests implements many of the RESTful API interfaces found in the NEM API.

Create a url.URL object for the NEM testnet or mainnet using default nodes
		testnetUrl := model.DefaultTestnet
		...
		mainnetURL := model.DefaultMainnet


Test that the endpoing is connected, alive, and well
		heartbeat, err := requests.Heartbeat(url)
		if err != nil {
			// handle error
		}
		// heartbeat is a NemRequestResult object
		// ...

Get the current height of the chain
		height, err := requests.Height(url)
		if err != nil {
			// handle error
		}
		// height is a BlockHeight object
		// ...

Get current account data
		accountMetaDataPair, err := requests.Data(url, address)
		if err != nil {
			// handl error
		}
		// accountMetaDataPair is an AccountMetaDataPair object
		// ...

Get namespace information
		namespaceInfo, err := requests.NamespaceInfo(url, "nw")
		if err != nil {
			// handle error
		}
		// namespaceInfo is a NamespaceData object
		// ...

Get Mosaic Definitions of a namespace for sub-namespace
		mosaicInfo, err := requests.MosaicDefinitions(url, "nw")
		if err != nil {
			// handle error
		}
		// mosaicInfo is a MosaicDefinition object
		// ...

Get transaction by hash
		var txHash = "161d7f74ab9d332acd46f96650e74371d65b6e1a0f47b076bdd7ccea37903175"
		txInHash, err := requests.ByHash(url, txHash)
		if err != nil {
			// handle error
		}
		// txInHash is a TransactionMetaDataPair object
		// ...

Get all transactions by account
		txForAccount, err := requests.AllTransactions(url, address, "", "")
		if err != nil {
			// handle error
		}
		// txForAccount is a slice of TransactionMetaDataPair objects
		// ...
*/
package requests
