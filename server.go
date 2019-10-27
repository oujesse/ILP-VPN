package main

import "log"

var paymentPackets = map[uint64]PaymentPacket {}
var nextHopPaymentPackets = map[uint64]PaymentPacket {}
var pki = map[string]string {}

type PaymentPacket struct {
	ID uint64
	nextHopID uint64
	sender string
	ledger string
	amount float64
	receiver string
	numHops uint64
	timestamp int64
}

type RejectionPacket struct {
	ID uint64
	paymentPacketID uint64
}

func verifyPacket(packet PaymentPacket, signature string, publicKey string) bool {
	return true
}

func signPacket(packet PaymentPacket, privKey string) string {
	return ""
}

func acceptablePacketConditions(packet PaymentPacket) bool {
	return true
}

func receivePaymentPacket(packet PaymentPacket, signature string) {
	if _, exists := paymentPackets[packet.ID]; exists {
		log.Fatal("Received Packet is a duplicate")
	} else if !verifyPacket(packet, signature, pki[packet.sender]){
		log.Fatal("Received Packet Sender's publicKey does not verify signature")
	} else if !acceptablePacketConditions(packet) {
		log.Fatal("Received Packet does not fulfill acceptable conditions")
	} else {
		paymentPackets[packet.ID] = packet
		var nextHopPacket = PaymentPacket {
			ID: packet.nextHopID,
			// nextHopID
			sender: packet.sender,
			ledger: packet.ledger,
			amount: packet.amount,
			receiver: packet.receiver,
			numHops: packet.numHops - 1,
			// timestamp
		}
		nextHopPaymentPackets[nextHopPacket.ID] = nextHopPacket
		VPNPacketHandler(nextHopPacket, signPacket(nextHopPacket, ""))
	}
}

func VPNPacketHandler(packet PaymentPacket, signature string) {

}

func main() {

}