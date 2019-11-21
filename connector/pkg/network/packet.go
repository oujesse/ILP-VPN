package network

import (
	"encoding/json"
	"fmt"
)

// This file should define all the packets used.
// To add a new packet, please inherit from _Packet and add the deserialization logic to DeserNetworkPacket

type PacketType int

// GOTCHA: THE ORDER BELOW CANNOT BE CHANGED
const (
	PAYMENT PacketType = iota // 0
	VPN
	ACCEPT
	REJECT
)

// Wraps packets for peers and returns a packet that implements Packet interface
type NetworkPacket struct {
	PacketType
	SerPacket []byte
}

type Packet interface {
	PacketType() PacketType
	Header() []byte
	NextHop() []byte
	SetHeader(b []byte)
}

func SerNetworkPacket(pkt Packet) (NetworkPacket, error) {
	var peerPkt NetworkPacket
	pktType := pkt.PacketType()
	serPkt, err := json.Marshal(pkt)
	if err == nil {
		peerPkt = NetworkPacket{
			PacketType: pktType,
			SerPacket: serPkt,
		}
	}
	return peerPkt, err
}

func DeserNetworkPacket(pkt NetworkPacket) (Packet, error) {
	switch pkt.PacketType {
	case PAYMENT: {
		var retPkt PaymentPacket
		err := json.Unmarshal(pkt.SerPacket, retPkt)
		return &retPkt, err
	}
	case VPN: {
		var retPkt VPNPacket
		err := json.Unmarshal(pkt.SerPacket, retPkt)
		return &retPkt, err
	}
	case ACCEPT: {
		var retPkt AcceptPacket
		err := json.Unmarshal(pkt.SerPacket, retPkt)
		return &retPkt, err
	}
	case REJECT: {
		var retPkt RejectPacket
		err := json.Unmarshal(pkt.SerPacket, retPkt)
		return &retPkt, err
	}
	default: {
		return &PaymentPacket{}, fmt.Errorf("packet type not found")
	}
	}
}

// packet that allows
type _Packet struct {
	packetType PacketType
	nextHop []byte // When sending, this is the next hop. When receiving, this should be current node
	header []byte
}

func (p *_Packet) PacketType() PacketType {
	return p.packetType
}

func (p *_Packet) Header() []byte {
	return p.header
}

func (p *_Packet) NextHop() []byte {
	return p.nextHop
}

func (p *_Packet) SetHeader(b []byte) {
	p.header = b
}

type PaymentPacket struct {
	_Packet
	Receiver []byte
	Commit []byte
	Ledger string
	Amount float64
	NumHops uint8
	Timestamp int64
	Timeout int64
}

type VPNPacket struct {
	_Packet
	Packets []Packet
}

type AcceptPacket struct {
	_Packet
	AcceptedPacketHeader []byte
	Precommit []byte
}

type RejectPacket struct {
	_Packet
	RejectedPacketHeader []byte
}

func (pkt *PaymentPacket) Copy() PaymentPacket {
	newPkt := *pkt
	return newPkt
}