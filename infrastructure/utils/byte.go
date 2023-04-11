package utils

func Byte32ToByte(data32 [32]byte) []byte {
	dataByte := make([]byte, len(data32))
	for i, b := range data32 {
		dataByte[i] = b
	}

	return dataByte
}
