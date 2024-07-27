package bitfield

// Bitfield represents the pieces that a peer has.
type Bitfield []byte

// HasPiece returns true if the bitfield has the specified index.
func (bf Bitfield) HasPiece(index int) bool {
    byteIndex := index / 8
    bitIndex := index % 8
    return bf[byteIndex]&(1<<uint(7-bitIndex)) != 0
}

// SetPiece marks the piece at the specified index as being available.
func (bf Bitfield) SetPiece(index int) {
    byteIndex := index / 8
    bitIndex := index % 8
    bf[byteIndex] |= 1 << uint(7-bitIndex)
}
