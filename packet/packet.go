package packet

import "reflect"

type Identifier interface {
	ID() int
}

type State int

// state->PacketMap
var ProtoMap map[State]PacketMap = map[State]PacketMap{}

// id->Packet
type PacketMap map[int]reflect.Type

func init() {
	ProtoMap[Handshake] = PacketMap{}
	ProtoMap[Status] = PacketMap{}
	ProtoMap[Login] = PacketMap{}
	ProtoMap[Play] = PacketMap{}

	Register(0x00, Handshake, HandshakePacket{})
}

func Register(id int, state State, packet interface{}) {
	ProtoMap[state][id] = reflect.TypeOf(packet)
}

const (
	Handshake State = iota
	Status
	Login
	Play
)
