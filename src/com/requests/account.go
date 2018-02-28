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
	"log"
	"net/http"
	"net/url"
	"strconv"
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
	Account AccountInfo
	// Meta contain the account meta data object.
	Meta AccountMetaData
}

func (amdp *AccountMetaDataPair) unmarshal(data []byte) ([]AccountMetaDataPair, error) {
	type accountMetaDataPairRecords struct {
		Records []AccountMetaDataPair `json:"data"`
	}
	var tmpData accountMetaDataPairRecords
	err := json.Unmarshal(data, &tmpData)
	if err != nil {
		return []AccountMetaDataPair{}, err
	}
	return tmpData.Records, nil
}

// GetBatchAccountData gets the AccountMetaDataPair of an array of accounts
func GetBatchAccountData(u url.URL, addresses []string) ([]AccountMetaDataPair, error) {
	u.Path = "/account/get/batch"
	var build []map[string]string
	for _, address := range addresses {
		addMap := map[string]string{"account": address}
		build = append(build, addMap)
	}
	payload, err := json.Marshal(map[string][]map[string]string{"data": build})
	if err != nil {
		return []AccountMetaDataPair{}, err
	}
	options := Options{
		URL:     u,
		Headers: JSON(payload),
		Method:  http.MethodPost,
		Body:    payload}
	fmt.Println(URLEncoded)
	resp, err := Send(options)
	fmt.Println(string(resp))
	if err != nil {
		return []AccountMetaDataPair{}, err
	}
	var amdp AccountMetaDataPair
	data, err := amdp.unmarshal(resp)
	if err != nil {
		return []AccountMetaDataPair{}, err
	}
	return data, nil
}

// Forwarded gets the AccountMetaDataPair of the account for which the given account is the delegate account
func Forwarded(u url.URL, address string) (AccountMetaDataPair, error) {
	u.Path = "/account/get/forwarded"
	q := u.Query()
	q.Set("address", address)
	u.RawQuery = q.Encode()
	options := Options{
		URL:     u,
		Headers: URLEncoded,
		Method:  http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return AccountMetaDataPair{}, err
	}
	var data AccountMetaDataPair
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return AccountMetaDataPair{}, err
	}
	return data, nil

}

// HarvestingBlocks gets the AccountMetaDataPair of an account.
func HarvestingBlocks(u url.URL, address string) (AccountMetaDataPair, error) {
	u.Path = "/account/get"
	q := u.Query()
	q.Set("address", address)
	u.RawQuery = q.Encode()
	options := Options{
		URL:     u,
		Headers: URLEncoded,
		Method:  http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return AccountMetaDataPair{}, err
	}
	data := AccountMetaDataPair{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return AccountMetaDataPair{}, err
	}
	return data, nil
}

// DataFromPublicKey gets the AccountMetaDataPair of an account with a public key
func DataFromPublicKey(u url.URL, publicKey string) (AccountMetaDataPair, error) {
	u.Path = "/account/get/from-public-key"
	q := u.Query()
	q.Set("publicKey", publicKey)
	u.RawQuery = q.Encode()
	options := Options{
		URL:     u,
		Headers: URLEncoded,
		Method:  http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return AccountMetaDataPair{}, err
	}
	data := AccountMetaDataPair{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return AccountMetaDataPair{}, err
	}
	return data, nil
}

// HarvestInfo contains information about a block that an account harvested.
type HarvestInfo struct {
	// TimeStamp is the number of seconds elapsed since the creation of the nemesis block.
	TimeStamp int `json:"timeStamp"`
	// ID is the database id for the harvested block.
	ID int `json:"id"`
	// Difficulty is the block difficulty.
	// The initial difficulty was set to 100000000000000.
	// The block difficulty is always between one tenth and ten times the initial difficulty.
	Difficulty int `json:"difficulty"`
	// TotalFee is the total fee collected by harvesting the block.
	TotalFee int `json:"totalFee"`
	// Height is the height of the harvested block.
	Height int `json:"height"`
}

func (hi *HarvestInfo) unmarshal(data []byte) ([]HarvestInfo, error) {
	type harvestInfoRecords struct {
		Records []HarvestInfo `json:"data"`
	}
	var tmpData harvestInfoRecords
	err := json.Unmarshal(data, &tmpData)
	if err != nil {
		return []HarvestInfo{}, err
	}
	return tmpData.Records, nil
}

