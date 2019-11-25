package network

import (
	"encoding/json"
	"../crypto"
	"errors"
	pq "github.com/jupp0r/go-priority-queue"
)

// Should abstract away the connections and verify that packets are correct coming in and out. Note that we do not deal
// with encryption here -- we can assume that the peer in the identity and TLS system we are working on will be the primitive

// This process should explicitly unbundle the packets coming in into their small VPN packets and all the packets leaving
// should be bundled into large packets.

type Peer struct {
	address []byte
	incoming chan NetworkPacket // TODO: Replace this with a peer connection wrapped by this
	outgoing chan NetworkPacket
	pendingPackets map[string]PaymentPacket
	pendingPacketsTOQueue pq.PriorityQueue
	nextVPNPacket VPNPacket
	nextVPNPacketSize int
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
	networkPkt := <-peer.incoming
	pkt, err := DeserNetworkPacket(networkPkt)
	if err != nil { return err }
	// PaymentPacket functionality
	if err := peer.checkValidPacket(pkt); err != nil { return err }
	if pkt.PacketType() == VPN {
		vpnPkt, ok := pkt.(*VPNPacket)
		if !ok { return errors.New("unable to convert peerPkt into VPNPacket") }
		for i := 0; i < len(vpnPkt.Packets); i++ {
			if err := peer.checkValidPacket(vpnPkt.Packets[i]); err != nil { return err }
			peer.core.RegisterIncomingPacket(vpnPkt.Packets[i])
		}
	} else {
		peer.core.RegisterIncomingPacket(pkt)
	}
	return nil
}

/*
Called by WriteToVPNPacket to check and see if the packets aggregated inside peer.nextVPNPacket should be flushed to the
network.
 */
func (peer *Peer) ShouldFlushVPNPacket() bool {
	if len(peer.nextVPNPacket.Packets) >= peer.nextVPNPacketSize {
		return true
	}
	return false
}

/*
Called by core to write an aggregated packet to this peer. Write to the outgoing channel, and that channel should handle
the actual processing
 */
func (peer *Peer) Write(packet Packet) error {
	if err := peer.checkValidPacket(packet); err != nil { return err }
	networkPkt, err := SerNetworkPacket(packet)
	if err != nil { return err }
	peer.outgoing <- networkPkt
	return nil
}

func (peer *Peer) WriteToVPNPacket(packet Packet) error {
	peer.nextVPNPacket.Packets = append(peer.nextVPNPacket.Packets, packet)
	if peer.ShouldFlushVPNPacket() {
		networkPkt, err := SerNetworkPacket(&peer.nextVPNPacket)
		if err != nil { return err }
		peer.outgoing <- networkPkt
		for i := 0; i < len(peer.nextVPNPacket.Packets); i++ {
			if peer.nextVPNPacket.Packets[i].PacketType() == PAYMENT {
				pkt, ok := peer.nextVPNPacket.Packets[i].(*PaymentPacket)
				if !ok { return errors.New("error during Packet to PaymentPacket conversion") }
				peer.pendingPackets[string(pkt.Header())] = *pkt
				peer.pendingPacketsTOQueue.Insert(pkt, float64(pkt.Timeout))
			}
		}
		peer.nextVPNPacket = VPNPacket{
			Packets: []Packet{},
		}
	}
	return nil
}