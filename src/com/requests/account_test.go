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

// RequesterMock is a mock Requester used for testing only!
type RequesterMock struct{}

// Send is a mock Send method for testing only!
func (RequesterMock) Send(s SenderOptions) ([]byte, error) {
	switch s.options.URL.Path {
	case "/account/get/forwarded", "/account/get", "/account/get/from-public-key":
		return json.Marshal(accountMetaDataPairMock)
	default:
		return []byte{}, nil
	}
}
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
