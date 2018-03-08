package nemgo

import (
	"errors"
	"net/http"
	"net/url"
)

const defaultTestnetHost = "23.228.67.85:7890"
const defaultMainnetHost = "209.126.98.204:7890"

// Client is a new client for connected to a NIS Node
type Client struct {
	Network byte
	URL     url.URL
	Request func(*http.Request) ([]byte, error)
}

// NewClient will return a Client object ready to be used
// on whichever chain is passed in
func NewClient(network byte) (Client, error) {
	// TODO instantiate new client
	var host string
	switch network {
	case byte(0x68):
		host = defaultMainnetHost
	case byte(0x98):
		host = defaultTestnetHost
	default:
		return Client{}, errors.New("please provide a valid network")
	}
	return Client{
		Network: network,
		URL:     url.URL{Scheme: "http", Host: host},
		Request: sendReq}, nil
}
