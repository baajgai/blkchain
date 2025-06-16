package mod

type HashSet struct {
	data [][32]byte
}

func NewHashSet(capacity int) *HashSet {
	return &HashSet{
		data: make([][32]byte, 0, capacity),
	}
}

// Add a new element to the hashset
func (h *HashSet) Add(key [32]byte) {
	for _, v := range h.data {
		if v == key {
			return
		}
	}

	h.data = append(h.data, key)
}

// Contains checks if 32 byte data exists
func (h *HashSet) Contains(key [32]byte) bool {
	for _, v := range h.data {
		if v == key {
			return true
		}
	}

	return false
}

// Returns total number of hash elements in the hashset
func (h *HashSet) Len() int {
	return len(h.data)
}
