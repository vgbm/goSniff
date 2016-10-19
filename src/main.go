package main

import(
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket"
	"fmt"
	"time"
)

func main() {
	deviceInterfaceName := "enp4s0"
	stream := captureStream(deviceInterfaceName)
	fmt.Println(*stream)
	for packet := range stream.Packets() {
		fmt.Println(packet)
	}
}

func captureStream(deviceInterfaceName string) *gopacket.PacketSource{

	handle, _ := pcap.OpenLive(
		deviceInterfaceName,
		int32(65535),		//snapshot len
		true,			//promiscuous mode
		-1 * time.Second,	//do not stop
	)

	defer handle.Close()

	packetSource := gopacket.NewPacketSource(
		handle,
		handle.LinkType(),
	)

	return packetSource
}