// HarvestedBlocks gets an array of harvest info objects for an account
func HarvestedBlocks(u url.URL, address string) ([]HarvestInfo, error) {
	u.Path = "/account/harvests"
	q := u.Query()
	q.Set("address", address)
	u.RawQuery = q.Encode()
	options := Options{
		URL:     u,
		Headers: URLEncoded,
		Method:  http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return []HarvestInfo{}, err
	}
	var hi HarvestInfo
	data, err := hi.unmarshal(resp)
	if err != nil {
		return []HarvestInfo{}, err
	}
	return data, nil
}

// Transaction represents data of any type of transaction
type Transaction struct {
	// All Transactions
	//
	// TimeStamp is the number of seconds elapsed since the creation of the nemesis block.
	TimeStamp int `json:"timeStamp"`
	// Fee is the fee for the transaction.
	// The higher the fee, the higher the priority of the transaction.
	// Transactions with high priority get included in a block before transactions with lower priority.
	Fee int `json:"fee"`
	// Type is the transaction type.
	Type int `json:"type"`
	// Deadline is the deadline of the transaction.
	// The deadline is given as the number of seconds elapsed since the creation of the nemesis block.
	// If a transaction does not get included in a block before the deadline is reached, it is deleted.
	Deadline int `json:"deadline"`
	// Version is the version of the structure
	Version int `json:"version"`
	// Signer is the public key of the account that created the transaction.
	Signer string `json:"signer"`
	// Signature is the transaction signature (missing if part of a multisig transaction).
	Signature string `json:"signature,omitempty"`
	//
	// Other Transaction Types
	//
	// ImportanceTransferTransaction
	//
	// Mode is the transaction mode.
	// Possible values are:
	// 1: Activate remote harvesting.
	// 2: Deactivate remote harvesting.
	// Mode is found in the following transactions:
	Mode int `json:"mode,omitempty"`
	// RemoteAccount is the public key of the receiving account as hexadecimal string.
	RemoteAccount string `json:"remoteAccount,omitempty"`
	//
	// MosaicDefinitionCreationTransaction
	//
	// CreationFee is the fee for the creation of the mosaic.
	// CreationFee is found in the following Transactions:
	CreationFee int `json:"creationFee,omitempty"`
	// CreationFeeSink is the public key of the account to which the creation fee is tranferred.
	CreationFeeSink string `json:"creationFeeSink,omitempty"`
	// MosaicDefinition is the actual mosaic definition.
	MosaicDefinition MosaicDefinition `json:"mosaicDefinition,omitempty"`
	//
	// MosaicSupplyChangeTransaction
	//
	// SupplyType is the supply type.
	// Supported supply types are:
	// 1: Increase in supply.
	// 2: Decrease in supply.
	SupplyType int `json:"supplyType,omitempty"`
	// Delta is the supply change in units for the mosaic.
	Delta int `json:"delta,omitempty"`
	// MosaicID is the mosaicID object
	MosaicID MosaicDefinitionID `json:"mosaicId,omitempty"`
	//
	// MultisigAggregateModificationTransaction
	//
	// Modifications is a JSON array of multisig modifications.
	Modifications []multisigCosignitaryModification `json:"modifications,omitempty"`
	// MinCosignatories is a JSON object that holds the minimum cosignatories modification.
	MinCosignatories minCosignatoriesDefinition `json:"minCosignatories,omitempty"`
	//
	// MultisigSignatureTransaction
	//
	// OtherHash is the hash of the inner transaction of the corresponding multisig transaction.
	OtherHash singleDataField `json:"otherHash,omitempty"`
	// OtherAccount is the address of the corresponding multisig account.
	OtherAccount string `json:"otherAccount,omitempty"`
	//
	// MultisigTransaction
	//
	// OtherTrans is the inner transaction.
	// The inner transaction can be a transfer transaction, an importance
	// transfer transaction or a multisig aggregate modification transaction.
	// The inner transaction does not have a valid signature.
	OtherTrans *Transaction `json:"otherTrans,omitempty"`
	//
	// MultisigTransaction
	//
	// Signatures is the JSON array of MulsigSignatureTransaction objects.
	Signatures *Transaction `json:"signatures,omitempty"`
	//
	// ProvisionNamespaceTransaction
	//
	// RentalFeeSink is the public key of the account to which the rental fee is transferred.
	RentalFeeSink string `json:"rentalFeeSink,omitempty"`
	// RentalFee is the fee for renting the namespace.
	RentalFee int `json:"rentalFee,omitempty"`
	// NewPart is the new part which is concatenated to the parent with a '.' as separator.
	NewPart string `json:"newPart,omitempty"`
	// Parent is the parent namespace. This can be nil if the transaction rents a root namespace.
	Parent string `json:"parent,omitempty"`
	//
	// TransferTransaction
	//
	// Recipient is the address of the recipient.
	Recipient string `json:"recipient,omitempty"`
	// Message is a message on the transaction.
	Message message `json:"message,omitempty"`
	// Amount is the amount of micro NEM that is transferred from sender to recipient.
	Amount int `json:"amount,omitempty"`
	// Mosaics is an array of Mosaic objects.
	Mosaics []Mosaic `json:"mosaics,omitempty"`
}

