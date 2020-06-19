package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	wol "github.com/sabhiram/go-wol"
)

func main() {
	var mac = flag.String("mac", "", "mac address of the device that should be waked up")
	var ip = flag.String("ip", "", "broadcat IP address for network")
	flag.Parse()

	parsedIP := net.ParseIP(*ip)
	if parsedIP == nil {
		fmt.Fprintf(os.Stderr, "error parsing IP address: %s", *ip)
		os.Exit(1)
	}

	uaddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:9", *ip))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error setting up UDP connection: %s", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, uaddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error setting up UDP connection: %s", err)
		os.Exit(1)
	}
	defer conn.Close()

	mp, err := wol.New(*mac)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating magic packet: %s", err)
		os.Exit(1)
	}

	bytes, err := mp.Marshal()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error marshal magic packet: %s", err)
		os.Exit(1)
	}

	_, err = conn.Write(bytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error sending magic packet: %s", err)
		os.Exit(1)
	}
}
