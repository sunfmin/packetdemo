package main

import (
	"log"

	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
	//"golang.org/x/net/ipv4"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"

)

func main() {
	config := water.Config{
		DeviceType: water.TUN,
	}
	config.Name = "utun2"

	ifce, err := water.New(config)
	if err != nil {
		log.Fatal(err)
	}
	var frame ethernet.Frame

	for {
		frame.Resize(1500)
		n, err := ifce.Read(frame)
		if err != nil {
			log.Fatal(err)
		}
		frame = frame[:n]
		//log.Printf("Dst: %s\n", frame.Destination())
		//log.Printf("Src: %s\n", frame.Source())
		//log.Printf("Ethertype: % x\n", frame.Ethertype())
		//log.Printf("Payload: % x\n", frame.Payload())
		//h, err := ipv4.ParseHeader(frame)
		//if err != nil {
		//	panic(err)
		//}
		//log.Printf("Header: %+v\n", h)
		log.Printf("data %+v\n", frame)

		packet := gopacket.NewPacket(frame, layers.LayerTypeIPv4, gopacket.Default)
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer != nil {
			ip, _ := ipLayer.(*layers.IPv4)
			log.Printf("SrcAddr: %+v, DstAddr: %+v \n", ip.SrcIP, ip.DstIP)
		}

		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		//log.Println("tcpLayer", tcpLayer)
		if tcpLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)
			log.Printf("SrcPort: %+v, DstPort: %+v \n", tcp.SrcPort, tcp.DstPort)
		}
	}
}
