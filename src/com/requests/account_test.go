package requests

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"

	"github.com/myndshft/nem-go-sdk/src/model"
)

var testURL = model.DefaultTestnet
var testAddress = "TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S"
var testPublicKey = "0257b05f601ff829fdff84956fb5e3c65470a62375a1cc285779edd5ca3b42f6"
var testTxHash = "161d7f74ab9d332acd46f96650e74371d65b6e1a0f47b076bdd7ccea37903175"
var testTxID = "100"

// RequesterMock is a mock Requester used for testing only!
type RequesterMock struct{}

// Send is a mock Send method for testing only!
func (RequesterMock) Send(s SenderOptions) ([]byte, error) {
	switch s.options.URL.Path {
	case "/account/get/forwarded", "/account/get", "/account/get/from-public-key":
		return json.Marshal(accountMetaDataPairMock)
	case "/account/harvests":
		return json.Marshal(harvestedBlocksArrayMock)
	case "/account/transfers/all", "/account/transfers/incoming", "/account/transfers/outgoing":
		return json.Marshal(transactionMetaDataPairArrayMock)
	case "/account/unconfirmedTransactions":
		return json.Marshal(unconfirmedTransactionMetaDataPairArrayMock)
	case "/account/unlocked/info":
		return json.Marshal(unlockInfoDataMock)
	case "/account/namespace/page":
		return json.Marshal(namespaceMetaDataPairArrayMock)
	case "/account/mosaic/owned":
		return json.Marshal(mosaicArrayMock)
	case "/account/mosaic/owned/definition", "/account/mosaic/definition/page":
		return json.Marshal(mosaicDefinitionArrayMock)
	case "/account/historical/get":
		return json.Marshal(historicalAccountDataArrayMock)
	default:
		return []byte{}, nil
	}
}

var accountMetaDataPairMock = AccountMetaDataPair{
	Account: AccountInfo{
		Address:         "TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S",
		Balance:         124446551689680,
		VestedBalance:   1041345514976241,
		Importance:      0.010263666447108395,
		PublicKey:       "a11a1a6c17a24252e674d151713cdf51991ad101751e4af02a20c61b59f1fe1a",
		Label:           "",
		HarvestedBlocks: 645},
	Meta: AccountMetaData{
		Status:        "LOCKED",
		RemoteStatus:  "ACTIVE",
		CosignatoryOf: []AccountInfo{},
		Cosignatories: []AccountInfo{}}}

