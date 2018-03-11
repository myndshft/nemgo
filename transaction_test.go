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

func TestIncomingTransactions(t *testing.T) {
	want := []TransactionMetadataPair{
		TransactionMetadataPair{
			TransactionMetadata{
				ID:     71245,
				Height: 40706,
				Hash: hash{
					Data: "15c373ad4c3fe6af47d1941379ff262f785bdcfa07c02ac3608bc10da27d5e82",
				},
			},
			Transaction{
				TimeStamp: 9106400,
				Amount:    1000000000,
				Signature: "449cd76ea8bda2220b3d6ad6f8db5f81d4e68ad3d4b0c3db9a3c267355657639eabed3dbcef8e0cc22953ae2b36a22ee7dc6327484c9649cccd686a511eca105",
				Fee:       3000000,
				Recipient: "TALICELCD3XPH4FFI5STGGNSNSWPOTG5E4DS2TOS",
				Type:      257,
				Deadline:  9149600,
				Message: message{
					Payload: "280000005444334b32493543524850595634425a5a5a4c335850454e4",
					Type:    2,
				},
				Version: -1744830463,
				Signer:  "c20a1dffe699c7a68328986273265e33fceebe074f274240ef890dd80ad55ed6",
			},
		},
		TransactionMetadataPair{
			TransactionMetadata{
				ID:     71356,
				Height: 40629,
				Hash: hash{
					Data: "37c34ead4c3fe6af42d994135798262f785ba2d807c02ac3608bc10da12e5f87",
				},
			},
			Transaction{
				TimeStamp: 9101541,
				Amount:    49997995000000,
				Signature: "57c3c48d2ae8b24240b57d72493f498cfeb61e2ab87237dc0e08c51007d5c7f15847d0e08c0286e68a72028925db5fa809ca9d57e2cb6eebe11822176a834c0b",
				Fee:       2005000000,
				Recipient: "TALICELCD3XPH4FFI5STGGNSNSWPOTG5E4DS2TOS",
				Type:      257,
				Deadline:  9144741,
				Message: message{
					Payload: "526f6262657279212121",
					Type:    1,
				},
				Version: -1744830463,
				Signer:  "546e4fb9c81db84e04d8e9e67380db0fe1f540df09a527fb995b589b5695ae24",
			},
		},
	}
	got, err := clientMock.IncomingTransactions("TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\nWanted: %v\n Got: %v", want, got)
	}
}

func TestOutgoingTransactions(t *testing.T) {
	want := []TransactionMetadataPair{
		TransactionMetadataPair{
			TransactionMetadata{
				ID:     71245,
				Height: 40706,
				Hash: hash{
					Data: "15c373ad4c3fe6af47d1941379ff262f785bdcfa07c02ac3608bc10da27d5e82",
				},
			},
			Transaction{
				TimeStamp: 9106400,
				Amount:    1000000000,
				Signature: "449cd76ea8bda2220b3d6ad6f8db5f81d4e68ad3d4b0c3db9a3c267355657639eabed3dbcef8e0cc22953ae2b36a22ee7dc6327484c9649cccd686a511eca105",
				Fee:       3000000,
				Recipient: "TALICELCD3XPH4FFI5STGGNSNSWPOTG5E4DS2TOS",
				Type:      257,
				Deadline:  9149600,
				Message: message{
					Payload: "280000005444334b32493543524850595634425a5a5a4c335850454e4",
					Type:    2,
				},
				Version: -1744830463,
				Signer:  "c20a1dffe699c7a68328986273265e33fceebe074f274240ef890dd80ad55ed6",
			},
		},
		TransactionMetadataPair{
			TransactionMetadata{
				ID:     71356,
				Height: 40629,
				Hash: hash{
					Data: "37c34ead4c3fe6af42d994135798262f785ba2d807c02ac3608bc10da12e5f87",
				},
			},
			Transaction{
				TimeStamp: 9101541,
				Amount:    49997995000000,
				Signature: "57c3c48d2ae8b24240b57d72493f498cfeb61e2ab87237dc0e08c51007d5c7f15847d0e08c0286e68a72028925db5fa809ca9d57e2cb6eebe11822176a834c0b",
				Fee:       2005000000,
				Recipient: "TALICELCD3XPH4FFI5STGGNSNSWPOTG5E4DS2TOS",
				Type:      257,
				Deadline:  9144741,
				Message: message{
					Payload: "526f6262657279212121",
					Type:    1,
				},
				Version: -1744830463,
				Signer:  "546e4fb9c81db84e04d8e9e67380db0fe1f540df09a527fb995b589b5695ae24",
			},
		},
	}
	got, err := clientMock.OutgoingTransactions("TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\nWanted: %v\n Got: %v", want, got)
	}
}

func TestAllTransactions(t *testing.T) {
	want := []TransactionMetadataPair{
		TransactionMetadataPair{
			TransactionMetadata{
				ID:     71245,
				Height: 40706,
				Hash: hash{
					Data: "15c373ad4c3fe6af47d1941379ff262f785bdcfa07c02ac3608bc10da27d5e82",
				},
			},
			Transaction{
				TimeStamp: 9106400,
				Amount:    1000000000,
				Signature: "449cd76ea8bda2220b3d6ad6f8db5f81d4e68ad3d4b0c3db9a3c267355657639eabed3dbcef8e0cc22953ae2b36a22ee7dc6327484c9649cccd686a511eca105",
				Fee:       3000000,
				Recipient: "TALICELCD3XPH4FFI5STGGNSNSWPOTG5E4DS2TOS",
				Type:      257,
				Deadline:  9149600,
				Message: message{
					Payload: "280000005444334b32493543524850595634425a5a5a4c335850454e4",
					Type:    2,
				},
				Version: -1744830463,
				Signer:  "c20a1dffe699c7a68328986273265e33fceebe074f274240ef890dd80ad55ed6",
			},
		},
		TransactionMetadataPair{
			TransactionMetadata{
				ID:     71356,
				Height: 40629,
				Hash: hash{
					Data: "37c34ead4c3fe6af42d994135798262f785ba2d807c02ac3608bc10da12e5f87",
				},
			},
			Transaction{
				TimeStamp: 9101541,
				Amount:    49997995000000,
				Signature: "57c3c48d2ae8b24240b57d72493f498cfeb61e2ab87237dc0e08c51007d5c7f15847d0e08c0286e68a72028925db5fa809ca9d57e2cb6eebe11822176a834c0b",
				Fee:       2005000000,
				Recipient: "TALICELCD3XPH4FFI5STGGNSNSWPOTG5E4DS2TOS",
				Type:      257,
				Deadline:  9144741,
				Message: message{
					Payload: "526f6262657279212121",
					Type:    1,
				},
				Version: -1744830463,
				Signer:  "546e4fb9c81db84e04d8e9e67380db0fe1f540df09a527fb995b589b5695ae24",
			},
		},
	}
	got, err := clientMock.AllTransactions("TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\nWanted: %v\n Got: %v", want, got)
	}
}
