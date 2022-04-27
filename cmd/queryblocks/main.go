package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	_ = client

	if err := queryBlockHeader(client, nil); err != nil {
		log.Fatal(err)
	}

	if err := queryFullBlock(client, nil); err != nil {
		log.Fatal(err)
	}
}

// queryBlockHeader if the blockNumber is nil -> query the latest block
func queryBlockHeader(client *ethclient.Client, blockNumber *big.Int) error {
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		return err
	}

	fmt.Printf("header: %+v\n\n", header)

	return nil
}

// queryFullBlock if the blockNumber is nil -> query the latest block
func queryFullBlock(client *ethclient.Client, blockNumber *big.Int) error {
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return err
	}

	fmt.Printf("full block: %+v\n\n", block)

	return queryTransactionCount(client, block.Hash())
}

func queryTransactionCount(client *ethclient.Client, blockHash common.Hash) error {
	txCount, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		return err
	}

	fmt.Printf("transaction count: %+v\n\n", txCount)

	return nil
}
