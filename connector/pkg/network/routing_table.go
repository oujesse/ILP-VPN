package network

// This file should include all functionality that allows the generation of a NextHopPacket. Given its generation,
// the core will then send

var forwardPacketMap = map[string]string{}
var returnPacketMap = map[string]string{}
var packetCache = map[string]interface{}{}

var peers = make(map[string]Peer)

/*
This function generates the packet that will be sent to the next hop. It needs to do a few things:
- Create a new Packet
- Copy over the fields from the old Packet
- Replace necessary fields for the nextHopPacket (including nextHop)
- Alter the forwardPacketMap and returnPacketMap to keep track of packets received and sent
Note that how the actual packet forwarding is handled and the new packet is generated depends on the packet type
 */
func genNextHopPacket(packet Packet) Packet {

}

/*
Gets the next hop peer given a packet. This function can have wildly varying functionality, and it is this function
that should actually access peer tables and compute the next hop. This is especially important if the packet is an
accept or reject packet, as it must be mapped to the correct next hop.
 */
func getNextHop(packet Packet) Peer {

}

func getPeer(address []byte) Peer {

}