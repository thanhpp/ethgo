package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/thanhpp/ethgo/pkg/explorerutil"
)

/*
A keystore is a file containing an encrypted wallet private key.
Keystores in go-ethereum can only contain one wallet key pair per file.
*/

func main() {
	createKeyStore("./tmp", "pass")

	f, err := explorerutil.LatestModFile("./tmp")
	if err != nil {
		log.Fatal(err)
	}

	importKeyStore(
		"./tmp1",
		fmt.Sprintf("./tmp/%s", f.Name()),
		"pass",
	)
}

func createKeyStore(dir, pass string) {
	ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(pass)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("created account", account.Address.Hex())
}

func importKeyStore(dir, file, pass string) {
	ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	account, err := ks.Import(fileBytes, pass, pass)
	if err != nil {
		log.Fatal("import key store: ", err)
	}

	fmt.Println("import account", account.Address.Hex())

	// remove to reuse the folder
	if err := os.RemoveAll(dir); err != nil {
		log.Fatal(err)
	}
}
