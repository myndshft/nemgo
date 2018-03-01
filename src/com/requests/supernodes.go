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

	"github.com/myndshft/nem-go-sdk/src/model"
)

// SuperNodeInfo contains information about all the current Supernodes on the network.
// These nodes form a backbone of support for light wallets, mobile wallets, and
// 3rd party apps so that users of these services might have access to the
// network that is easy, fast, and reliable without having to sync a blockchain
// by themselves or use untrustworthy centralized services.
// Supernodes are expected to be high performance and reliable nodes.
// They are regularly tested on their bandwidth, chain height, chain parts,
// computing power, version, ping, and responsiveness to make sure that they
// are performing to high standards.
// If they meet all these requirements, they are randomly given rewards.
type SuperNodeInfo struct {
	// Nodes contains information about all current Supernodes.
	Nodes []SuperNodeDefinition `json:"nodes"`
	// NodeCount contains the number of Supernodes currently on the network.
	NodeCount int `json:"nodeCount"`
}

// SuperNodeDefinition contains all information about a single Supernode.
type SuperNodeDefinition struct {
	// ID contains the node id.
	// Sometimes this is a string and sometimes it is an int
	// Validate before using this field!
	ID interface{} `json:"id"`
	// Alias contains the name of the node.
	Alias string `json:"alias"`
	// IP contains the IP address of the node.
	IP string `json:"ip"`
	// NisPort contains the exposed port on the node.
	NisPort int `json:"nisPort"`
	// PubKey contains the public key of the node.
	PubKey string `json:"pubKey"`
	// ServantPort contains the exposed servant port on the node.
	ServantPort int `json:"servantPort"`
	// Status contains the node's current status.
	Status int `json:"status"`
	// Latitude contains the node's current latitude.
	Latitude float64 `json:"latitude"`
	// Longitude contains the node's current longitude.
	Longitude float64 `json:"longitude"`
	// PayoutAddress contains the address of the account to which any bonus payout
	// will be sent.
	PayoutAddress string `json:"payoutAddress"`
	// Distance is the distance from the coordinates input (if exists).
	Distance float64 `json:"distance,omitempty"`
	// MaxUnlocked contains the maximum number of harvesters allowed on the node (if exists).
	MaxUnlocked int `json:"maxUnlocked,omitempty"`
	// numUnlocked contains the number of harvesters currently harvesting on the node (if exists).
	NumUnlocked int `json:"numUnlocked,omitempty"`
}

type superNodeDefinitionRecords struct {
	Records []SuperNodeDefinition `json:"data"`
}

func (sni *SuperNodeDefinition) unmarshal(data []byte) ([]SuperNodeDefinition, error) {
	var tmpData superNodeDefinitionRecords
	err := json.Unmarshal(data, &tmpData)
	if err != nil {
		return []SuperNodeDefinition{}, err
	}
	return tmpData.Records, nil
}

// All gets all nodes of the node reward program
func All(sender Sender) (SuperNodeInfo, error) {
	options := Options{
		URL:    model.Supernodes,
		Method: http.MethodGet}
	senderOpts := NewDefaultSenderOptions(options)
	resp, err := sender.Send(senderOpts)
	if err != nil {
		return SuperNodeInfo{}, err
	}
	var data SuperNodeInfo
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return SuperNodeInfo{}, err
	}
	return data, nil
}

// Coordinates contains the Latitude and Longitude of a coorindate on Earth
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// Nearest gets the nearest supernodes up to total
func Nearest(sender Sender, coords Coordinates, total int) ([]SuperNodeDefinition, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"latitude":  coords.Latitude,
		"longitude": coords.Longitude,
		"numNodes":  total})
	if err != nil {
		return []SuperNodeDefinition{}, err
	}
	options := Options{
		URL:    model.NearestSupernodes,
		Method: http.MethodPost,
		Body:   payload}
	senderOpts := NewDefaultSenderOptions(options)
	resp, err := sender.Send(senderOpts)
	if err != nil {
		return []SuperNodeDefinition{}, err
	}
	var snd SuperNodeDefinition
	data, err := snd.unmarshal(resp)
	if err != nil {
		return []SuperNodeDefinition{}, err
	}
	return data, nil
}

// Get gets all the supernodes by status
func Get(sender Sender, status int) ([]SuperNodeDefinition, error) {
	payload, err := json.Marshal(map[string]int{"status": status})
	if err != nil {
		return []SuperNodeDefinition{}, err
	}
	options := Options{
		URL:    model.SupernodesByStatus,
		Method: http.MethodPost,
		Body:   payload}
	senderOpts := NewDefaultSenderOptions(options)
	resp, err := sender.Send(senderOpts)
	if err != nil {
		return []SuperNodeDefinition{}, err
	}
	var snd SuperNodeDefinition
	data, err := snd.unmarshal(resp)
	if err != nil {
		return []SuperNodeDefinition{}, err
	}
	return data, nil

}
