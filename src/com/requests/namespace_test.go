package requests

import (
	"log"
	"reflect"
	"testing"
)

var namespaceMetaDataPairMock = NamespaceMetaDataPair{
	Meta: map[string]int{"id": 100},
	Namespace: NamespaceData{
		FQN:    testTxHash,
		Owner:  testAddress,
		Height: 100}}

func TestRoots(t *testing.T) {
	want := namespaceMetaDataPairMock
	got, err := Roots(RequesterMock{}, testURL, 123)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

var mosaicDefinitionMock = MosaicDefinition{
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
		Fee: 100}}

func TestMosaicDefinitions(t *testing.T) {
	want := mosaicDefinitionMock
	got, err := MosaicDefinitions(RequesterMock{}, testURL, "123")
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}

var namespaceDataMock = NamespaceData{
	FQN:    testTxHash,
	Owner:  testAddress,
	Height: 100}

func TestNamespaceInfo(t *testing.T) {
	want := namespaceDataMock
	got, err := NamespaceInfo(RequesterMock{}, testURL, "123")
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v; want %v", got, want)
	}
}
