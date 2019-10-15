package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"strconv"
)

var prefix = []byte{0xff}

type LegacyHandshakeIn struct{}

type LegacyHandshakePacket struct {
	Protocol               int
	MinecraftVersion, MOTD string
	Players, MaxPlayers    int
}

func (LegacyHandshakePacket) ID() int {
	return 0xFE
}

func (l LegacyHandshakePacket) Encode() []byte {
	encoded := []byte{0xFF}

	utf16 := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewEncoder()
	msg := strconv.Itoa(47) + "\x00"
	msg += l.MinecraftVersion + "\x00"
	msg += l.MOTD + "\x00"
	msg += strconv.Itoa(l.Players) + "\x00"
	msg += strconv.Itoa(l.MaxPlayers)

	e, _ := utf16.Bytes([]byte(msg))

	msgLen := bytes.NewBuffer([]byte{})

	binary.Write(msgLen, binary.BigEndian, int16(len(msg)+3))
	//msgLen := []byte{byte((len(msg) >> 8) & 0xFF), byte(len(msg) & 0xFF)}

	fmt.Println("msgLen size:", len(msgLen.Bytes()))

	encoded = append(encoded, msgLen.Bytes()...)
	encoded = append(encoded, 0x00, 0xa7, 0x00, 0x31, 0x00, 0x00)
	encoded = append(encoded, e...)

	return encoded
}
