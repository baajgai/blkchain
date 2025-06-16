package mod

import "crypto/sha256"

// Output struct contains destination address and the value
type Output struct {
	ToAddress string
	Value     uint64
}

func (o Output) Bytes() []byte {
	var bytes []byte

	// Add address as bytes
	addr := []byte(o.ToAddress)
	bytes = append(bytes, addr[:]...)

	// Add value as bytes
	val_bytes := U64Bytes(o.Value)
	bytes = append(bytes, val_bytes[:]...)

	return bytes
}

// Transaction struct contains inputs and outputs
type Transaction struct {
	Inputs  []Output
	Outputs []Output
}

func (t *Transaction) InputValue() uint64 {

	var sum uint64

	for _, val := range t.Inputs {

		sum = sum + val.Value
	}

	return sum

}

func (t Transaction) OutputValue() uint64 {

	var sum uint64

	for _, val := range t.Outputs {

		sum = sum + val.Value
	}

	return sum

}

func (t Transaction) InputHash() *HashSet {

	var checksum [32]byte
	hashSet := NewHashSet(len(t.Inputs))

	for _, input := range t.Inputs {

		checksum = sha256.Sum256(input.Bytes())
		hashSet.Add(checksum)

	}
	return hashSet

}

func (t Transaction) OutputHash() *HashSet {

	var checksum [32]byte
	hashSet := NewHashSet(len(t.Outputs))

	for _, output := range t.Outputs {

		checksum = sha256.Sum256(output.Bytes())
		hashSet.Add(checksum)

	}

	return hashSet

}

func (t Transaction) Bytes() []byte {

	var bytes []byte

	for _, input := range t.Inputs {

		bytes = append(bytes, input.Bytes()...)
	}

	for _, output := range t.Outputs {

		bytes = append(bytes, output.Bytes()...)
	}

	return bytes

}

func (t *Transaction) IsCoinbase() bool {
	return len(t.Inputs) == 0
}