func TestForwarded(t *testing.T) {
	want := accountMetaDataPairMock
	got, err := Forwarded(RequesterMock{}, testURL, testAddress)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

func TestData(t *testing.T) {
	want := accountMetaDataPairMock
	got, err := Data(&RequesterMock{}, testURL, testAddress)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

func TestDataFromPublicKey(t *testing.T) {
	want := accountMetaDataPairMock
	got, err := DataFromPublicKey(RequesterMock{}, testURL, testPublicKey)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

var harvestedBlocksArrayMock = harvestInfoRecords{
	Records: []HarvestInfo{
		HarvestInfo{
			TimeStamp:  8963798,
			ID:         254378,
			Difficulty: 46534789865332,
			TotalFee:   2041299054,
			Height:     38453}}}

func TestHarvestedBlocks(t *testing.T) {
	want := harvestedBlocksArrayMock.Records
	got, err := HarvestedBlocks(RequesterMock{}, testURL, testAddress)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

var transactionMetaDataPairArrayMock = transactionMetaDataPairRecords{
	Records: []TransactionMetaDataPair{
		TransactionMetaDataPair{
			Transaction: Transaction{
				TimeStamp:     9111526,
				Signature:     "651a19ccd09c1e0f8b25f6a0aac5825b0a20f158ca4e0d78f2abd904a3966b6e3599a47b9ff199a3a6e1152231116fa4639fec684a56909c22cbf6db66613901",
				Fee:           150000,
				Mode:          1,
				RemoteAccount: "cc6c9485d15b992501e57fe3799487e99de272f79c5442de94eeb998b45e0144",
				Type:          2049,
				Deadline:      9154726,
				Version:       1744830465,
				Signer:        "a1aaca6c17a24252e674d155713cdf55996ad00175be4af02a20c67b59f9fe8a"},
			Meta: TransactionMetaData{
				Height: 1300,
				ID:     1000,
				Hash:   map[string]string{"data": testTxHash}}}}}

func TestAllTransactions(t *testing.T) {
	want := transactionMetaDataPairArrayMock.Records
	got, err := AllTransactions(RequesterMock{}, testURL, testAddress, testTxHash, testTxID)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

func TestIncomingTransactions(t *testing.T) {
	want := transactionMetaDataPairArrayMock.Records
	got, err := IncomingTransactions(RequesterMock{}, testURL, testAddress, testTxHash, testTxID)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

func TestOutgoingTransactions(t *testing.T) {
	want := transactionMetaDataPairArrayMock.Records
	got, err := OutgoingTransactions(RequesterMock{}, testURL, testAddress, testTxHash, testTxID)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

var unconfirmedTransactionMetaDataPairArrayMock = unconfirmedTransactionMetaDataPairRecords{
	Records: []UnconfirmedTransactionMetaDataPair{
		UnconfirmedTransactionMetaDataPair{
			Meta: unconfirmedTransactionMetaData{
				Data: testTxHash},
			Transaction: Transaction{
				TimeStamp:     9111526,
				Signature:     "651a19ccd09c1e0f8b25f6a0aac5825b0a20f158ca4e0d78f2abd904a3966b6e3599a47b9ff199a3a6e1152231116fa4639fec684a56909c22cbf6db66613901",
				Fee:           150000,
				Mode:          1,
				RemoteAccount: "cc6c9485d15b992501e57fe3799487e99de272f79c5442de94eeb998b45e0144",
				Type:          2049,
				Deadline:      9154726,
				Version:       1744830465,
				Signer:        "a1aaca6c17a24252e674d155713cdf55996ad00175be4af02a20c67b59f9fe8a"}}}}

func TestUnconfirmedTransactions(t *testing.T) {
	want := unconfirmedTransactionMetaDataPairArrayMock.Records
	got, err := UnconfirmedTransactions(RequesterMock{}, testURL, testAddress)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

var unlockInfoDataMock = UnlockInfoData{
	NumUnlocked: 5,
	MaxUnlocked: 10}

func TestUnlockInfo(t *testing.T) {
	want := unlockInfoDataMock
	got, err := UnlockInfo(RequesterMock{}, testURL)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

var namespaceMetaDataPairArrayMock = namespaceMetaDataPairRecords{
	Records: []NamespaceMetaDataPair{
		NamespaceMetaDataPair{
			Meta: map[string]int{"id": 100},
			Namespace: NamespaceData{
				FQN:    testTxHash,
				Owner:  testAddress,
				Height: 100}}}}

func TestNamespacesOwned(t *testing.T) {
	want := namespaceMetaDataPairArrayMock.Records
	got, err := NamespacesOwned(RequesterMock{}, testURL, testAddress, testAddress)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

var mosaicArrayMock = mosaicRecords{
	Records: []Mosaic{
		Mosaic{
			MosaicID: MosaicDefinitionID{
				NamespaceID: "testNamespaceIDString",
				Name:        testAddress},
			Quantity: 100}}}

func TestMosaicsOwned(t *testing.T) {
	want := mosaicArrayMock.Records
	got, err := MosaicsOwned(RequesterMock{}, testURL, testAddress)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

var mosaicDefinitionArrayMock = mosaicDefinitionRecords{
	Records: []MosaicDefinition{
		MosaicDefinition{
			Creator: testAddress,
			ID: MosaicDefinitionID{
				NamespaceID: "namespaceid",
				Name:        testTxHash},
			Description: "descriptionhere",
			Properties: []MosaicDefinitionProperties{
				MosaicDefinitionProperties{
					Name:  "namehere",
					Value: "valuehere"}},
			Levy: MosaicDefinitionLevy{
				Type:      100,
				Recipient: testAddress,
				MosaicID: MosaicDefinitionID{
					NamespaceID: "idhere",
					Name:        "namehere"},
				Fee: 100}}}}

func TestAccountMosaicDefinitions(t *testing.T) {
	want := mosaicDefinitionArrayMock.Records
	got, err := AccountMosaicDefinitions(RequesterMock{}, testURL, testAddress)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

func TestMosaicDefinitionsCreated(t *testing.T) {
	want := mosaicDefinitionArrayMock.Records
	got, err := MosaicDefinitionsCreated(RequesterMock{}, testURL, testAddress, testAddress)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

var historicalAccountDataArrayMock = historicalAccountDataRecords{
	Records: []HistoricalAccountData{
		HistoricalAccountData{
			PageRank:        10.45,
			Address:         testAddress,
			Balance:         1034,
			Important:       11.34,
			VestedBalance:   123,
			UnvestedBalance: 911,
			Height:          1234}}}

func TestGetHistoricalAccountData(t *testing.T) {
	want := historicalAccountDataArrayMock.Records
	got, err := GetHistoricalAccountData(RequesterMock{}, testURL, testAddress, 1234)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}
