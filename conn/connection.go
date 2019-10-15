package conn

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
	"sync"

	"github.com/rbrick/mcserver/packet"
)

type Connection struct {
	m     *sync.Mutex
	conn  net.Conn
	state packet.State
}

func NewConnection(c net.Conn) *Connection {
	return &Connection{&sync.Mutex{}, c, packet.Handshake}
}

func (c *Connection) Close() {
	defer c.m.Unlock()
	c.m.Lock()
	c.conn.Close()
}

func (c *Connection) SetState(state packet.State) {
	defer c.m.Unlock()
	c.m.Lock()
	c.state = state
}

func (c *Connection) Write(b []byte) {
	c.conn.Write(b)
}

func (c *Connection) ReadPacket() interface{} {
	rd := bufio.NewReader(c.conn)

	if b, _ := rd.ReadByte(); b == 0xFE {
		// legacy ping
		fmt.Println("legacy ping")
		return packet.LegacyHandshakeIn{}
	}

	_ = rd.UnreadByte()

	size, err := binary.ReadUvarint(rd)
	if err != nil {
		return nil
	}

	fmt.Println("size:", size)

	buf := make([]byte, size)
	_, _ = io.ReadAtLeast(rd, buf, int(size))

	pr := bufio.NewReader(bytes.NewReader(buf))

	id, _ := binary.ReadUvarint(pr)

	fmt.Println("id:", id)

	if v, ok := packet.ProtoMap[c.state][int(id)]; !ok {
		log.Panic("No packet for ID:", id)
	} else {
		p := reflect.New(v).Elem()

		for i := 0; i < p.NumField(); i++ {
			f := p.Field(i)
			if codec, ok := CodecMap[f.Type()]; ok {
				val := codec.Decode(pr)
				fmt.Println("val", val)
				f.Set(reflect.ValueOf(val))
				fmt.Println("type: ", f.Type())
			}
		}

		return p.Interface()
	}

	return nil
}
