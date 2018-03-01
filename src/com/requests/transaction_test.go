package requests

import (
	"log"
	"reflect"
	"testing"
)

func TestAnnounce(t *testing.T) {
	want := nemRequestResultMock
	got, err := Announce(RequesterMock{}, testURL, RequestAnnounce{})
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

var transactionMetaDataPairMock = TransactionMetaDataPair{
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
		Hash:   map[string]string{"data": testTxHash}}}

func TestByHash(t *testing.T) {
	want := transactionMetaDataPairMock
	got, err := ByHash(RequesterMock{}, testURL, testTxHash)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}
