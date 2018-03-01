package requests

import (
	"log"
	"reflect"
	"testing"
)

var superNodeDefinitionArrayMock = superNodeDefinitionRecords{
	Records: []SuperNodeDefinition{
		SuperNodeDefinition{
			ID:            "string",
			Alias:         "aliasstring",
			IP:            "ipstring",
			NisPort:       1234,
			PubKey:        "pubkeystring",
			ServantPort:   1234,
			Status:        1,
			Latitude:      45.678,
			Longitude:     123.456,
			PayoutAddress: "payoutaddressstring",
			Distance:      123,
			MaxUnlocked:   3,
			NumUnlocked:   2}}}

func TestNearest(t *testing.T) {
	want := superNodeDefinitionArrayMock.Records
	got, err := Nearest(RequesterMock{}, coords, 1)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

func TestGet(t *testing.T) {
	want := superNodeDefinitionArrayMock.Records
	got, err := Get(RequesterMock{}, 1)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

var superNodeInfoMock = SuperNodeInfo{
	Nodes: []SuperNodeDefinition{
		SuperNodeDefinition{
			ID:            "string",
			Alias:         "aliasstring",
			IP:            "ipstring",
			NisPort:       1234,
			PubKey:        "pubkeystring",
			ServantPort:   1234,
			Status:        1,
			Latitude:      45.678,
			Longitude:     123.456,
			PayoutAddress: "payoutaddressstring",
			Distance:      123,
			MaxUnlocked:   3,
			NumUnlocked:   2}},
	NodeCount: 1}

func TestAll(t *testing.T) {
	want := superNodeInfoMock
	got, err := All(RequesterMock{})
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}
