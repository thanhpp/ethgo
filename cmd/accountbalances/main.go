package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x829BD824B016326A401d083B33D092293333A830")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance", balance)

	blockNo := big.NewInt(14664874)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balanceAt", balanceAt)

	fBalance, _ := new(big.Float).SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("ethValue", ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pendingBalance", pendingBalance)
}
