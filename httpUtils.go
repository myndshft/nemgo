package nemgo

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func (c Client) buildReq(params map[string]string, body []byte, method string) (*http.Request, error) {
	if params != nil {
		q := c.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		c.URL.RawQuery = q.Encode()
	}
	req, err := http.NewRequest(method, c.URL.String(), bytes.NewBuffer(body))
	if err != nil {
		return &http.Request{}, err
	}
	return req, nil
}

func sendReqMock(req *http.Request) ([]byte, error) {
	switch req.URL.Path {
	case "/account/batch":
		return []byte(accountMetaDataPairNested), nil
	case "/account/get", "/account/get/forwarded":
		return []byte(accountMetaDataPair), nil
	case "/account/status":
		return []byte(accountMetaData), nil
	case "/account/harvests":
		return []byte(harvestInfo), nil
	case "/account/mosaic/owned":
		return []byte(ownedmosaic), nil
	case "/chain/height":
		return []byte(blockHeight), nil
	default:
		return nil, nil
	}
}

func sendReq(req *http.Request) ([]byte, error) {
	client := http.Client{Timeout: 10 * time.Second}
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

const accountMetaDataPairNested = `{
   "data":[
      {
         "account":{
            "Address":"TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S",
            "Balance":124446551689680,
            "VestedBalance":1041345514976241,
            "Importance":0.010263666447108395,
            "PublicKey":"a11a1a6c17a24252e674d151713cdf51991ad101751e4af02a20c61b59f1fe1a",
			"Label": null,
			"HarvestedBlocks": 645
         },
         "meta":{
            "Status":"LOCKED",
            "RemoteStatus":"ACTIVE",
            "CosignatoryOf":[

            ],
            "Cosignatories":[

            ]
         }
      },
      {
         "account":{
            "Address":"TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S",
            "Balance":124446551689680,
            "VestedBalance":1041345514976241,
            "Importance":0.010263666447108395,
            "PublicKey":"a11a1a6c17a24252e674d151713cdf51991ad101751e4af02a20c61b59f1fe1a",
			"Label": null,
			"HarvestedBlocks": 645
         },
         "meta":{
            "Status":"LOCKED",
            "RemoteStatus":"ACTIVE",
            "CosignatoryOf":[

            ],
            "Cosignatories":[

            ]
         }
      }
   ]
}`

const accountMetaDataPair = `{
  "account":{
     "Address":"TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S",
     "Balance":124446551689680,
     "VestedBalance":1041345514976241,
     "Importance":0.010263666447108395,
     "PublicKey":"a11a1a6c17a24252e674d151713cdf51991ad101751e4af02a20c61b59f1fe1a",
     "Label": null,
     "HarvestedBlocks": 645
  },
  "meta":{
     "Status":"LOCKED",
     "RemoteStatus":"ACTIVE",
     "CosignatoryOf":[

     ],
     "Cosignatories":[

     ]
  }
}`

const accountMetaData = `{
     "Status":"LOCKED",
     "RemoteStatus":"ACTIVE",
     "CosignatoryOf":[

     ],
     "Cosignatories":[

     ]
  }`

const harvestInfo = `{
       "data": [{
              "timeStamp": 8879051,
              "difficulty": 26453656336676,
              "totalFee": 102585065,
              "id": 1262068,
              "height": 37015
       },{
              "timeStamp": 8879051,
              "difficulty": 26453656336676,
              "totalFee": 102585065,
              "id": 1262068,
              "height": 37015}
	   ]
}`

const ownedmosaic = `{
        "data": [{
            "mosaicId": {
                "namespaceId": "alice.drinks",
                "name": "orange juice"
            },
            "quantity": 123
        },{
            "mosaicId": {
                "namespaceId": "alice.drinks",
                "name": "orange juice"
            },
            "quantity": 123
        }]
}`

const blockHeight = `{
	"height": 12345
}`
