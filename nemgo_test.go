package nemgo

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestNewWithOptions(t *testing.T) {
	want := Client{
		network: byte(0x98),
		url:     url.URL{Scheme: "http", Host: "23.228.67.85:7890"},
		request: sendReq}
	got := New(WithNIS("23.228.67.85:7890", byte(0x98)))
	if reflect.DeepEqual(want, got) {
		t.Fatalf("\nWanted: %v\n Got: %v", want, got)
	}
}

func TestNew(t *testing.T) {
	want := Client{
		url:     url.URL{Scheme: "http", Host: "209.126.98.204:7890"},
		request: sendReq}
	got := New()
	if reflect.DeepEqual(want, got) {
		t.Fatalf("\nWanted: %v\n Got: %v", want, got)
	}
}

func ExampleNew() {
	c := New()
	// use c here
	fmt.Println(c)
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
	case "/chain/score":
		return []byte(blockScore), nil
	case "/chain/last-block", "/block/at/public":
		return []byte(block), nil
	case "/node/extended-info":
		return []byte(node), nil
	case "/namespace/root/page":
		return []byte(namespaceMetaDataPair), nil
	case "/namespace":
		return []byte(namespaceInfo), nil
	case "/account/transfers/incoming", "/account/transfers/outgoing", "/account/transfers/all":
		return []byte(transactionMetadataPairArray), nil
	default:
		return nil, nil
	}
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

const blockScore = `{
	"score": "18722d5a7d590deb"
}`

const block = `{
       "timeStamp": 9232968,
       "signature": "0a1351ef3e9b19c601e804a6d329c9ade662051d1da2c12c3aec9934353e421c79de7d8e59b127a8ca9b9d764e3ca67daefcf1952f71bc36f747c8a738036b05",
       "prevBlockHash": {
              "data": "58efa578aea719b644e8d7c731852bb26d8505257e03a897c8102e8c894a99d6"
       },
       "type": 1,
       "transactions": [
       ],
       "version": 1744830465,
       "signer": "2afca04d2cb8d16cf3656274bc55b95e60be823cfb7230d82f791ed42a309ee7",
       "height": 42804
}`

const node = `{
       "node": {
              "metaData":
              {
                     "features": 1,
                     "application": "NIS",
                     "networkId": -104,
                     "version": "0.4.33-BETA",
                     "platform": "Oracle Corporation (1.8.0_25) on Windows 8"
              },
              "endpoint":
              {
                     "protocol": "http",
                     "port": 7890,
                     "host": "81.224.224.156"
              },
              "identity":
              {
                     "name": "Alice",
                     "public-key": "a1aaca6c17a24252e674d155713cdf55996ad00175be4af02a20c67b59f9fe8a"
              }
       },
       "nisInfo":
       {
              "currentTime": 9288341,
              "application": "NEM Infrastructure Server",
              "startTime": 9238484,
              "version": "0.4.33-BETA",
              "signer": "CN=VeriSign Class 3 Code Signing 2010 CA,OU=Terms of use at https://www.verisign.com/rpa (c)10,OU=VeriSign Trust Network,O=VeriSign\\, Inc.,C=US"
       }
}`

const namespaceMetaDataPair = `{
        "data": [{
            "meta": {
                "id": 26264
            },
            "namespace": {
                "fqn": "makoto.metal.coins",
                "owner": "TD3RXTHBLK6J3UD2BH2PXSOFLPWZOTR34WCG4HXH",
                "height": 13465
            }
        },{
            "meta": {
                "id": 25421
            },
            "namespace": {
                "fqn": "gimre.vouchers",
                "owner": "TDGIMREMR5NSRFUOMPI5OOHLDATCABNPC5ID2SVA",
                "height": 12392
            }
        }]
}`

const namespaceInfo = `{
        "fqn": "makoto.metal.coins",
        "owner": "TD3RXTHBLK6J3UD2BH2PXSOFLPWZOTR34WCG4HXH",
        "height": 13465
}`

const transactionMetadataPairArray = `{
       "data": [
       {
              "meta":
              {
                     "id": 71245,
                     "height": 40706,
                     "hash": {
                         "data":"15c373ad4c3fe6af47d1941379ff262f785bdcfa07c02ac3608bc10da27d5e82"
                     }
              },
              "transaction":
              {
                     "timeStamp": 9106400,
                     "amount": 1000000000,
                     "signature": "449cd76ea8bda2220b3d6ad6f8db5f81d4e68ad3d4b0c3db9a3c267355657639eabed3dbcef8e0cc22953ae2b36a22ee7dc6327484c9649cccd686a511eca105",
                     "fee": 3000000,
                     "recipient": "TALICELCD3XPH4FFI5STGGNSNSWPOTG5E4DS2TOS",
                     "type": 257,
                     "deadline": 9149600,
                     "message":
                     {
                           "payload": "280000005444334b32493543524850595634425a5a5a4c335850454e4",
                           "type": 2
                     },
                     "version": -1744830463,
                     "signer": "c20a1dffe699c7a68328986273265e33fceebe074f274240ef890dd80ad55ed6"
              }
       },
       {
              "meta":
              {
                     "id": 71356,
                     "height": 40629,
                     "hash": {
                         "data":"37c34ead4c3fe6af42d994135798262f785ba2d807c02ac3608bc10da12e5f87"
                     }
              },
              "transaction":
              {
                     "timeStamp": 9101541,
                     "amount": 49997995000000,
                     "signature": "57c3c48d2ae8b24240b57d72493f498cfeb61e2ab87237dc0e08c51007d5c7f15847d0e08c0286e68a72028925db5fa809ca9d57e2cb6eebe11822176a834c0b",
                     "fee": 2005000000,
                     "recipient": "TALICELCD3XPH4FFI5STGGNSNSWPOTG5E4DS2TOS",
                     "type": 257,
                     "deadline": 9144741,
                     "message":
                     {
                           "payload": "526f6262657279212121",
                           "type": 1
                     },
                     "version": -1744830463,
                     "signer": "546e4fb9c81db84e04d8e9e67380db0fe1f540df09a527fb995b589b5695ae24"
              }
       }]
}`
