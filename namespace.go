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
	"strconv"
)

// NamespaceMetadataPair contains info and metadata about a namespace
type NamespaceMetadataPair struct {
	NamespaceInfo     NamespaceInfo
	NamespaceMetaData NamespaceMetadata
}

// NamespaceMetadata contains meta information about a namespace
type NamespaceMetadata struct {
	ID int
}

// NamespaceInfo contains information about a namespace
type NamespaceInfo struct {
	FQN    string
	Owner  string
	Height int
}

// RootNamespace will get all the root namespaces in batch of a specified PageSize.
func (c Client) RootNamespace(ID int, PageSize int) ([]NamespaceMetadataPair, error) {
	data := struct{ Data []NamespaceMetadataPair }{}
	c.url.Path = "/namespace/root/page"
	req, err := c.buildReq(map[string]string{"id": strconv.Itoa(ID), "pagesize": strconv.Itoa(PageSize)}, nil, http.MethodGet)
	if err != nil {
		return data.Data, err
	}
	body, err := c.request(req)
	if err != nil {
		return data.Data, err
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return data.Data, nil
	}
	return data.Data, nil
}

// Namespace will return a NamespaceInfo about a namespace
func (c Client) Namespace(namespace string) (NamespaceInfo, error) {
	var data NamespaceInfo
	c.url.Path = "/namespace"
	req, err := c.buildReq(map[string]string{"namespace": namespace}, nil, http.MethodGet)
	if err != nil {
		return data, err
	}
	body, err := c.request(req)
	if err != nil {
		return data, err
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return data, err
	}
	return data, nil
}