// TransactionMetaData contains additional information about the transaction.
type TransactionMetaData struct {
	// Height is the height of the block in which the transaction was included.
	Height int `json:"height"`
	// ID is the id of the transaction.
	ID int `json:"id"`
	// Hash is the transaction hash.
	Hash map[string]string `json:"hash"`
}

// TransactionMetaDataPair contains additional information about the transaction.
type TransactionMetaDataPair struct {
	// Meta contains the transaction meta data object.
	Meta TransactionMetaData `json:"meta"`
	// Transaction contains the transaction object.
	Transaction Transaction `json:"transaction"`
}

func (tmdp *TransactionMetaDataPair) unmarshal(data []byte) ([]TransactionMetaDataPair, error) {
	type transactionMetaDataPairRecords struct {
		Records []TransactionMetaDataPair `json:"data"`
	}
	var tmpData transactionMetaDataPairRecords
	err := json.Unmarshal(data, &tmpData)
	if err != nil {
		return []TransactionMetaDataPair{}, err
	}
	return tmpData.Records, nil
}

// AllTransactions gets all transactions of an account
func AllTransactions(u url.URL, address string, txHash string, txID string) ([]TransactionMetaDataPair, error) {
	u.Path = "/account/transfers/all"
	q := u.Query()
	q.Set("address", address)
	if txHash != "" {
		q.Set("hash", txHash)
	}
	if txID != "" {
		q.Set("id", txID)
	}
	u.RawQuery = q.Encode()
	options := Options{
		URL:     u,
		Headers: URLEncoded,
		Method:  http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return []TransactionMetaDataPair{}, err
	}
	var tmdp TransactionMetaDataPair
	data, err := tmdp.unmarshal(resp)
	if err != nil {
		return []TransactionMetaDataPair{}, err
	}
	return data, nil
}

// IncomingTransactions gets an array of TransactionMetaDataPair objects where
// the sender has the address given as parameter to the request
func IncomingTransactions(u url.URL, address string, txHash string, txID string) ([]TransactionMetaDataPair, error) {
	u.Path = "/account/transfers/incoming"
	q := u.Query()
	q.Set("address", address)
	if txHash != "" {
		q.Set("hash", txHash)
	}
	if txID != "" {
		q.Set("id", txID)
	}
	u.RawQuery = q.Encode()
	options := Options{
		URL:     u,
		Headers: URLEncoded,
		Method:  http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return []TransactionMetaDataPair{}, err
	}
	var tmdp TransactionMetaDataPair
	data, err := tmdp.unmarshal(resp)
	if err != nil {
		return []TransactionMetaDataPair{}, err
	}
	return data, nil
}

// OutgoingTransactions gets an array of TransactionMetaDataPair objects where
// the sender has the address given as parameter to the request.
func OutgoingTransactions(u url.URL, address string, txHash string, txID string) ([]TransactionMetaDataPair, error) {
	u.Path = "/account/transfers/outgoing"
	q := u.Query()
	q.Set("address", address)
	if txHash != "" {
		q.Set("hash", txHash)
	}
	if txID != "" {
		q.Set("id", txID)
	}
	u.RawQuery = q.Encode()
	options := Options{
		URL:     u,
		Headers: URLEncoded,
		Method:  http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return []TransactionMetaDataPair{}, err
	}
	var tmdp TransactionMetaDataPair
	data, err := tmdp.unmarshal(resp)
	if err != nil {
		return []TransactionMetaDataPair{}, err
	}
	return data, nil
}

