// +build OMIT

package main

import (
	"fmt"
	"log"

	"github.com/myndshft/nemgo"
)

func main() {
	// Connect to the mainnet
	client := nemgo.New()
	// Obtain account data using address
	a, err := client.AccountData(nemgo.Address("NBMBTUB6JIXGBSETDJBMCGLB2GPTI6GMPAYFNH3P"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", a)
	// Obtain account data using public key
	b, err := client.AccountData(nemgo.PublicKey("4f790ccd43426dcb663a80283294ed3e852a855723d87763a7608f1b45d27e20"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", b)
}
