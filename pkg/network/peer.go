package network

import (
	"encoding/json"
	"github.com/exasys/connector/pkg/crypto"
	pq "github.com/jupp0r/go-priority-queue"
)

// Should abstract away the connections and verify that packets are correct coming in and out. Note that we do not deal
// with encryption here -- we can assume that the peer in the identity and TLS system we are working on will be the primitive

// This process should explicitly unbundle the packets coming in into their small VPN packets and all the packets leaving
// should be bundled into large packets.

type Peer struct {
	incoming chan Packet // TODO: Replace this with a peer connection wrapped by this
	outgoing chan Packet
	pendingPackets map[string]PaymentPacket
	pendingPacketsTOQueue pq.PriorityQueue
	core *Core
}

/*
Constructor
When a new peer is created it should insert itself into the routing table, and upon destruction it should be removed
from the routing table.
 */
func newPeer() {

}

func (peer *Peer) checkValidPacket(pktWrapper Packet) error {
	tempHeader := pktWrapper.Header()
	pktWrapper.SetHeader([]byte{})
	ser, err := json.Marshal(pktWrapper)
	pktWrapper.SetHeader(tempHeader)

	if err != nil { return err }
	payPktHash, err := crypto.Hash(ser)
	if err != nil { return err }
	err = crypto.CheckHashEqual(payPktHash, pktWrapper.Header())
	if err != nil { return err }
	return nil
}

/*
Takes an aggregated packet and submits the split, inner VPN packets to core. Does this by running a coroutine reading
incoming and taking those packets, splitting them, and submitting them to core.
 */
func (peer *Peer) Read() error {
	// Server functionality
	pkt := <-peer.incoming
	// PaymentPacket functionality
	if pkt.PacketType() == VPN {

	} else {
		if err := peer.checkValidPacket(pkt); err != nil { return err }
		peer.core.RegisterIncomingPacket(pkt)
	}
	return nil
}

/*
Called by core to write an aggregated packet to this peer. Write to the outgoing channel, and that channel should handle
the actual processing
 */
func (peer *Peer) write(packet Packet) {

}