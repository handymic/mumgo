package mumgo

import (
	"encoding/binary"
)

func uint16tbs(i uint16) []byte {
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, i)
	return bs
}

func uint32tbs(i uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, i)
	return bs
}
