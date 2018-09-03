package main

import (
	"fmt"
	"strings"

	"github.com/bigzhu/collision_cardano/conf"

	bip32 "github.com/tyler-smith/go-bip32"
	bip39 "github.com/tyler-smith/go-bip39"
)

func main() {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seedConf := conf.GetSeedConf()
	seed := bip39.NewSeed(mnemonic, seedConf.Seed)

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)
	splitMnemonic := strings.Fields(mnemonic)
	// fmt.Println(splitMnemonic, len(splitMnemonic))
	stringFiles := strings.Join(splitMnemonic, `","`)
	fmt.Println(`"` + stringFiles + `"`)
}
