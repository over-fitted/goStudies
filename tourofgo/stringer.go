package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ipA IPAddr) String() string {
	var ret string = strconv.Itoa(int(ipA[0]))
	for _, val := range ipA[1:] {
		ret += "." + strconv.Itoa(int(val))
	}
	return ret
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
