package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/alejoacosta74/chains/eth"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	var nonce uint64 = 0
	var gasPrice int64 = 1000000000
	var gasLimit uint64 = 21000
	var address string = "0x96216849c49358B10257cb55b28eA603c874b05E"
	var amount int64 = 1000000000000000000 // 1 ETH

	tx, _ := eth.NewEthereumTx(nonce, gasPrice, gasLimit, address, amount, nil)

	privKeyECDSA, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	privKey := hex.EncodeToString(crypto.FromECDSA(privKeyECDSA))
	fmt.Printf("Generated private key: %s\n", privKey)
	pubKey, err := eth.PrivToPubKey(privKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Public key: %s\n", pubKey)

	signedTx, err := eth.SignEthereumTx(tx, types.NewEIP155Signer(big.NewInt(1)), privKey)
	if err != nil {
		panic(err)
	}
	signedTxBytes, err := signedTx.MarshalJSON()
	if err != nil {
		panic(err)
	}
	logPretty("Signed transaction", signedTxBytes)

	encodedTx, err := eth.EncodeRawTx(signedTx)
	if err != nil {
		panic(err)
	}
	decodedTx, err := eth.DecodeRawTx(encodedTx)
	if err != nil {
		panic(err)
	}
	decodedTxBytes, err := decodedTx.MarshalJSON()
	if err != nil {
		panic(err)
	}
	logPretty("==> Decoded transaction", decodedTxBytes)

	result := eth.VerifyTxSignature(decodedTx, pubKey)
	fmt.Printf("\n==> Signature verified?: %t\n", result)

}

func logPretty(msg string, output []byte) {
	if len(output) > 0 {
		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, output, "", "    "); err != nil {
			fmt.Printf("Error decoding JSON: %v\n", err)
		} else {
			fmt.Printf("\n%s :\n%s\n", msg, prettyJSON.String())

		}
	}
}
