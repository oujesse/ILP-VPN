package crypto

import (
	"bytes"
	"testing"
	"../network"
)

func TestPacketHash(t *testing.T) {
	PacketHash([]byte("Arbitrary byte string"))
}

/*func TestMarshalPacketHash(t *testing.T) {
	packetHash := PacketHash([]byte("Arbitrary byte string"))
	bytestr := MarshalPacketHash(packetHash)
	print(bytestr)
}*/

func TestSerializeDeserializePaymentPacket(t *testing.T) {
	// Serialize
	pp := network.PaymentPacket{
		Header: []byte("Header Test"),
	}
	byteStr, err := SerializePacket(pp)
	if err != nil { t.Fatal(err) }
	expected := []byte(`{"Header":"SGVhZGVyIFRlc3Q=","Receiver":null,"Commit":null,"Ledger":"","Amount":0,"NumHops":0,"Timestamp":0,"Timeout":0}`)
	if bytes.Compare(byteStr, expected) != 0 {
		t.Fatal("byteStr not correct")
	}
	// Deserialize
	pp2 := network.PaymentPacket{}
	err = DeserializePacket(byteStr, &pp2)
	if err != nil { t.Fatal(err) }
	if string(pp2.Header) != "Header Test" {
		t.Fatal("Deserialize returns incorrect result")
	}
}
