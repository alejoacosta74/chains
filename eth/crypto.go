package eth

import (
	"crypto/ecdsa"
	"encoding/hex"
	"regexp"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

// PrivToPubKey converts a ECDSA private key to a Ethereum public key.
func PrivToPubKey(privKeyHex string) (string, error) {
	privKeyHex, err := sanitizePrivateKey(privKeyHex)
	if err != nil {
		return "", err
	}
	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return "", errors.Wrapf(err, "error converting private key to ECDSA: %s", privKeyHex)
	}
	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}
	pubKeyBytes := crypto.FromECDSAPub(pubKeyECDSA)
	return hex.EncodeToString(pubKeyBytes), nil
}

// PubKeyToAddress converts a public key to an Ethereum address.
func PubKeyToAddress(pubKeyHex string) (string, error) {
	pubKeyBytes, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return "", err
	}
	pubKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return "", err
	}
	address := crypto.PubkeyToAddress(*pubKey)
	return address.Hex(), nil
}

func verifyHexAddress(address string) bool {
	if len(address) > 2 && address[:2] == "0x" {
		address = address[2:]
	}
	if len(address) != 40 {
		return false
	}
	re := regexp.MustCompile(`[0-9a-fA-F]{40}$`)
	return re.MatchString(address)
}

func sanitizePrivateKey(privKeyHex string) (string, error) {
	if len(privKeyHex) > 2 && privKeyHex[:2] == "0x" {
		privKeyHex = privKeyHex[2:]
	}
	if len(privKeyHex) != 64 {
		return "", errors.New("invalid private key length: " + privKeyHex)
	}
	re := regexp.MustCompile(`[0-9a-fA-F]{64}$`)
	if re.MatchString(privKeyHex) {
		return privKeyHex, nil
	}
	return "", errors.New("invalid private key: " + privKeyHex)
}
