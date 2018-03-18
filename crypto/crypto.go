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

package crypto

// func AccountFromPubKey(pubkey string, network nemgo.Network) (string, error) {
// 	h := sha3.Sum256([]byte(pubkey))
// 	r := ripemd160.New()
// 	_, err := r.Write(h[:])
// 	if err != nil {
// 		return "", err
// 	}
// 	b := append(network, r.Sum(nil)...)
// 	h = sha3.Sum256(b)
// 	a := append(b, h[:4])
// 	return a

// }