// TODO add in more complex transaction fields

type multisigCosignitaryModification struct {
	// ModificationType is the type of modification.
	// Possible values are:
	// 1: Add a new cosignatory.
	// 2: Delete an existing cosignatory.
	ModificationType int `json:"modificationType,omitempty"`
	// CosignatoryAccount is the public key of the cosignatory account as hexadecimal string.
	CosignatoryAccount string `json:"cosignatoryAccount,omitempty"`
}

type singleDataField struct {
	Data string `json:"data"`
}

type minCosignatoriesDefinition struct {
	// RelativeChange is a value indicating the relative change of the minimum cosignatories.
	RelativeChange int `json:"relativeChange"`
}

type message struct {
	Payload string `json:"payload"`
	Type    int    `json:"type"`
}

// unconfirmedTransactionMetaData contains the hash of the inner transaction in case the transaction is a multisig transaction.
// This data is need to initiate a multisig signature transaction.
type unconfirmedTransactionMetaData struct {
	// Data is the hash of the inner transaction or null if the transaction is not a multisig transaction.
	Data string `json:"data"`
}

// UnconfirmedTransactionMetaDataPair contains additional information about the transaction
type UnconfirmedTransactionMetaDataPair struct {
	// Meta contains the transaction meta data object.
	Meta unconfirmedTransactionMetaData `json:"meta"`
	// Transaction contains the transaction object.
	Transaction Transaction `json:"transaction"`
}

func (utmdp *UnconfirmedTransactionMetaDataPair) unmarshal(data []byte) ([]UnconfirmedTransactionMetaDataPair, error) {
	type unconfirmedTransactionMetaDataPairRecords struct {
		Records []UnconfirmedTransactionMetaDataPair `json:"data"`
	}
	var tmpData unconfirmedTransactionMetaDataPairRecords
	err := json.Unmarshal(data, &tmpData)
	if err != nil {
		return []UnconfirmedTransactionMetaDataPair{}, err
	}
	return tmpData.Records, nil

}

// UnconfirmedTransactions gets the array of transactions for which an account
// is the sender or receiver and which have not yet been included in a block
func UnconfirmedTransactions(u url.URL, address string) ([]UnconfirmedTransactionMetaDataPair, error) {
	u.Path = "/account/unconfirmedTransactions"
	q := u.Query()
	q.Set("address", address)
	u.RawQuery = q.Encode()
	options := Options{
		URL:     u,
		Headers: URLEncoded,
		Method:  http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return []UnconfirmedTransactionMetaDataPair{}, err
	}
	var utmdp UnconfirmedTransactionMetaDataPair
	data, err := utmdp.unmarshal(resp)
	if err != nil {
		return []UnconfirmedTransactionMetaDataPair{}, err
	}
	return data, nil
}

// UnlockInfoData contains information about the maximum number of allowed
// harvesters, and how many harvesters are already using the node.
// Each node can allow users to harvest with their delegated key on that node.
// The NIS configuration has entries for configuring the maximum number of
// allowed harvesters and optionally allow harvesting only for certain account addresses.
// The unlock info gives information about the maximum number of allowed
// harvesters and how many harvesters are already using the node.
type UnlockInfoData struct {
	// NumUnlocked contains the number of currently unlocked harvesters on the node.
	NumUnlocked int `json:"num-unlocked"`
	// MaxUnlocked contains the number of allowable unlocked harvesters on the node.
	MaxUnlocked int `json:"max-unlocked"`
}

// UnlockInfo gets information about the maximum number of allowed harvesters and how many harvesters are already using the node
func UnlockInfo(u url.URL) (UnlockInfoData, error) {
	u.Path = "/account/unlocked/info"
	options := Options{
		URL:    u,
		Method: http.MethodPost}
	resp, err := Send(options)
	if err != nil {
		return UnlockInfoData{}, err
	}
	var data UnlockInfoData
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return UnlockInfoData{}, err
	}
	return data, nil

}

