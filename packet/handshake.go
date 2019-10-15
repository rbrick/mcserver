package packet

import "github.com/rbrick/mcserver/util"

type HandshakePacket struct {
	ProtoVersion *util.Uvarint
	ServerAddr   string
	ServerPort   uint16
	NextState    *util.Uvarint
}

func (HandshakePacket) ID() int {
	return 0x00
}
