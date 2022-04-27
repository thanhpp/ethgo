package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	if err := queryFullBlock(client, nil); err != nil {
		log.Fatal(err)
	}
}

// queryFullBlock if the blockNumber is nil -> query the latest block
func queryFullBlock(client *ethclient.Client, blockNumber *big.Int) error {
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return err
	}

	fmt.Printf("full block: %+v\n\n", block)

	listAllTransactions(block.Transactions())

	if err := getReceipt(client, block.Transactions()[0].Hash()); err != nil {
		return err
	}

	return nil
}

func listAllTransactions(txs types.Transactions) {
	for i := range txs {
		fmt.Printf("%d %+v\n\n", i, txs[i])
	}
}

func getSenderAddress(client *ethclient.Client, tx types.Transaction, baseFee *big.Int) error {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return err
	}

	msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), baseFee)
	if err != nil {
		return err
	}

	fmt.Println("sender address", msg.From().Hex())

	return nil
}

func getReceipt(client *ethclient.Client, txHash common.Hash) error {
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return err
	}

	fmt.Printf("receipt %+v", receipt)

	return nil
}
