package eth

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
)

const ()

func TestTransactions(t *testing.T) {
	privKey := "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
	var nonce uint64 = 0
	var gasPrice int64 = 1000000000
	var gasLimit uint64 = 21000
	var address string = "0x96216849c49358B10257cb55b28eA603c874b05E"
	var amount int64 = 1000000000000000000 // 1 ETH

	pubKey, err := PrivToPubKey(privKey)
	handleFatalError(t, err)

	tx, _ := NewEthereumTx(nonce, gasPrice, gasLimit, address, amount, nil)

	tests := []struct {
		name   string
		signer types.Signer
	}{
		{"HomesteadSigner", types.HomesteadSigner{}},
		{"EIP155Signer", types.NewEIP155Signer(big.NewInt(1))},
	}

	for _, tt := range tests {

		signedTx, err := SignEthereumTx(tx, tt.signer, privKey)
		handleFatalError(t, err)

		encodedTx, err := EncodeRawTx(signedTx)
		handleFatalError(t, err)

		t.Run("DecodeRawTx("+tt.name+")", func(t *testing.T) {
			decodedTx, err := DecodeRawTx(encodedTx)
			if err != nil {
				t.Fatal(err)
			}
			want, err := signedTx.MarshalJSON()
			handleFatalError(t, err)
			got, err := decodedTx.MarshalJSON()
			handleFatalError(t, err)
			if string(want) != string(got) {
				t.Errorf("got %s, want %s", string(got), string(want))
			}
		})

		t.Run("VerifySignature("+tt.name+")", func(t *testing.T) {
			result := VerifyTxSignature(signedTx, pubKey)
			if !result {
				t.Errorf("got %t, want %t", result, true)
			}
		})
	}

}
