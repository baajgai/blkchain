package mod

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
)

// Transaction struct contains source/destination addresses and the value
type Transaction struct {
	ToAddress   string // Receiver address
	FromAddress string // Sender address
	Value       uint64 // Transaction value
	Signature   []byte // Signature
}

func NewTransaction(to_address, from_address string, value uint64) *Transaction {
	return &Transaction{
		ToAddress:   to_address,
		FromAddress: from_address,
		Value:       value,
	}
}

func (tx *Transaction) Hash() [32]byte {
	data, _ := json.Marshal(tx)

	return sha256.Sum256(data)
}

func (tx *Transaction) SignTransaction() {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	publicKey := &privateKey.PublicKey

	fmt.Println(" Pub key = ", publicKey)
	tx.Signature = nil
	hashedTx := tx.Hash()
	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, hashedTx[:])
	signature := append(r.Bytes(), s.Bytes()...)
	tx.Signature = signature

	fmt.Printf("Signed Transaction: %s\n", hex.EncodeToString(tx.Signature))

}

func (tx *Transaction) VerifyTransaction(pub *ecdsa.PublicKey) bool {

	// Save and clear signature to hash same data
	sign := tx.Signature
	tx.Signature = nil

	hashedTx := tx.Hash()
	tx.Signature = sign

	r := new(big.Int).SetBytes(sign[:len(sign)/2])
	s := new(big.Int).SetBytes(sign[len(sign)/2:])

	return ecdsa.Verify(pub, hashedTx[:], r, s)
}
