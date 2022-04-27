package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	// key pair generation
	private, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateBytes := crypto.FromECDSA(private)
	fmt.Println("private key", hexutil.Encode(privateBytes))

	public := private.Public()
	publicECSDA, ok := public.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("invalid public key")
	}

	publicBytes := crypto.FromECDSAPub(publicECSDA)
	fmt.Println("public key", hexutil.Encode(publicBytes))

	// generate wallet
	address := crypto.PubkeyToAddress(*publicECSDA).Hex() // using lib
	fmt.Println("address", address)

	addrHash := sha3.NewLegacyKeccak256()                               // manual
	addrHash.Write(publicBytes[1:])                                     // manual
	fmt.Println("address hash", hexutil.Encode(addrHash.Sum(nil)[12:])) //  take the last 40 characters (20 bytes) and prefix it with 0x
}
