package main

import (
	"net/url"

	"github.com/myndshft/nem-go-sdk/src/com/requests"
)

const defaultTestnet = "bigalice2.nem.ninja:7890"

var testAddress = "TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S"
var privateKey = ""

func main() {
	// data := requests.MosaicDefinitionProperties{Name: "name", Value: "value"}
	url := url.URL{Scheme: "http", Host: defaultTestnet}
	// resp, err := requests.GetHistoricalAccountData(url, testAddress, 1000)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp)
	// coords := requests.Coordinates{Latitude: 27.2, Longitude: 80.1}
	// ra := requests.RequestAnnounce{
	// Data:      "010100000100000000000000200000002b76078fa709bbe6752222b215abc7ec0152ffe831fb4f9aed3e7749a425900a00093d0000000000000000002800000054444e46555946584f5353334e4e4c4f35465a5348535a49354c33374b4e5149454850554d584c54c0d45407000000000b00000001000000030000000c3215",
	// Signature: "db2473513c7f0ce9f8de6345f0fbe773dc687eb571123d08eab4d98f96849eaeb63fa8756fb6c59d9b9d0e551537c1cdad4a564747ff9291db4a88b65c97c10d"}
	requests.GetBatchHistoricalAccountData(url, []string{testAddress, testAddress}, 3300)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp)
	// resp, err := requests.GetBatchAccountData(url, []string{testAddress})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp)
	// requests.TestSend(data)
	// resp, err := requests.Supply(url, "")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp)
	// resp, err := requests.Heartbeat(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp)
	// resp, err := requests.HarvestedBlocks(url, testAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp[0].Difficulty)
	// err := requests.StartHarvesting(url, privateKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // _, err := requests.IncomingTransactions(url, testAddress, "", "")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp)
	// data, err := requests.DataFromPublicKey(url, testAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(data)
	// account, err := requests.HarvestingBlocks(url, testAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(account)

	// height, err := height(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(height)
}

// // Get the height of the current blockchain
// func height(u url.URL) (map[string]interface{}, error) {
// 	u.Path = "/chain/height"
// 	options := requests.Options{
// 		URL:    u,
// 		Method: http.MethodGet}
// 	resp, err := requests.Send(options)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp, nil
// }
