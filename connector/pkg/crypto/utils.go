package crypto

import (
	"fmt"
	"golang.org/x/crypto/blake2s"
)

const HASHKEY string = "APP_BASE"

func Hash(bs []byte) ([]byte, error) {
	h, err := blake2s.New128([]byte(HASHKEY))
	if err != nil {
		return []byte{}, err
	}
	h.Write(bs)
	return h.Sum(nil), nil
}

func CheckHashEqual(b1 []byte, b2 []byte) error {
	if len(b1) != len(b2) {
		return fmt.Errorf("hashes of different lengths")
	}
	for i:=1; i<len(b1); i++ {
		if (b1[i] != b2[i]) {
			return fmt.Errorf("hashes are not equal")
		}
	}
	return nil
}