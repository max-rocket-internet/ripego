package ripego

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNicProvider(t *testing.T) {

	tests := []struct {
		Whois string
		IP    string
	}{
		{"whois.apnic.net", "1.5.5.1"},
		{"whois.lacnic.net", "177.148.56.7"},
		{"whois.ripe.net", "51.2.25.4"},
		{"whois.afrinic.net", "154.125.148.148"},
		{"whois.arin.net", "184.12.95.8"},
	}

	for _, test := range tests {
		assert.Equal(t, test.Whois, getIPv4Server(net.ParseIP(test.IP).To4()))
	}
}
