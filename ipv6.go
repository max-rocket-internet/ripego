package ripego

import "net"

// ipv6prefix contains a prefix and the corresponding WhoIs server
type ipv6Prefix struct {
	net   net.IPNet
	whois string
}

var ipv6prefixes = make([]ipv6Prefix, 0, 64)

func addIPv6Prefix(prefix, whois string) {
	_, ipnet, err := net.ParseCIDR(prefix)
	if err != nil {
		panic(err)
	}

	ipv6prefixes = append(ipv6prefixes, ipv6Prefix{
		net:   *ipnet,
		whois: whois,
	})
}
