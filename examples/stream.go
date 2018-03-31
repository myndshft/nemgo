package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/myndshft/nemgo"
)

func main() {
	// Connect to the testnet
	client := nemgo.New(nemgo.WithNIS("193.70.91.98:7890", nemgo.Mainnet))
	var wg sync.WaitGroup
	// Obtain blocks channel
	blocks, err := client.SubscribeConfirmedTX("ND2JRPQIWXHKAA26INVGA7SREEUMX5QAI6VU7HNR")
	if err != nil {
		log.Fatal(err)
	}
	// Iterate over channel
	go func() {
		wg.Add(1)
		for block := range blocks {
			fmt.Println(block)
		}
		wg.Done()
	}()
	heights, err := client.SubscribeHeight()
	if err != nil {
		log.Fatal(err)
	}
	// Iterate over channel
	go func() {
		wg.Add(1)
		for height := range heights {
			fmt.Println(height)
		}
		wg.Done()
	}()
	wg.Wait()
}
