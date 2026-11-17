package bts

// NextOffset gets nexxt offset that can be devided by 4
func NextOffset(
	offset int32,
	length int32,
) int32 {
	return ((offset + length + 3) / 4) * 4
}
