package main

import (
	"encoding/base32"
	"encoding/hex"
	"fmt"

	"github.com/myndshft/nemgo"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

var pub = "7657e3b29e6d64f951f1ca3371630dfc017b5a440f7fa6f9bc4f7f6c191534ab"
var want = "TCQKEW2GQ25L44CGZJVQRGL4LGO7OR2OX43YYTKA"

func main() {
	pk, err := hex.DecodeString(pub)
	if err != nil {
		panic(err)
	}
	h := sha3.Sum256(pk)
	r := ripemd160.New()
	_, err = r.Write(h[:])
	if err != nil {
		panic(err)
	}
	b := append([]byte{byte(0x98)}, r.Sum(nil)...)
	h = sha3.Sum256(b)
	got := append(b, h[:4]...)
	fmt.Println("GOT")
	fmt.Println(base32.StdEncoding.EncodeToString(got[:]))
	fmt.Println(want)
	fmt.Println("WANT")
	fmt.Println("-----------------------------")
	c := nemgo.New(nemgo.WithNIS("159.203.1.94:7890", nemgo.Testnet))
	a, err := c.AccountData("TCQKEW2GQ25L44CGZJVQRGL4LGO7OR2OX43YYTKA")
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}
