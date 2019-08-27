package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

const api = "https://faucet.ropsten.be/donate/"

var ethAddrRegExp = regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)

func checkEthAddress(address string) bool {
	return ethAddrRegExp.MatchString(address)
}

var (
	address string
	help    bool
)

func main() {
	flag.StringVar(&address, "address", "", "Your ethereum address")
	flag.BoolVar(&help, "help", false, "Print help text")
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	if address == "" {
		fmt.Println("address is empty")
		return
	}

	if !checkEthAddress(address) {
		fmt.Println("Invalid address")
		return
	}

	url := fmt.Sprintf("%s%s", api, address)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Request error", err)
		return
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Response error", err)
		return
	}

	var out bytes.Buffer
	if err := json.Indent(&out, res, "", "  "); err != nil {
		fmt.Println("Format response error", err)
		return
	}
	out.WriteTo(os.Stdout)
}
