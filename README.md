# chains

This repository is meant for educational purposes. It contains a collection of golang functions that can be used to implement different operations related to blockchain.

The code in this repo is explained in detailed in the following article: [https://medium.com/@alejoacosta2020/signing-and-verifying-ethereum-transactions-with-golang-a9cdc4061fbe](https://medium.com/@alejoacosta2020/signing-and-verifying-ethereum-transactions-with-golang-a9cdc4061fbe)


## How to use
Running the main.go file will execute the following functions:
- Create a sample ethereum transaction
- Generate a random private key and sign the transaction
- Encoded the signed transaction to RLP
- Decode the RLP encoded transaction
- Verify the signature of the transaction
  
```bash
❯ go run main.go
Generated private key: b6ead5c2d7fb9f29fbc58524cfda0239b1fc1d188df68ca176336fe6d8c21ef8
Public key: 04346630cc30cb8ecd0fd144c91a895679af6dfc47f698079186ec347eecec3d4fe84e48a5015247e0d3b35f78d9e55ffb45d5e4caadc3bb826bc2b00bd257481e

Signed transaction :
{
    "type": "0x0",
    "nonce": "0x0",
    "gasPrice": "0x3b9aca00",
    "maxPriorityFeePerGas": null,
    "maxFeePerGas": null,
    "gas": "0x5208",
    "value": "0xde0b6b3a7640000",
    "input": "0x",
    "v": "0x25",
    "r": "0x2354c0dab305bbe296eaa7927fff5179fe6aa6125c9c77fab93bb9efa02ca140",
    "s": "0x125b7b232674e0e85f03ab7dd7e58e2405d090afb1f70ebd39585f161a0ef0e6",
    "to": "0x96216849c49358b10257cb55b28ea603c874b05e",
    "hash": "0xb12e3223559cfef70122a1fc7077ae08037ff3a01937cb285a2582ab662e5b8a"
}

==> Decoded transaction :
{
    "type": "0x0",
    "nonce": "0x0",
    "gasPrice": "0x3b9aca00",
    "maxPriorityFeePerGas": null,
    "maxFeePerGas": null,
    "gas": "0x5208",
    "value": "0xde0b6b3a7640000",
    "input": "0x",
    "v": "0x25",
    "r": "0x2354c0dab305bbe296eaa7927fff5179fe6aa6125c9c77fab93bb9efa02ca140",
    "s": "0x125b7b232674e0e85f03ab7dd7e58e2405d090afb1f70ebd39585f161a0ef0e6",
    "to": "0x96216849c49358b10257cb55b28ea603c874b05e",
    "hash": "0xb12e3223559cfef70122a1fc7077ae08037ff3a01937cb285a2582ab662e5b8a"
}

==> Signature verified?: true
```
