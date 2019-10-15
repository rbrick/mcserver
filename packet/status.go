package packet

type StatusServerRequestPacket struct {
}

func (StatusServerRequestPacket) ID() int {
	return 0x00
}

type StatusClientResponsePacket struct {
	JSON string
}

func (StatusClientResponsePacket) ID() int {
	return 0x00
}
