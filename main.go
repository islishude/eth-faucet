package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func checkEthAddress(address string) bool {
	return regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`).MatchString(address)
}

const api = "https://faucet.ropsten.be/donate/"

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`ropsten ethereum faucet

example:

$ eth-faucet 0x0581a31Bc9d1c030B004951D8bC520A07fcb3897`)
		return
	}

	address := os.Args[1]
	if !checkEthAddress(address) {
		fmt.Println("Invalid address")
		return
	}

	resp, err := http.Get(api + address)
	if err != nil {
		fmt.Println("Request error", err)
		return
	}
	defer resp.Body.Close()
	_, _ = io.Copy(os.Stdout, resp.Body)
}
