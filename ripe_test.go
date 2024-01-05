package ripego

import (
	"testing"
)

func TestRipeWhois(t *testing.T) {
	// This is a megafon NAT IP, belongs to RIPE
	w, err := RipeCheck4("31.173.101.1", "whois.ripe.net")
	if err != nil {
		t.Fatal("Failed to download RIPE data")
	}

	if w.Inetnum == "" {
		t.Fatal("Inetnum is empty, failed to get information")
	}
	if w.Netname == "" {
		t.Fatal("Netname is empty")
	}
	if w.Descr == "" {
		t.Fatal("Empty network description")
	}
	if w.Country == "" {
		t.Fatal("Empty country description")
	}

	t.Logf("inetnum: %s", w.Inetnum)
	t.Logf("netname: %s", w.Netname)
	t.Logf("descr: %s", w.Descr)
	t.Logf("country: %s", w.Country)
	t.Logf("organization: %s", w.Organization)
}
