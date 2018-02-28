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
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// TODO finish this method

// MosaicSupplyInfo contains information regarding the current supply of a mosaic
type MosaicSupplyInfo struct {
	something int
}

// Supply gets teh current supply of a mosaic
func Supply(u url.URL, ID string) (MosaicSupplyInfo, error) {
	u.Path = "/mosaic/supply"
	q := u.Query()
	q.Set("mosaicId", ID)
	u.RawQuery = q.Encode()
	options := Options{
		URL:    u,
		Method: http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return MosaicSupplyInfo{}, err
	}
	fmt.Println(string(resp))
	var data MosaicSupplyInfo
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return MosaicSupplyInfo{}, err
	}
	return data, nil
}
