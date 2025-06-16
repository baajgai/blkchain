package mod

func U32Bytes(u uint32) [4]byte {

	return [4]byte{
		byte(u >> 8 * 0),
		byte(u >> 8 * 1),
		byte(u >> 8 * 2),
		byte(u >> 8 * 3),
	}

}

func U64Bytes(u uint64) [8]byte {

	return [8]byte{
		byte(u >> 8 * 0),
		byte(u >> 8 * 1),
		byte(u >> 8 * 2),
		byte(u >> 8 * 3),

		byte(u >> 8 * 4),
		byte(u >> 8 * 5),
		byte(u >> 8 * 6),
		byte(u >> 8 * 7),
	}

}

func DifficultyBytesAsU128(v []byte) uint64 {
	if len(v) < 32 {
		panic("input slice must have at least 32 bytes")
	}
	return (uint64(v[31]) << (0xf * 8)) |
		(uint64(v[30]) << (0xe * 8)) |
		(uint64(v[29]) << (0xd * 8)) |
		(uint64(v[28]) << (0xc * 8)) |
		(uint64(v[27]) << (0xb * 8)) |
		(uint64(v[26]) << (0xa * 8)) |
		(uint64(v[25]) << (0x9 * 8)) |
		(uint64(v[24]) << (0x8 * 8)) |
		(uint64(v[23]) << (0x7 * 8)) |
		(uint64(v[22]) << (0x6 * 8)) |
		(uint64(v[21]) << (0x5 * 8)) |
		(uint64(v[20]) << (0x4 * 8)) |
		(uint64(v[19]) << (0x3 * 8)) |
		(uint64(v[18]) << (0x2 * 8)) |
		(uint64(v[17]) << (0x1 * 8)) |
		(uint64(v[16]) << (0x0 * 8))
}
