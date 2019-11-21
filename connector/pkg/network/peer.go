package network

import (
	"../crypto"
	"encoding/json"
	"github.com/jupp0r/go-priority-queue"
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
	payPktHash, err := crypto.Hash(pktWrapper.SerPacket)
	if err != nil { return err }
	err = crypto.CheckHashEqual(payPktHash, pktWrapper.Header)
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
	if err := peer.checkValidPacket(pkt); err != nil { return err }
	if pkt.Type == VPN {
		var vpnPacket VPNPacket
		err := json.Unmarshal(pkt.SerPacket, vpnPacket)
		if err != nil { return err }
		for i := 0; i < len(vpnPacket.Packets); i++ {
			if err := peer.checkValidPacket(vpnPacket.Packets[i]); err != nil { return err }
			peer.core.RegisterIncomingPacket(pkt)
		}
	} else {
		peer.core.RegisterIncomingPacket(pkt)
	}
	return nil
}

/*
Called by core to write an aggregated packet to this peer. Write to the outgoing channel, and that channel should handle
the actual processingi
 */
func (peer *Peer) write(packet Packet) {

}