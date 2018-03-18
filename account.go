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

	"github.com/pkg/errors"
)

// AccountInfo describes basic information for an account.
type AccountInfo struct {
	// Each account has a unique address. First letter of an address
	// indicate the network the account belongs to. Currently two networks
	// are defined: the test network whose account addresses start with a
	// capital T and the main network whose account addresses always start
	// with a capital N. Addresses have always a length of 40 characters
	// and are base-32 encoded.
	Address string
	// Each account has a balance which is an integer greater or equal to
	// zero and denotes the number of micro NEMs which the account owns.
	// Thus a balance of 123456789 means the account owns 123.456789 NEM.
	// A balance is split into its vested and unvested part.
	// Only the vested part is relevant for the importance calculation.
	// For transfers from one account to another only the balance itself
	// is relevant.
	Balance int
	// vestedBalance contains the vested part of the balance of the account
	// in micro NEM.
	VestedBalance float64
	// Each account is assigned an importance. The importance is a decimal
	// number between 0 and 1. It denotes the probability of an account to
	// harvest the next block in case the account has harvesting turned on
	// and all other accounts are harvesting too. The exact formula for
	// calculating the importance is not public yet.
	// Accounts need at least 10k vested NEM to be included
	// in the importance calculation.
	Importance float64
	// The public key of an account can be used to verify signatures of the
	// account. Only accounts that have already published a transaction have
	// a public key assigned to the account. Otherwise the field is null.
	PublicKey string
	// Label has the label of the account (not used, always null).
	Label string
	// Harvesting is the process of generating new blocks. The field
	// denotes the number of blocks that the account harvested so far.
	// For a new account the number is 0.
	HarvestedBlocks int
}

// AccountMetadata describes additional information for the account.
type AccountMetadata struct {
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

// AccountMetadataPair includes durable information for an account and additional information about its state.
type AccountMetadataPair struct {
	// Account contains the account object.
	Account AccountInfo `json:"account"`
	// Meta contain the account meta data object.
	Meta AccountMetadata `json:"meta"`
}

// GetBatchAccountData gets the AccountMetaDataPair of an array of accounts
func (c Client) GetBatchAccountData(addresses []string) ([]AccountMetadataPair, error) {
	data := struct{ Data []AccountMetadataPair }{}
	var pb []map[string]string
	for _, address := range addresses {
		pb = append(pb, map[string]string{"account": address})
	}
	payload, err := json.Marshal(map[string][]map[string]string{"data": pb})
	if err != nil {
		return data.Data, err
	}
	c.url.Path = "/account/batch"
	req, err := c.buildReq(nil, payload, http.MethodPost)
	if err != nil {
		return data.Data, err
	}
	body, err := c.request(req)
	if err != nil {
		return data.Data, err
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return data.Data, err
	}
	return data.Data, nil
}

type XXXXXXXX interface {
	isValid() bool
	String() string
}

type Address string

func (a Address) isValid() bool {
	// check if address is valid
	return true
}

func (a Address) String() string {
	return string(a)
}

type PublicKey string

func (pk PublicKey) isValid() bool {
	// check if address is valid
	return true
}

func (pk PublicKey) String() string {
	return string(pk)
}

// AccountData gets all information for a given address
func (c Client) AccountData(acc XXXXXXXX) (AccountMetadataPair, error) {
	var data AccountMetadataPair
	var req *http.Request
	var err error
	if acc.isValid() {
		switch acc.(type) {
		case Address:
			c.url.Path = "/account/get"
			req, err = c.buildReq(map[string]string{"address": acc.String()}, nil, http.MethodGet)
			if err != nil {
				return data, err
			}
		case PublicKey:
			c.url.Path = "/account/get/from-public-key"
			req, err = c.buildReq(map[string]string{"publicKey": acc.String()}, nil, http.MethodGet)
			if err != nil {
				return data, err
			}
		default:
			return data, errors.New("Use an Address or PublicKey type")
		}
	}
	body, err := c.request(req)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return data, err
	}
	return data, nil
}

// GetDelegated returns the account meta and data info for the account
// for which the given account is the delegate account
func (c Client) GetDelegated(address string) (AccountMetadataPair, error) {
	var data AccountMetadataPair
	c.url.Path = "/account/get/forwarded"
	req, err := c.buildReq(map[string]string{"address": address}, nil, http.MethodGet)
	if err != nil {
		return data, err
	}
	body, err := c.request(req)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return data, err
	}
	return data, nil
}

// AccountStatus gets the current metadata about an account
func (c Client) AccountStatus(address string) (AccountMetadata, error) {
	var data AccountMetadata
	c.url.Path = "/account/status"
	req, err := c.buildReq(map[string]string{"address": address}, nil, http.MethodGet)
	if err != nil {
		return data, err
	}
	body, err := c.request(req)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return data, err
	}
	return data, nil
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
	var data = struct{ Data []HarvestInfo }{}
	c.url.Path = "/account/harvests"
	req, err := c.buildReq(map[string]string{"address": address, "hash": hash}, nil, http.MethodGet)
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
	var data = struct{ Data []OwnedMosaic }{}
	c.url.Path = "/account/mosaic/owned"
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
