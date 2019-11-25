package network

import "testing"

func setup() Peer {
	peer := Peer {
	}
	peer.incoming <- NetworkPacket{}
	return peer
}

func TestPeer(t *testing.T){
	setup()
}
