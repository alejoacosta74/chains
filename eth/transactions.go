package eth

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
)

// NewEthereumTx() creates a new unsigned ethereum transaction
// with the given parameters.
// An error is returned if the address is invalid.
func NewEthereumTx(nonce uint64, gasPrice int64, gasLimit uint64, address string, amount int64, data []byte) (*types.Transaction, error) {
	v := verifyHexAddress(address)
	if !v {
		return nil, errors.New("invalid address: " + address)
	}
	gasPriceBig := big.NewInt(gasPrice)
	toAddr := common.HexToAddress(address)
	amountBig := big.NewInt(amount)
	tx := types.NewTransaction(nonce, toAddr, amountBig, gasLimit, gasPriceBig, data)
	return tx, nil
}

// SignEthereumTx() signs an ethereum transaction with a private key.
//
// Params:
//   - tx: the transaction to sign
//   - signer: the signer to use (e.g. types.HomesteadSigner{})
//   - privKeyHex: the private key to sign the transaction with
func SignEthereumTx(tx *types.Transaction, signer types.Signer, privKeyHex string) (*types.Transaction, error) {
	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return nil, err
	}
	signedTx, err := types.SignTx(tx, signer, privKey)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}

// DecodeRawTX decodes a raw ethereum RLP encoded transaction
// and returns a go-ethereum types.Transaction
func DecodeRawTx(rawtx string) (*types.Transaction, error) {

	rawtx = strings.TrimPrefix(rawtx, "0x")
	rawtxBytes, err := hex.DecodeString(rawtx)
	if err != nil {
		return nil, err
	}

	var tx = &types.Transaction{}

	err = rlp.DecodeBytes(rawtxBytes, tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// EncodeRawTx encodes a go-ethereum types.Transaction
// and returns a raw ethereum RLP encoded transaction
func EncodeRawTx(tx *types.Transaction) (string, error) {

	var buf bytes.Buffer
	if err := rlp.Encode(&buf, tx); err != nil {
		return "", err
	}

	return hex.EncodeToString(buf.Bytes()), nil
}

// RecoverFromAddress returns the signer address from a signed ethereum transaction
func RecoverFromAddress(tx *types.Transaction) (*common.Address, error) {
	recoveredPubKey, err := RecoverPublicKey(tx)
	if err != nil {
		return nil, errors.Wrap(err, "error recovering pubkey")
	}
	recoveredAddress := crypto.PubkeyToAddress(*recoveredPubKey)

	return &recoveredAddress, nil

}

// RecoverPublicKey returns the signer public key from a signed ethereum transaction
func RecoverPublicKey(tx *types.Transaction) (*ecdsa.PublicKey, error) {
	signature, signer := getSignatureAndSigner(tx)
	recoveredPubKey, err := crypto.SigToPub(signer.Hash(tx).Bytes(), signature)
	if err != nil {
		return nil, errors.Wrap(err, "error recovering pubkey")
	}
	return recoveredPubKey, nil
}

// VerifyTxSignature() verifies the signature of a signed ethereum transaction
// against a public key.
func VerifyTxSignature(tx *types.Transaction, pubKeyHex string) bool {

	//1. Get the signature and signer from the signed transaction
	signature, signer := getSignatureAndSigner(tx)

	//2. Recreate the raw transaction from the signed transaction
	var rawTx = types.NewTransaction(tx.Nonce(), *tx.To(), tx.Value(), tx.Gas(), tx.GasPrice(), tx.Data())

	//3. RLP encode the raw transaction
	var buf bytes.Buffer
	if err := rlp.Encode(&buf, rawTx); err != nil {
		return false
	}
	//4. Hash the raw transaction
	rawTxHashed := signer.Hash(rawTx)
	digest := rawTxHashed.Bytes()

	//6. Recover the public key from the signature and message digest
	recoveredPubKey, err := crypto.SigToPub(digest, signature)
	if err != nil {
		return false
	}

	//7. Convert the recovered public key to a hex string
	recoveredPubKeyHex := hex.EncodeToString(crypto.FromECDSAPub(recoveredPubKey))

	//8. Compare the recovered public key hex with the expected public key hex
	return recoveredPubKeyHex == pubKeyHex
}

// getSignatureAndSigner returns the signature and signer from
// a signed ethereum transaction
func getSignatureAndSigner(tx *types.Transaction) (signature []byte, signer types.Signer) {
	v, r, s := tx.RawSignatureValues()

	signature = make([]byte, 65)
	copy(signature[32-len(r.Bytes()):32], r.Bytes())
	copy(signature[64-len(s.Bytes()):64], s.Bytes())

	if tx.Protected() {
		signer = types.NewEIP155Signer(tx.ChainId())
		signature[64] = byte(v.Uint64() - 35 - 2*tx.ChainId().Uint64())
	} else {
		signer = types.HomesteadSigner{}
		signature[64] = byte(v.Uint64() - 27)
	}

	return signature, signer
}
