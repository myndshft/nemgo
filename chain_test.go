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

func TestBlockHeight(t *testing.T) {
	want := 12345
	got, err := clientMock.Height()
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("\nWanted: %v\n   Got: %v", want, got)
	}

}

func TestScore(t *testing.T) {
	want := "18722d5a7d590deb"
	got, err := clientMock.Score()
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("\nWanted: %v\n   Got: %v", want, got)
	}

}

func TestLastBlock(t *testing.T) {
	want := Block{
		TimeStamp: 9232968,
		Signature: "0a1351ef3e9b19c601e804a6d329c9ade662051d1da2c12c3aec9934353e421c79de7d8e59b127a8ca9b9d764e3ca67daefcf1952f71bc36f747c8a738036b05",
		// TODO(tyler): Fix this to be a simple string
		PrevBlockHash: "58efa578aea719b644e8d7c731852bb26d8505257e03a897c8102e8c894a99d6",
		Type:          1,
		Transactions:  []Transaction{},
		Version:       1744830465,
		Signer:        "2afca04d2cb8d16cf3656274bc55b95e60be823cfb7230d82f791ed42a309ee7",
		Height:        42804}
	got, err := clientMock.BlockInfo(42804)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nWanted: %v\n   Got: %v", want, got)
	}
}

func TestBlockInfo(t *testing.T) {
	want := Block{
		TimeStamp: 9232968,
		Signature: "0a1351ef3e9b19c601e804a6d329c9ade662051d1da2c12c3aec9934353e421c79de7d8e59b127a8ca9b9d764e3ca67daefcf1952f71bc36f747c8a738036b05",
		// TODO(tyler): Fix this to be a simple string
		PrevBlockHash: "58efa578aea719b644e8d7c731852bb26d8505257e03a897c8102e8c894a99d6",
		Type:          1,
		Transactions:  []Transaction{},
		Version:       1744830465,
		Signer:        "2afca04d2cb8d16cf3656274bc55b95e60be823cfb7230d82f791ed42a309ee7",
		Height:        42804}
	got, err := clientMock.LastBlock()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nWanted: %v\n   Got: %v", want, got)
	}

}
