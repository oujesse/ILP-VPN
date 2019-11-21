package network

// This file should define all the packets used. Each packet type is wrapped by the Packet struct, and upon sending,
// should be able to be turned into a generic Packet.

type PacketType int

// GOTCHA: THE ORDER BELOW CANNOT BE CHANGED
const (
	PAYMENT PacketType = iota // 0
	VPN
	ACCEPT
	REJECT
)

// Wrapper Packet
type Packet struct {
	Type PacketType
	NextHop []byte // When sending, this is the next hop. When receiving, this should be current node
	Header []byte
	SerPacket []byte
}

type PaymentPacket struct {
	Receiver []byte
	Commit []byte
	Ledger string
	Amount float64
	NumHops uint8
	Timestamp int64
	Timeout int64
}

type VPNPacket struct {
	Packets []Packet
}

type AcceptPacket struct {
	AcceptedPacketHeader []byte
	Precommit []byte
}

type RejectionPacket struct {
	RejectedPacketHeader []byte
}

func (pkt *PaymentPacket) Copy() PaymentPacket {
	newPkt := *pkt
	return newPkt
}