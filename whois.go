package ripego

import (
	"errors"
	"net"
)

//go:generate go run gen.go

type lookupFunc func(string, string) (*WhoisInfo, error)

var lookupFunctions = map[string]lookupFunc{
	"whois.afrinic.net4": AfrinicCheck,
	"whois.apnic.net4":   ApnicCheck,
	"whois.arin.net4":    ArinCheck,
	"whois.lacnic.net4":  LacnicCheck,
	"whois.ripe.net4":    RipeCheck4,

	"whois.afrinic.net6": AfrinicCheck,
	"whois.apnic.net6":   ApnicCheck6,
	"whois.arin.net6":    ArinCheck,
	"whois.lacnic.net6":  LacnicCheck,
	"whois.ripe.net6":    RipeCheck6,
}

func init() {
	initIPv6()
}

// IPLookup function that returns IP information at provider and returns information.
func IPLookup(ipaddr string) (*WhoisInfo, error) {
	var ipVersion string
	var whoisServer string
	ip := net.ParseIP(ipaddr)

	if ip4 := ip.To4(); ip4 != nil {
		ipVersion = "4"
		whoisServer = getIPv4Server(ip4)
	} else if ip6 := ip.To16(); ip6 != nil {
		ipVersion = "6"
		whoisServer = getIPv6Server(ip6)
	} else {
		return nil, errors.New("Invalid IP address: " + ipaddr)
	}

	if whoisServer == "" {
		return nil, errors.New("unable to find WhoIs-Server for: " + ipaddr)
	}

	lf := lookupFunctions[whoisServer+ipVersion]
	if lf == nil {
		return nil, errors.New("Unable to find whois function for: " + whoisServer)
	}

	return lf(ipaddr, whoisServer)
}

// getIPv4Server returns the whois server fot the given IPv4 address
func getIPv4Server(ip net.IP) string {
	return ipv4prefixes[ip[0]]
}

// getIPv6Server returns the whois server fot the given IPv6 address
func getIPv6Server(ip net.IP) string {
	for i := range ipv6prefixes {
		entry := &ipv6prefixes[i]
		if entry.net.Contains(ip) {
			return entry.whois
		}
	}
	return ""
}

// WhoisInfo struct with information on IP address range.
type WhoisInfo struct {
	Noc          string
	Inetnum      string
	Netname      string
	Descr        string
	Country      string
	Organization string
	AdminC       string
	TechC        string
	MntLower     string
	Status       string
	MntBy        string
	Created      string
	LastModified string
	Source       string
	MntRoutes    string
	Person       WhoisPerson
	Route        WhoisRoute
}

// WhoisPerson struct for Person information from provider.
type WhoisPerson struct {
	Name         string
	Address      string
	Phone        string
	AbuseMailbox string
	NicHdl       string
	MntBy        string
	Created      string
	LastModified string
	Source       string
}

// WhoisRoute struct for Route and Network information from provider.
type WhoisRoute struct {
	Route        string
	Descr        string
	Origin       string
	MntBy        string
	Created      string
	LastModified string
	Source       string
}
