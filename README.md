Ropsten Ethereum Faucet Tool 
---

[![Build Status](https://travis-ci.org/islishude/eth-faucet.svg?branch=master)](https://travis-ci.org/islishude/eth-faucet)

## Usage

```console
$ GO111MODULE=on go get -u github.com/islishude/eth-faucet
$ eth-faucet 0x0581a31Bc9d1c030B004951D8bC520A07fcb3897 | jq .
{
  "address": "0x0581a31bc9d1c030b004951d8bc520a07fcb3897",
  "txhash": "0x7d4e60e582f2af443627dc921446b3bbe830d984e177e57407bf4f3c9b9702dc",
  "amount": 1000000000000000000
}
```

## Response fields

- paydate the unix timestamp when the transaction will be executed. Depends on the current length of the queue
- address the address where the payment will be done
- amount the amount in Wei that will be transferred
- txhash transaction hash : if the queue is empty, you will immediately receive the transaction hash - if the queue is not empty - your request will be queued until paydate and the txhash field will be empty.