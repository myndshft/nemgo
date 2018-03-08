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
	"fmt"
	"net/http"
)

// AccountInfo describes basic information for an account.
type AccountInfo struct {
	// Address contains the address of the account.
	Address string
	// Balance contains the balance of the account in micro NEM.
	Balance float64
	// vestedBalance contains the vested part of the balance of the account in micro NEM.
	VestedBalance float64
	// Importance contains the importance of the account.
	Importance float64
	// PublicKey contains the public key of the account.
	PublicKey string
	// Label has the label of the account( not used, always null).
	Label string
	// HarvestedBlocks contains the number of blocks that the account already harvested.
	HarvestedBlocks int
}

// AccountMetaData describes additional information for the account.
type AccountMetaData struct {
	// Status contains the harvesting status of a queried account.
	// The harvesting status can be one of the following values:
	// "UNKNOWN": The harvesting status of the account is not known.
	// "LOCKED": The account is not harvesting.
	// "UNLOCKED": The account is harvesting.
	Status string
	// RemoteStatus contains the status of teh remote harvesting of a queried account.
	// The remote harvesting status can be one of the following values:
	// "REMOTE": The account is a remote account and therefore remoteStatus is not applicable for it.
	// "ACTIVATING": The account has activated remote harvesting but it is not yet active.
	// "ACTIVE": The account has activated remote harvesting and remote harvesting is active.
	// "DEACTIVATING": The account has deactivated remote harvesting but remote harvesting is still active.
	// "INACTIVE": The account has inactive remote harvesting, or it has deactivated remote harvesting
	// and deactivation is operational.
	RemoteStatus string
	// CosignatoryOf is a JSON array of AccountInfo structures.
	// The account is cosignatory for each of the accounts in the array.
	CosignatoryOf []AccountInfo
	// Cosignatories is a JSON array of AccountInfo structures.
	// The array holds all accounts that are a cosignatory for this account.
	Cosignatories []AccountInfo
}

// AccountMetaDataPair includes durable information for an account and additional information about its state.
type AccountMetaDataPair struct {
	// Account contains the account object.
	Account AccountInfo `json:"account"`
	// Meta contain the account meta data object.
	Meta AccountMetaData `json:"meta"`
}

// GetBatchAccountData gets the AccountMetaDataPair of an array of accounts
func (c Client) GetBatchAccountData(addresses []string) ([]AccountMetaDataPair, error) {
	var payloadBuilder []map[string]string
	for _, address := range addresses {
		payloadBuilder = append(payloadBuilder, map[string]string{"account": address})
	}
	payload, err := json.Marshal(map[string][]map[string]string{"data": payloadBuilder})
	if err != nil {
		return []AccountMetaDataPair{}, err
	}
	c.URL.Path = "/account/batch"
	req, err := c.buildReq(nil, payload, http.MethodPost)
	if err != nil {
		return []AccountMetaDataPair{}, err
	}
	body, err := c.Request(req)
	if err != nil {
		return []AccountMetaDataPair{}, err
	}
	// The data is returned as a nested json array
	// This enables us to not return the array nested
	// as a value under a "data" key
	data := struct{ Data []AccountMetaDataPair }{}
	if err = json.Unmarshal(body, &data); err != nil {
		return []AccountMetaDataPair{}, err
	}
	return data.Data, nil
}

// AccountInfo gets all information for a given address
func (c Client) AccountInfo(address string) (AccountMetaDataPair, error) {
	c.URL.Path = "/account/get"
	req, err := c.buildReq(map[string]string{"address": address}, nil, http.MethodGet)
	if err != nil {
		return AccountMetaDataPair{}, err
	}
	body, err := c.Request(req)
	if err != nil {
		return AccountMetaDataPair{}, err
	}
	var data AccountMetaDataPair
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return AccountMetaDataPair{}, err
	}
	return data, nil
}

// GetDelegated returns the account meta and data info for the account
// for which the given account is the delegate account
func (c Client) GetDelegated(address string) (AccountMetaDataPair, error) {
	c.URL.Path = "/account/get/forwarded"
	req, err := c.buildReq(map[string]string{"address": address}, nil, http.MethodGet)
	if err != nil {
		return AccountMetaDataPair{}, err
	}
	body, err := c.Request(req)
	if err != nil {
		return AccountMetaDataPair{}, err
	}
	var data AccountMetaDataPair
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return AccountMetaDataPair{}, err
	}
	return data, nil
}

// AccountStatus gets the current metadata about an account
func (c Client) AccountStatus(address string) (AccountMetaData, error) {
	c.URL.Path = "/account/status"
	req, err := c.buildReq(map[string]string{"address": address}, nil, http.MethodGet)
	if err != nil {
		return AccountMetaData{}, err
	}
	body, err := c.Request(req)
	if err != nil {
		return AccountMetaData{}, err
	}
	var data AccountMetaData
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return AccountMetaData{}, err
	}
	return data, nil
}

func (c Client) IncomingTransactions(address string) {
	// TODO
	// /account/transfers/incoming
}

func (c Client) OutgoingTransactions(address string) {
	// TODO
	// /account/transfers/outgoing
}

func (c Client) AccountTransfers(address string) {
	// TODO
	// /account/transfers/all
}

// HarvestInfo is information about harvested blocks
type HarvestInfo struct {
	TimeStamp  int
	Difficulty int
	TotalFee   int
	ID         int
	Height     int
}

// Harvested gets an array of harvest info objects for an account
func (c Client) Harvested(address string, hash string) ([]HarvestInfo, error) {
	c.URL.Path = "/account/harvests"
	req, err := c.buildReq(map[string]string{"address": address, "hash": hash}, nil, http.MethodGet)
	if err != nil {
		return []HarvestInfo{}, err
	}
	body, err := c.Request(req)
	if err != nil {
		return []HarvestInfo{}, err
	}
	var data = struct{ Data []HarvestInfo }{}
	if err := json.Unmarshal(body, &data); err != nil {
		return []HarvestInfo{}, err
	}
	return data.Data, nil
}

// OwnedMosaic is an array of basic information about a mosaic
type OwnedMosaic struct {
	MosaicID struct {
		NamespaceID string
		Name        string
	}
	Quantity int
}

// MosaicsOwned will find information about what mosaics an address
// currently holds
func (c Client) MosaicsOwned(address string) ([]OwnedMosaic, error) {
	c.URL.Path = "/account/mosaic/owned"
	req, err := c.buildReq(map[string]string{"address": address}, nil, http.MethodGet)
	if err != nil {
		return []OwnedMosaic{}, err
	}
	body, err := c.Request(req)
	if err != nil {
		return []OwnedMosaic{}, err
	}
	var data = struct{ Data []OwnedMosaic }{}
	if err := json.Unmarshal(body, &data); err != nil {
		return []OwnedMosaic{}, err
	}
	return data.Data, nil
}
