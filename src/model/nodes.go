package model

import "net/url"

// DefaultTestnet is the default testnet node
const DefaultTestnet = "bigalice2.nem.ninja"

// DefaultMainnet is the default mainnet node
const DefaultMainnet = "alice6.nem.ninja"

// DefaultMijin is the default mijin node
const DefaultMijin = ""

// MainnetExplorer is the default mainnet block explorer
const MainnetExplorer = "chain.nem.ninja/#/transfer/"

// TestnetExplorer is the default mainnet block explorer
const TestnetExplorer = "bob.nem.ninja:8765/#/transfer/"

// MijinExplorer is the default mijin block explorer
const MijinExplorer = ""

type nodeInformation struct {
	uri      url.URL
	location string
}

// TODO look into how we can make these const

// SearchOnTestnet is an array of the nodes allowing search by transaction hash on testnet
var SearchOnTestnet = []nodeInformation{
	{uri: url.URL{Scheme: "http", Host: "bigalice2.nem.ninja"},
		location: "America / New_York"},
	{uri: url.URL{Scheme: "http", Host: "192.3.61.243"},
		location: "America / Los_Angeles"},
	{uri: url.URL{Scheme: "http", Host: "23.228.67.85"},
		location: "America / Los_Angeles"}}

// SearchOnMainnet is an array of the nodes allowing search by transaction hash on mainnet
var SearchOnMainnet = []nodeInformation{
	{uri: url.URL{Scheme: "http", Host: "62.75.171.41"},
		location: "Germany"},
	{uri: url.URL{Scheme: "http", Host: "104.251.212.131"},
		location: "USA"},
	{uri: url.URL{Scheme: "http", Host: "45.124.65.125"},
		location: "Hong Kong"},
	{uri: url.URL{Scheme: "http", Host: "185.53.131.101"},
		location: "Netherlands"},
	{uri: url.URL{Scheme: "http", Host: "sz.nemchina.com"},
		location: "China"}}

// SearchOnMijin is an array of nodes allowing earch by transaction hash on mijin
var SearchOnMijin = []nodeInformation{
	{uri: url.URL{},
		location: ""}}

// Testnet are the testnet nodes
var Testnet = []nodeInformation{
	{uri: url.URL{Scheme: "http", Host: "104.128.226.60"}},
	{uri: url.URL{Scheme: "http", Host: "23.228.67.85"}},
	{uri: url.URL{Scheme: "http", Host: "192.3.61.243"}},
	{uri: url.URL{Scheme: "http", Host: "50.3.87.123"}},
	{uri: url.URL{Scheme: "http", Host: "localhost"}}}

// Mainnet are the mainnet nodes
var Mainnet = []nodeInformation{
	{uri: url.URL{Scheme: "http", Host: "62.75.171.41"}},
	{uri: url.URL{Scheme: "http", Host: "san.nem.ninja"}},
	{uri: url.URL{Scheme: "http", Host: "go.nem.ninja"}},
	{uri: url.URL{Scheme: "http", Host: "hachi.nem.ninja"}},
	{uri: url.URL{Scheme: "http", Host: "jusan.nem.ninja"}},
	{uri: url.URL{Scheme: "http", Host: "nijuichi.nem.ninja"}},
	{uri: url.URL{Scheme: "http", Host: "alice2.nem.ninja"}},
	{uri: url.URL{Scheme: "http", Host: "alice3.nem.ninja"}},
	{uri: url.URL{Scheme: "http", Host: "alice4.nem.ninja"}},
	{uri: url.URL{Scheme: "http", Host: "alice5.nem.ninja"}},
	{uri: url.URL{Scheme: "http", Host: "alice6.nem.ninja"}},
	{uri: url.URL{Scheme: "http", Host: "alice7.nem.ninja"}},
	{uri: url.URL{Scheme: "http", Host: "localhost"}}}

// Mijin are the mijin nodes
var Mijin = []nodeInformation{
	{uri: url.URL{}}}

// ApostilleAuditServer is the server verifying signed apostilles
var ApostilleAuditServer = url.URL{Scheme: "https", Host: "185.117.22.58:4567", Path: "/verify"}

// Supernodes is the API endpoint to get all supernodes
var Supernodes = url.URL{Scheme: "https", Host: "supernodes.nem.io", Path: "/nodes"}

// NearestSupernodes is the API endpoint to get the nearest supernodes
var NearestSupernodes = url.URL{Scheme: "http", Host: "199.217.113.179:7782", Path: "/nodes/nearest"}

// SupernodesByStatus is the API endpoint to get the supernodes by current status
var SupernodesByStatus = url.URL{Scheme: "http", Host: "199.217.113.179:7782", Path: "/nodes"}

// MarketInfo is the API endpoint to get XEM/BTC market data
var MarketInfo = url.URL{Scheme: "https", Host: "poloniex.com", Path: "/public"}

// BTCPrice is the API to get BTC/USD market data
var BTCPrice = url.URL{Scheme: "https", Host: "blockchain.info", Path: "/ticker"}

// DefaultPort is the default endpoint port
const DefaultPort = 7890

// MijinPort is the default endpoint port for mijin network
const MijinPort = 7895

// WebsocketPort is the default websocket port
const WebsocketPort = 7778
