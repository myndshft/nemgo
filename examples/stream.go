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
