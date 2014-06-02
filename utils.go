package mumgo

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
)

// Derives the current working directory of running program
func pwd() string {
	pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		fmt.Errorf("cannot get working directory: %s", err)
		os.Exit(1)
	}

	return pwd
}

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