// NamespaceData consists of information about a namespace
// A namespace is the NEM version of a domain.
// You can rent a namespace for the duration of a year by paying a fee.
// The naming of the parts of a namespace has certain restrictions,
// see the corresponding chapter on namespaces.
type NamespaceData struct {
	// FQN is the fully qualified name of the namespace, also named namespace id.
	FQN string `json:"fqn"`
	// Owner is the owner of the namespace.
	Owner string `json:"owner"`
	// Height is the height at which the the ownership begins.
	Height int `json:"height"`
}

// NamespaceMetaDataPair consists of a namespace object and a database id.
// The id is needed for requests that support paging.
type NamespaceMetaDataPair struct {
	// Meta contains one key: value pair with the key "id" and the value
	// is the database id for the namespace object.
	Meta map[string]int `json:"meta"`
	// Namespace contains the namespace data object.
	Namespace NamespaceData `json:"namespace"`
}

func (nmdp *NamespaceMetaDataPair) unmarshal(data []byte) ([]NamespaceMetaDataPair, error) {
	type namespaceMetaDataPairRecords struct {
		Records []NamespaceMetaDataPair `json:"data"`
	}
	var tmpData namespaceMetaDataPairRecords
	err := json.Unmarshal(data, &tmpData)
	if err != nil {
		return []NamespaceMetaDataPair{}, err
	}
	return tmpData.Records, nil
}

// NamespacesOwned gets namespaces that an account owns
func NamespacesOwned(u url.URL, address string, parent string) ([]NamespaceMetaDataPair, error) {
	u.Path = "/account/namespace/page"
	q := u.Query()
	q.Set("address", address)
	q.Set("parent", parent)
	u.RawQuery = q.Encode()
	options := Options{
		URL:     u,
		Headers: URLEncoded,
		Method:  http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return []NamespaceMetaDataPair{}, err
	}
	var nmdp NamespaceMetaDataPair
	data, err := nmdp.unmarshal(resp)
	if err != nil {
		return []NamespaceMetaDataPair{}, err
	}
	return data, nil
}

// MosaicDefinitionID contains the namespace ID information of the mosaic
type MosaicDefinitionID struct {
	NamespaceID string `json:"namespaceId"`
	Name        string `json:"name"`
}

// Mosaic describes an instance of a mosaic definition.
// Mosaics can be transferred by means of a transfer transaction.
type Mosaic struct {
	// MosaicID is the mosaic id
	MosaicID MosaicDefinitionID `json:"mosaicId"`
	// Quantity is the mosaic quantity.
	// The quantity is always given in smallest units for the mosaic,
	// i.e. if it has a divisibility of 3 the quantity is given in millis.
	Quantity int `json:"quantity"`
}

func (m *Mosaic) unmarshal(data []byte) ([]Mosaic, error) {
	type mosaicRecords struct {
		Records []Mosaic `json:"data"`
	}
	var tmpData mosaicRecords
	err := json.Unmarshal(data, &tmpData)
	if err != nil {
		return []Mosaic{}, err
	}
	return tmpData.Records, nil
}

// MosaicsOwned gets mosaics that an account owns
func MosaicsOwned(u url.URL, address string) ([]Mosaic, error) {
	u.Path = "/account/mosaic/owned"
	q := u.Query()
	q.Set("address", address)
	u.RawQuery = q.Encode()
	options := Options{
		URL:     u,
		Headers: URLEncoded,
		Method:  http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return []Mosaic{}, err
	}
	var m Mosaic
	data, err := m.unmarshal(resp)
	if err != nil {
		return []Mosaic{}, err
	}
	return data, nil
}

// MosaicDefinitionLevy describes the levy on the mosaic.
// A mosaic definition can optionally specify a levy for transferring those mosaics.
// This might be needed by legal entities needing to collect some taxes for transfers.
type MosaicDefinitionLevy struct {
	// Type describes the levy type.
	// The following types are supported:
	// 1: The levy is an absolute fee. The field 'fee' states how many sub-units
	// of the specified mosaic will be transferred to the recipient.
	// 2: The levy is calculated from the transferred amount. The field 'fee'
	// states how many percentiles of the transferred quantity will transferred to the recipient.
	Type int `json:"type"`
	// Recipient is the recipient of the levy.
	Recipient string `json:"recipient"`
	// MosaicID is the mosaic in which the levy is paid.
	MosaicID MosaicDefinitionID `json:"mosaicId"`
	// Fee is the fee. The interpretation is dependent on the type of the levy.
	Fee int `json:"fee"`
}

