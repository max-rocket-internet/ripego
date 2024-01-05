package ripego

import (
	"bytes"
	"testing"
)

const (
	ARIN_XML_REPLY = `<?xml version='1.0'?><?xml-stylesheet type='text/xsl' href='https://whois.arin.net/xsl/website.xsl' ?>
	<net xmlns="https://www.arin.net/whoisrws/core/v1" xmlns:ns2="https://www.arin.net/whoisrws/rdns/v1" xmlns:ns3="https://www.arin.net/whoisrws/netref/v2" inaccuracyReportUrl="https://www.arin.net/public/whoisinaccuracy/index.xhtml" termsOfUse="https://www.arin.net/whois_tou.html">
	<registrationDate>2014-03-14T16:52:05-04:00</registrationDate>
	<ref>https://whois.arin.net/rest/net/NET-8-8-8-0-1</ref>
	<endAddress>8.8.8.255</endAddress>
	<handle>NET-8-8-8-0-1</handle>
	<name>LVLT-GOGL-8-8-8</name>
	<netBlocks>
		<netBlock>
			<cidrLength>24</cidrLength>
			<endAddress>8.8.8.255</endAddress>
			<description>Reallocated</description>
			<type>A</type>
			<startAddress>8.8.8.0</startAddress>
		</netBlock>
	</netBlocks>
	<resources inaccuracyReportUrl="https://www.arin.net/public/whoisinaccuracy/index.xhtml" termsOfUse="https://www.arin.net/whois_tou.html">
		<limitExceeded limit="256">false</limitExceeded>
	</resources>
	<orgRef handle="GOGL" name="Google Inc.">https://whois.arin.net/rest/org/GOGL</orgRef>
	<parentNetRef handle="NET-8-0-0-0-1" name="LVLT-ORG-8-8">https://whois.arin.net/rest/net/NET-8-0-0-0-1</parentNetRef>
	<startAddress>8.8.8.0</startAddress>
	<updateDate>2014-03-14T16:52:05-04:00</updateDate>
	<version>4</version>
	</net>`
)

func TestArinParser(t *testing.T) {
	w, err := parseArinReply(bytes.NewReader([]byte(ARIN_XML_REPLY)))
	if err != nil {
		t.Fatalf("Failed to parse XML reply, got: %s", err.Error())
	}

	t.Log(w.Inetnum)

}
func TestArinWhois(t *testing.T) {
	// This is a megafon NAT IP, belongs to RIPE
	w, err := ArinCheck("8.8.8.8", "whois.arin.net")
	if err != nil {
		t.Fatal("Failed to download RIPE data")
	}

	t.Logf("inetnum: %s", w.Inetnum)
	t.Logf("netname: %s", w.Netname)
	t.Logf("descr: %s", w.Descr)
	t.Logf("country: %s", w.Country)
	t.Logf("organization: %s", w.Organization)

	if w.Inetnum == "" {
		t.Fatal("Inetnum is empty, failed to get information")
	}
	if w.Netname == "" {
		t.Fatal("Netname is empty")
	}
}
