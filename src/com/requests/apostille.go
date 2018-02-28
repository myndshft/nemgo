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

package requests

import (
	"fmt"
	"net/http"

	"github.com/myndshft/nem-go-sdk/src/model"
)

// TODO test this and make sure it works

// Audit will audit an apostille file
func Audit(sender Sender, publicKey string, data string, signedData string) (bool, error) {
	u := model.ApostilleAuditServer
	q := u.Query()
	q.Set("publicKey", publicKey)
	q.Set("data", data)
	q.Set("signedData", signedData)
	u.RawQuery = q.Encode()
	options := Options{
		URL:    u,
		Method: http.MethodPost}
	senderOpts := NewDefaultSenderOptions(options)
	resp, err := sender.Send(senderOpts)
	if err != nil {
		return false, err
	}
	// TODO this is unfinished, check the response before returning
	fmt.Println(string(resp))
	return true, nil

}
