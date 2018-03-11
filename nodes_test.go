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

func TestNodeInfo(t *testing.T) {
	want := NodeInfo{
		Node: Node{
			metadata{
				Features:    1,
				Application: "NIS",
				NetworkID:   -104,
				Version:     "0.4.33-BETA",
				Platform:    "Oracle Corporation (1.8.0_25) on Windows 8"},
			endpoint{
				Protocol: "http",
				Port:     7890,
				Host:     "81.224.224.156"},
			identity{
				Name:      "Alice",
				PublicKey: "a1aaca6c17a24252e674d155713cdf55996ad00175be4af02a20c67b59f9fe8a"}},
		NISInfo: nisInfo{
			CurrentTime: 9288341,
			Application: "NEM Infrastructure Server",
			StartTime:   9238484,
			Version:     "0.4.33-BETA",
			Signer:      "CN=VeriSign Class 3 Code Signing 2010 CA,OU=Terms of use at https://www.verisign.com/rpa (c)10,OU=VeriSign Trust Network,O=VeriSign\\, Inc.,C=US"}}
	got, err := clientMock.NodeInfo()
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted: %v\n    Got: %v", want, got)
	}
}
