package convert

import "encoding/binary"

func ConvertIDToByte(id uint64) []byte {
	idBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(idBytes, id)
	return idBytes
}
