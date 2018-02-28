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

// Send uses a set of options to create and send a request to the NEM blockchain.
func Send(options Options) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(options.Method, options.URL.String(), bytes.NewBuffer(options.Body))
	if err != nil {
		return nil, err
	}
	for k, v := range options.Headers {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
