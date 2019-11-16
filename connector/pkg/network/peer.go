package network

import (
	"../crypto"
	"fmt"
	"github.com/bgadrian/data-structures/priorityqueue"
	"hash"
	"time"
)

// Should abstract away the connections and verify that packets are correct coming in and out. Note that we do not deal
// with encryption here -- we can assume that the peer in the identity and TLS system we are working on will be the primitive

// This process should explicitly unbundle the packets coming in into their small VPN packets and all the packets leaving
// should be bundled into large packets.

type Peer struct {
	incoming chan Packet // TODO: Replace this with a peer connection wrapped by this
	outgoing chan Packet
	receivedAcceptedPaymentPacketsQueue priorityqueue.HierarchicalHeap // PaymentPackets ordered by timeout
	receivedAcceptedPaymentPackets map[hash.Hash]PaymentPacket
	sentPendingPaymentPackets map[hash.Hash]PaymentPacket
	PaymentPacketFeePercent float64
}

func (peer * Peer) acceptablePaymentPacketConditions(paymentPacket PaymentPacket) bool {
	return true
}

func newPeer() {

}

func (peer *Peer) read() {
	// Server functionality
	packetWrapper := <-peer.incoming
	// PaymentPacket functionality
	if packetWrapper.packetType == "Payment" {
		var paymentPacket PaymentPacket
		err := crypto.DeserializePacket(packetWrapper.serPacket, &paymentPacket)
		if err != nil { fmt.Errorf("Error in Deserialization", err) }
		paymentPacketHash, err := crypto.PacketHash(packetWrapper.serPacket)
		if err != nil { fmt.Errorf("Error in hashing", err) }
		if _, exists := peer.receivedAcceptedPaymentPackets[paymentPacketHash]; exists {
			fmt.Errorf("packet hash already exists")
		} else if !peer.acceptablePaymentPacketConditions(paymentPacket) {
			// TODO: Send Rejection Packet
			fmt.Errorf("Received PaymentPacket does not fulfill Peer's acceptable conditions")
		} else {
			peer.receivedAcceptedPaymentPackets[paymentPacketHash] = paymentPacket
			err := peer.receivedAcceptedPaymentPacketsQueue.Enqueue(paymentPacket, int(paymentPacket.Timeout))
			if err != nil { fmt.Errorf("Failed to enqueue PaymentPacket onto Queue", err) }
			newPaymentPacket := PaymentPacket{
				// TODO: Add Header
				// TODO: Add Receiver
				// TODO: Add Commit
				Ledger: paymentPacket.Ledger,
				Amount: paymentPacket.Amount * (1 - peer.PaymentPacketFeePercent),
				NumHops: paymentPacket.NumHops - 1,
				Timestamp: time.Now().Unix(),
				Timeout: paymentPacket.Timeout,
			}
			if newPaymentPacket.NumHops != 0 {
				// TODO: Send to VPN Packet Handler
				// TODO: Add to sentPendingPaymentPackets
			} else {

			}
			/*acceptPacket := AcceptPacket{
				// TODO: Add Header
				AcceptedPacketHeader: paymentPacket.Header,
				// TODO: Add Precommit
			}*/
			// TODO: Send Accept Packet backward
		}
	}
}

func (peer *Peer) write(packet Packet) {

}
