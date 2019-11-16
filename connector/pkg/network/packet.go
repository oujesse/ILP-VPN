package network

// This file should define all the packets used. Each packet type is wrapped by the Packet struct, and upon sending,
// should be able to be turned into a generic Packet.


type Packet struct {
	packetType string
	nextHop []byte
	serPacket []byte
}

type PaymentPacket struct {
	Header []byte
	Receiver []byte
	Commit []byte
	Ledger string
	Amount float64
	NumHops uint8
	Timestamp int64
	Timeout int64
}

type VPNPacket struct {
	packets []PaymentPacket
}

type AcceptPacket struct {
	Header []byte
	AcceptedPacketHeader []byte
	Precommit []byte
}

type RejectionPacket struct {
	header []byte
	rejectedPacketHeader []byte
}