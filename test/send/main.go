package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	c, err := net.ListenPacket("udp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	dst, err := net.ResolveUDPAddr("udp", "255.255.255.255:8888")
	if err != nil {
		log.Fatal(err)
	}

	var packetNo, commandNo int
	var user, host, mes string
	packetNo = 1
	commandNo = 32
	user = "yamasaki"
	host = "localhost"
	mes = "Hello"
	if _, err := c.WriteTo([]byte(fmt.Sprintf("1:%d:%s:%s:%d:%s", packetNo, user, host, commandNo, mes)), dst); err != nil {
		log.Fatal(err)
	}
	b := make([]byte, 512)
	n, peer, err := c.ReadFrom(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n, "bytes read from", peer)
}
