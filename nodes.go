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

type metadata struct {
	Features    int
	Application string
	NetworkID   int
	Version     string
	Platform    string
}
type endpoint struct {
	Protocol string
	Port     int
	Host     string
}
type identity struct {
	Name      string
	PublicKey string
}

// Node is a node on the Nem blockchain
type Node struct {
	MetaData metadata
	Endpoint endpoint
	Identity identity
}

type nisInfo struct {
	CurrentTime int
	Application string
	StartTime   int
	Version     string
	Signer      string
}

// NodeInfo contains extended information about the current node
type NodeInfo struct {
	Node    Node
	NISInfo nisInfo
}

// NodeInfo gets basic information about a node
func (c Client) NodeInfo() (NodeInfo, error) {
	var data NodeInfo
	c.url.Path = "/node/extended-info"
	req, err := c.buildReq(nil, nil, http.MethodGet)
	if err != nil {
		return data, err
	}
	body, err := c.request(req)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return data, err
	}
	return data, nil
}
