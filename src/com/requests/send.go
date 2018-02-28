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
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Options contains the information needed to successfully communicate with
// the NEM blockchain.
type Options struct {
	// URL contains all information regarding the Scheme, Host, Port, and Path
	// as well as any additional query parameters.
	URL url.URL
	// Method contains a string of the HTTP method to be used.
	Method string
	// Headers contains the information regarding what type of request
	// is being sent.
	Headers map[string]string
	// Body contains information to be sent in the request body.
	// If left empty, body will not affect the query.
	Body []byte
}

// Sender is the interface to Requester objects
type Sender interface {
	Send(SenderOptions) ([]byte, error)
}

// Requester is used to send requests to NEM blockchain
type Requester struct{}

// Send will use the passed in builder and parser to process information
// from the NEM blockchain
func (Requester) Send(s SenderOptions) ([]byte, error) {
	req, err := s.builder(s.options)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	return s.parser(resp)
}

// SenderOptions has everything needed for a Requester to send data to NEM
type SenderOptions struct {
	client  *http.Client
	options Options
	builder func(options Options) (*http.Request, error)
	parser  func(resp *http.Response) ([]byte, error)
}

// NewDefaultSenderOptions will create a new SenderOptions object with
// default values. This is used for ease of readibility and not needing to
// instantiate the entier SenderOptions by hand every time.
func NewDefaultSenderOptions(options Options) SenderOptions {
	return SenderOptions{
		client:  &http.Client{},
		options: options,
		builder: buildRequest,
		parser:  parseResponse}
}

func buildRequest(options Options) (*http.Request, error) {
	req, err := http.NewRequest(options.Method, options.URL.String(), bytes.NewBuffer(options.Body))
	if err != nil {
		return nil, err
	}
	for k, v := range options.Headers {
		req.Header.Add(k, v)
	}
	return req, nil
}

func parseResponse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
