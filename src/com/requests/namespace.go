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
	"net/http"
	"net/url"
	"strconv"
)

// Roots gets root namespaces
func Roots(u url.URL, ID int) (NamespaceMetaDataPair, error) {
	u.Path = "/namespace/root/page"
	q := u.Query()
	q.Set("pageSize", "100")
	if ID != 0 {
		q.Set("id", strconv.Itoa(ID))
	}
	u.RawQuery = q.Encode()
	options := Options{
		URL:    u,
		Method: http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return NamespaceMetaDataPair{}, err
	}
	var data NamespaceMetaDataPair
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return NamespaceMetaDataPair{}, err
	}
	return data, nil

}

// MosaicDefinitions gets mosaic definitions of a namespace
func MosaicDefinitions(u url.URL, ID string) (MosaicDefinition, error) {
	u.Path = "/namespace/mosaic/definition/page"
	q := u.Query()
	q.Set("namespace", ID)
	u.RawQuery = q.Encode()
	options := Options{
		URL:    u,
		Method: http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return MosaicDefinition{}, err
	}
	var data MosaicDefinition
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return MosaicDefinition{}, err
	}
	return data, nil
}

// NamespaceInfo gets the namespace information for a given ID
func NamespaceInfo(u url.URL, ID string) (NamespaceData, error) {
	u.Path = "/namespace"
	q := u.Query()
	q.Set("namespace", ID)
	u.RawQuery = q.Encode()
	options := Options{
		URL:     u,
		Method:  http.MethodGet,
		Headers: URLEncoded}
	resp, err := Send(options)
	if err != nil {
		return NamespaceData{}, err
	}
	var data NamespaceData
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return NamespaceData{}, err
	}
	return data, nil
}