// MosaicDefinitionProperties allows Unmarshall of data into name/value pairs
// Each mosaic definition comes with a set of properties.
// Each property has a default value which will be applied in case it was not specified.
// Future release may add additional properties to the set of available properties.
// The available properties and their default values are:
// "Divisibility": defines the smallest sub-unit that a mosaic can be divided into.
// A divisibility of 0 means that only entire units can be transferred while a
// divisibility of 3 means the mosaic can be transferred in milli-units.
// "InitialSupply": defines how many units of the mosaic are initially created.
// These mosaics are credited to the creator of the mosaic.
// The initial supply has an upper limit of 9,000,000,000 units.
// "SupplyMutable": determines whether or not the supply can be changed by the
// creator at a later point using a MosaicSupplyChangeTransaction.
// Possible values are "true" and "false", the former meaning the supply can be
// changed and the latter that the supply is fixed for all times.
// "Transferable": determines whether or not the a mosaic can be transferred to
// a user other than the creator. In certain scenarios it is not wanted that
// user are able to trade the mosaic (for example when the mosaic represents bonus
// points which the company does not want to be tranferable to other users).
// Possible values are "true" and "false", the former meaning the mosaic can be
// arbitrarily transferred among users and the latter meaning the mosaic can
// only be transferred to and from the creator.
type MosaicDefinitionProperties struct {
	// Name is the name of the mosaic property.
	Name string `json:"name"`
	// Value is the value of the mosaic property.
	Value string `json:"value"`
}

// MosaicDefinition describes an asset class.
type MosaicDefinition struct {
	// Creator is the public key of the mosaic definition creator.
	Creator string `json:"creator"`
	// ID is the mosaic id.
	ID MosaicDefinitionID `json:"id"`
	// Description is tmosaic description.
	// The description may have a length of up to 512 characters and cannot be empty.
	Description string `json:"description"`
	// Properties are the mosaic properties.
	// The properties may be an empty array in which case default values for
	// all properties are applied.
	Properties []MosaicDefinitionProperties `json:"properties"`
	// Levy is the levy for the mosaic.
	// A creator can demand that each mosaic transfer induces an additional fee.
	Levy MosaicDefinitionLevy `json:"levy"`
}

func (md *MosaicDefinition) unmarshal(data []byte) ([]MosaicDefinition, error) {
	type mosaicDefinitionRecords struct {
		Records []MosaicDefinition `json:"data"`
	}
	var tmpData mosaicDefinitionRecords
	err := json.Unmarshal(data, &tmpData)
	if err != nil {
		return []MosaicDefinition{}, err
	}
	return tmpData.Records, nil
}

// AccountMosaicDefinitions gets mosaic definitions that an account owns
func AccountMosaicDefinitions(u url.URL, address string) ([]MosaicDefinition, error) {
	u.Path = "/account/mosaic/owned/definition"
	q := u.Query()
	q.Set("address", address)
	u.RawQuery = q.Encode()
	options := Options{
		URL:     u,
		Headers: URLEncoded,
		Method:  http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return []MosaicDefinition{}, err
	}
	var md MosaicDefinition
	data, err := md.unmarshal(resp)
	if err != nil {
		return []MosaicDefinition{}, err
	}

	return data, nil
}

// MosaicDefinitionsCreated gets mosaic definitions that an account has created
func MosaicDefinitionsCreated(u url.URL, address string, parent string) ([]MosaicDefinition, error) {
	u.Path = "account/mosaic/definition/page"
	q := u.Query()
	q.Set("address", address)
	q.Set("parent", parent)
	u.RawQuery = q.Encode()
	options := Options{
		URL:     u,
		Headers: URLEncoded,
		Method:  http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return []MosaicDefinition{}, err
	}
	var md MosaicDefinition
	data, err := md.unmarshal(resp)
	if err != nil {
		return []MosaicDefinition{}, err
	}
	return data, nil
}

