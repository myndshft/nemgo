package requests

import (
	"log"
	"reflect"
	"testing"
)

var nemRequestResultMock = NemRequestResult{
	Type:    1,
	Code:    123,
	Message: "messagehere",
	TransactionHash: nemRequestResultData{
		Data: "datahere"},
	InnerTransactionHash: nemRequestResultData{
		Data: "datahere"}}

func TestHeartbeat(t *testing.T) {
	want := nemRequestResultMock
	got, err := Heartbeat(RequesterMock{}, testURL)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}
