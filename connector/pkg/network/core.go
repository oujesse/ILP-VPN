package network

import (
	"errors"
)

// This is the driver. It will feature a goroutine that will read off packets, aggregate them into large buckets, and then
// forward them out.

type Core struct {
	packetQueue chan Packet
}

/*
Submits this packet to the packetQueue.
 */
func (core *Core) RegisterIncomingPacket(pkt Packet) {
	core.packetQueue <- pkt
}

func (core *Core) handlePaymentPacket(pkt Packet) error {
	_, ok := pkt.(*PaymentPacket);
	if !ok { return errors.New("pkt cannot be converted into PaymentPacket") }
	return nil
}

/*
Runs as a goRoutine upon creation of Core. Reads in packets from the packetQueue and delegates calls to the routing_table
to know where to forward packets to as well as generate next hop packets. Finally, calls peer.write() to write this packet
given the peer.
 */
func (core *Core) PacketHandlerRoutine() error {
	select {
	case rcvPkt := <-core.packetQueue: {
		// Check if packet has reached destination or last hop.
		var err error
		switch rcvPkt.PacketType() {
		case PAYMENT: err = core.handlePaymentPacket(rcvPkt)
		}
		return err
	}
	}
}