// HistoricalAccountData has information regarding account data at a specific block.
// This information is historical in nature and is a snapshot of what the account
// information was at a specified block height.
type HistoricalAccountData struct {
	// PageRank is the page rank part of the importance algorithm.
	PageRank float64 `json:"pageRank"`
	// Address is the address of the account.
	Address string `json:"address"`
	// Balance is the balance of the account.
	Balance int `json:"balance"`
	// Importance is the importance score of the account.
	// For more information about the importance score read the blog post here:
	// https://blog.nem.io/what-are-poi-and-vesting/
	Important float64 `json:"importance"`
	// VestedBalance contains the vested part of the balance of the account in micro NEM.
	VestedBalance int `json:"vestedBalance"`
	// UnvestedBalance contains the unvested part of the balance of the account in micro NEM.
	UnvestedBalance int `json:"unvestedBalance"`
	// Height is the height of the blockchain at which the snapshot is taken.
	Height int `json:"height"`
}

func (had *HistoricalAccountData) unmarshal(data []byte) ([]HistoricalAccountData, error) {
	type historicalAccountDataRecords struct {
		Records []HistoricalAccountData `json:"data"`
	}
	var tmpData historicalAccountDataRecords
	err := json.Unmarshal(data, &tmpData)
	if err != nil {
		return []HistoricalAccountData{}, err
	}
	return tmpData.Records, nil
}

// GetHistoricalAccountData gets the HistoricalAccountData of an account from a certain block
func GetHistoricalAccountData(u url.URL, address string, block int) ([]HistoricalAccountData, error) {
	u.Path = "/account/historical/get"
	q := u.Query()
	q.Set("address", address)
	q.Set("startHeight", strconv.Itoa(block))
	q.Set("endHeight", strconv.Itoa(block))
	q.Set("increment", strconv.Itoa(1))
	u.RawQuery = q.Encode()
	options := Options{
		URL:     u,
		Headers: URLEncoded,
		Method:  http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return []HistoricalAccountData{}, err
	}
	fmt.Println(string(resp))
	var amdp HistoricalAccountData
	data, err := amdp.unmarshal(resp)
	if err != nil {
		return []HistoricalAccountData{}, err
	}
	return data, nil
}

type batchHistoricalAccountDataRequest struct {
	Accounts    []batchHistoricalAccountDataRequestAccount `json:"accounts"`
	StartHeight int                                        `json:"startHeight"`
	EndHeight   int                                        `json:"endHeight"`
	IncrementBy int                                        `json:"incrementBy"`
}

type batchHistoricalAccountDataRequestAccount struct {
	Account string `json:"account"`
}

// TODO decide how to unmarshal this nested set of arrays into a struct

// GetBatchHistoricalAccountData gets the AccountMetaDataPair of an array of accounts from an historical height.
func GetBatchHistoricalAccountData(u url.URL, addresses []string, block int) {
	buildPayload := batchHistoricalAccountDataRequest{Accounts: []batchHistoricalAccountDataRequestAccount{}, StartHeight: block, EndHeight: block, IncrementBy: 1}
	for _, address := range addresses {
		buildPayload.Accounts = append(buildPayload.Accounts, batchHistoricalAccountDataRequestAccount{address})
	}
	payload, err := json.Marshal(buildPayload)
	if err != nil {
		log.Fatal(err)
	}
	u.Path = "/account/historical/get/batch"
	options := Options{
		URL:     u,
		Method:  http.MethodPost,
		Headers: JSON(payload),
		Body:    payload}
	resp, err := Send(options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(resp))
}

// StartHarvesting unlocks an account
func StartHarvesting(u url.URL, privateKey string) error {
	u.Path = "/account/unlock"
	payload, err := json.Marshal(map[string]string{"privateKey": privateKey})
	if err != nil {
		return err
	}
	options := Options{
		URL:     u,
		Headers: JSON(payload),
		Method:  http.MethodPost,
		Body:    payload}
	_, err = Send(options)
	if err != nil {
		return err
	}
	return nil
}

// StopHarvesting locks an account
func StopHarvesting(u url.URL, privateKey string) error {
	u.Path = "/account/lock"
	payload, err := json.Marshal(map[string]string{"privateKey": privateKey})
	if err != nil {
		return err
	}
	options := Options{
		URL:    u,
		Method: http.MethodPost,
		Body:   payload}
	_, err = Send(options)
	if err != nil {
		return err
	}
	return nil
}
