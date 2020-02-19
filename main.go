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

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`ropsten eth faucet
		example:
		
		$ eth-faucet 0x0581a31Bc9d1c030B004951D8bC520A07fcb3897`)
		return
	}

	address := os.Args[1]
	if address == "" {
		fmt.Println("address is empty")
		return
	}

	if !checkEthAddress(address) {
		fmt.Println("Invalid address")
		return
	}

	const api = "https://faucet.ropsten.be/donate/"
	resp, err := http.Get(api + address)
	if err != nil {
		fmt.Println("Request error", err)
		return
	}
	defer resp.Body.Close()
	_, _ = io.Copy(os.Stdout, resp.Body)
}
