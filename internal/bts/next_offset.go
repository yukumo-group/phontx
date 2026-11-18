package bts

// NextOffset gets nexxt offset that can be devided by 4
func NextOffset(
	offset uint32,
	length uint32,
) uint32 {
	return ((offset + length + 16 + 3) / 4) * 4
}
