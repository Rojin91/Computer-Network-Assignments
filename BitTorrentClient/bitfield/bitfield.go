package bitfield

type Bitfield []byte

func (bf Bitfield) HasPiece(index int) bool {
	byteIndex := index / 8
	bitIndex := index % 8
	return bf[byteIndex]&(1<<(7-bitIndex)) != 0
}

func (bf Bitfield) SetPiece(index int) {
	byteIndex := index / 8
	bitIndex := index % 8
	bf[byteIndex] |= 1 << (7 - bitIndex)
}
