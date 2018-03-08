package main

import (
	"fmt"
	"log"

	"github.com/myndshft/nemgo"
)

func main() {
	// Connect to the testnet
	client, err := nemgo.NewClient(byte(0x98))
	if err != nil {
		log.Fatal(err)
	}
	// Obtain blocks channel
	blocks, err := client.SubscribeHeight()
	if err != nil {
		log.Fatal(err)
	}
	// Iterate over channel
	for block := range blocks {
		fmt.Println(block)
	}
	close(blocks)
}
