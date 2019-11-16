package crypto

import (
	"encoding/json"
	"golang.org/x/crypto/blake2s"
	"hash"
)

func PacketHash(serializedPacket []byte) (hash.Hash, error) {
	return blake2s.New128(serializedPacket)
}

func SerializePacket(packet interface{}) ([]byte, error) {
	return json.Marshal(packet)
}

func DeserializePacket(packetBytes []byte, packetReference interface{}) error {
	return json.Unmarshal(packetBytes, packetReference)
}
