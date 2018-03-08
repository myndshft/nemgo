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
	"encoding/json"
	"net/http"
)

// BlockHeight contains a chain height
type BlockHeight struct {
	Height int
}

// Height gets the current height of the blockchain
func (c Client) Height() (BlockHeight, error) {
	c.URL.Path = "/chain/height"
	req, err := c.buildReq(nil, nil, http.MethodGet)
	if err != nil {
		return BlockHeight{}, err
	}
	body, err := c.Request(req)
	if err != nil {
		return BlockHeight{}, err
	}
	var data BlockHeight
	if err := json.Unmarshal(body, &data); err != nil {
		return BlockHeight{}, err
	}
	return data, nil
}
