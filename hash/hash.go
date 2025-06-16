package hash

type Hashable interface {
	Bytes() []byte
	Hash() []byte
}
