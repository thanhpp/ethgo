package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// starts with 0x
	// follows by 40 characters(a-f OR A-F OR 0-9) (hexa)
	re, err := regexp.Compile("^0x[0-9a-fA-F]{40}$")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("address check", re.MatchString("0xdefa4e8a7bcba345f687a2f1456f5edd9ce97202"))

	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	accountCheck, err := isAccount(client, "0xdefa4e8a7bcba345f687a2f1456f5edd9ce97202")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("account check", accountCheck)
}

func isAccount(client *ethclient.Client, addressStr string) (bool, error) {
	address := common.HexToAddress(addressStr)

	byteCode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		return false, err
	}

	// When there's no bytecode at the address
	// then we know that it's not a smart contract and it's a standard ethereum account
	if len(byteCode) == 0 {
		return true, nil
	}

	return false, nil
}
