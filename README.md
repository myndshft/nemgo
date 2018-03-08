<img src=".github/nemgo.png" height="136">

# nemgo

A pure golang SDK for the Nem blockchain. 

This project is in it's infancy and looking for more contributors! If you are working in Go, are interested in blockchain technologies, or just want to join a friendly open source project you are welcome!

## Getting Started

`go get` the package using the following command:

```bash
$ go get github.com/myndshft/nemgo
```

Open up your favorite text editor, create a new `Client` and interact with the blockchain!

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/myndshft/nemgo"
)

// testnet = byte(0x98)
// mainnet = byte(0x68

func main() {
    client, err := nemgo.NewClient(byte(0x98))
    if err != nil {
        log.Fatal(err)
    }
    
    // Get account information
    address = "YOUR ACCOUNT ADDRESS"
    actInfo, err := client.AccountInfo(address)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(actInfo)
    
    // Get the current height of the chain
    height, err := client.Height()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(height)
    
    // Subscribe to transactions related to an account
    // This will return a go channel
    txs, err := client.SubscribeUnconfirmedTX(address)
    if err != nil {
        log.Fatal(err)
    }
    defer close(txs)
    for tx := range txs {
        fmt.Println(tx)
    }
```

## Helping out

Check out the `CONTRIBUTING.md` documents in the `docs` folder. We always welcome any contribution, large or small! 

_The gopher logo is the work of Renee French. The Nem logo is licensed under CC0 1.0_
