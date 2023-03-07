package ent

import (
	"net"

	"github.com/eyedeekay/sam3"
	"github.com/eyedeekay/i2pkeys"
)

type Peer struct {
	Addresses []net.Addr
}

type Net struct {
	ListenPacketConn net.PacketConn
	SAMSession       *sam3.SAM
	Keys             i2pkeys.I2PKeys
}

func (this *Net) Close() {}

func (this *Net) Send(message string, addr net.Addr) (n int, err error) {
	return this.ListenPacketConn.WriteTo([]byte(message), addr)
}

type Message struct {
	message string
	addr    net.Addr
}

func (this *Net) Receive() (message Message, err error) {
	buf := make([]byte, 1024)
	n, addr, err := this.ListenPacketConn.ReadFrom(buf)
	if err != nil {
		return
	}
	message.message = string(buf[:n])
	message.addr = addr
	return
}

func NewUDP() (this *Net, err error) {
	this.ListenPacketConn, err = net.ListenPacket("udp", "127.0.0.1:31337")
	return
}

func NewI2P() (this *Net, err error) {
	this.SAMSession, err = sam3.NewSAM("127.0.0.1:7656")
	if err != nil {
		return
	}
	this.Keys, err = i2pkeys.LoadKeys("game-sam")
	if err != nil {
		return
	}
	this.ListenPacketConn, err = this.SAMSession.NewDatagramSession("game-sam", this.Keys, sam3.Options_Fat, 0)
	return
}
