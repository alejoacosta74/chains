package eth

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
)

const ()

func TestTransactions(t *testing.T) {
	privKey := "0x85cbc7b1adfe877051d746c3996a01c2bc3e7a6988490439b1f4b4c2b465322d"
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
