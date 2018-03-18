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
	"reflect"
	"testing"
)

func TestRootNamespace(t *testing.T) {
	want := []NamespaceMetadataPair{
		NamespaceMetadataPair{
			NamespaceMetaData: NamespaceMetadata{
				ID: 26264},
			NamespaceInfo: NamespaceInfo{
				FQN:    "makoto.metal.coins",
				Owner:  "TD3RXTHBLK6J3UD2BH2PXSOFLPWZOTR34WCG4HXH",
				Height: 13465}},
		NamespaceMetadataPair{
			NamespaceMetaData: NamespaceMetadata{
				ID: 25421},
			NamespaceInfo: NamespaceInfo{
				FQN:    "gimre.vouchers",
				Owner:  "TDGIMREMR5NSRFUOMPI5OOHLDATCABNPC5ID2SVA",
				Height: 12392}}}
	got, err := clientMock.RootNamespace(1, 5)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(want, got) {
		t.Fatalf("\nWanted: %v\n   Got: %v", want, got)
	}
}

func TestNamespace(t *testing.T) {
	want := NamespaceInfo{
		FQN:    "makoto.metal.coins",
		Owner:  "TD3RXTHBLK6J3UD2BH2PXSOFLPWZOTR34WCG4HXH",
		Height: 13465}
	got, err := clientMock.Namespace("makoto.metal.coins")
	if err != nil {
		t.Fatal(err)
	}
	// BUG(tyler): reflect.DeepEqual shows Namespaces are different in test incorrectly
	if want != got {
		t.Fatalf("\nWanted: %v\n   Got: %v", want, got)
	}
}

// func TestMosaic(t *testing.T) {
// 	want := MosaicInfo{}
// 	got, err := clientMock.Mosaic()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if reflect.DeepEqual(want, got) {
// 		t.Fatalf("\nWanted: %v\n   Got: %v", want, got)
// 	}
// }
