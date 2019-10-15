package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"

	"github.com/rbrick/mcserver/conn"
	"github.com/rbrick/mcserver/packet"
)

func main() {
	l, err := net.Listen("tcp", ":25566")

	if err != nil {
		log.Fatal(err)
	}

	for {
		con, err := l.Accept()

		if err != nil {
			log.Println(err)
		} else {
			go handleConnection(conn.NewConnection(con))
		}
	}

}

func handleConnection(c *conn.Connection) {
	defer c.Close()
	for {
		p := c.ReadPacket()
		if p == nil {
			return
		}
		switch x := p.(type) {
		case packet.HandshakePacket:
			{
				log.Println("ServerAddr:", x.ServerAddr)
				log.Println("ServerPort:", x.ServerPort)
				log.Println("Protocol Version:", x.ProtoVersion.Val)

				c.SetState(packet.State(x.NextState.Val))
			}
		case packet.LegacyHandshakeIn:
			{
				fmt.Println(hex.Dump(packet.LegacyHandshakePacket{
					Protocol:         74,
					MinecraftVersion: "1.8.7",
					MOTD:             "A Minecraft Server",
					Players:          0,
					MaxPlayers:       20,
				}.Encode()))
				c.Write(packet.LegacyHandshakePacket{
					Protocol:         74,
					MinecraftVersion: "ur gay",
					MOTD:             "lol no u",
					Players:          69696969,
					MaxPlayers:       69696969,
				}.Encode())
				c.Close()
				return
			}
		}
		return
	}
}
