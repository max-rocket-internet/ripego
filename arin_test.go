package ripego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ARIN_IP_XML_REPLY  = `<?xml version='1.0'?><?xml-stylesheet type='text/xsl' href='https://whois.arin.net/xsl/website.xsl' ?><net xmlns="http://www.arin.net/whoisrws/core/v1" xmlns:ns2="http://www.arin.net/whoisrws/rdns/v1" xmlns:ns3="http://www.arin.net/whoisrws/netref/v2" copyrightNotice="Copyright 1997-2024, American Registry for Internet Numbers, Ltd." inaccuracyReportUrl="https://www.arin.net/resources/registry/whois/inaccuracy_reporting/" termsOfUse="https://www.arin.net/resources/registry/whois/tou/"><registrationDate>2023-12-28T17:24:33-05:00</registrationDate><rdapRef>https://rdap.arin.net/registry/ip/8.8.8.0</rdapRef><ref>https://whois.arin.net/rest/net/NET-8-8-8-0-2</ref><endAddress>8.8.8.255</endAddress><handle>NET-8-8-8-0-2</handle><name>GOGL</name><netBlocks><netBlock><cidrLength>24</cidrLength><endAddress>8.8.8.255</endAddress><description>Direct Allocation</description><type>DA</type><startAddress>8.8.8.0</startAddress></netBlock></netBlocks><resources copyrightNotice="Copyright 1997-2024, American Registry for Internet Numbers, Ltd." inaccuracyReportUrl="https://www.arin.net/resources/registry/whois/inaccuracy_reporting/" termsOfUse="https://www.arin.net/resources/registry/whois/tou/"><limitExceeded limit="256">false</limitExceeded></resources><orgRef handle="GOGL" name="Google LLC">https://whois.arin.net/rest/org/GOGL</orgRef><parentNetRef handle="NET-8-0-0-0-0" name="NET8">https://whois.arin.net/rest/net/NET-8-0-0-0-0</parentNetRef><startAddress>8.8.8.0</startAddress><updateDate>2023-12-28T17:24:56-05:00</updateDate><version>4</version></net>`
	ARIN_ORG_XML_REPLY = `<?xml version='1.0'?><?xml-stylesheet type='text/xsl' href='https://whois.arin.net/xsl/website.xsl' ?><org xmlns="http://www.arin.net/whoisrws/core/v1" xmlns:ns2="http://www.arin.net/whoisrws/rdns/v1" xmlns:ns3="http://www.arin.net/whoisrws/netref/v2" copyrightNotice="Copyright 1997-2024, American Registry for Internet Numbers, Ltd." inaccuracyReportUrl="https://www.arin.net/resources/registry/whois/inaccuracy_reporting/" termsOfUse="https://www.arin.net/resources/registry/whois/tou/"><registrationDate>2000-03-30T00:00:00-05:00</registrationDate><rdapRef>https://rdap.arin.net/registry/entity/GOGL</rdapRef><ref>https://whois.arin.net/rest/org/GOGL</ref><canAllocate>Y</canAllocate><city>Mountain View</city><iso3166-1><code2>US</code2><code3>USA</code3><name>United States</name><e164>1</e164></iso3166-1><handle>GOGL</handle><name>Google LLC</name><postalCode>94043</postalCode><comment><line number="0">Please note that the recommended way to file abuse complaints are located in the following links. </line><line number="1"></line><line number="2">To report abuse and illegal activity: https://www.google.com/contact/</line><line number="3"></line><line number="4">For legal requests: http://support.google.com/legal </line><line number="5"></line><line number="6">Regards, </line><line number="7">The Google Team</line></comment><iso3166-2>CA</iso3166-2><streetAddress><line number="0">1600 Amphitheatre Parkway</line></streetAddress><updateDate>2019-10-31T15:45:45-04:00</updateDate><resources copyrightNotice="Copyright 1997-2024, American Registry for Internet Numbers, Ltd." inaccuracyReportUrl="https://www.arin.net/resources/registry/whois/inaccuracy_reporting/" termsOfUse="https://www.arin.net/resources/registry/whois/tou/"><limitExceeded limit="256">false</limitExceeded></resources></org>`
)

func TestArinParser(t *testing.T) {
	data := []byte(ARIN_IP_XML_REPLY)
	_, err := parseArinIpReply(&data)

	assert.NoError(t, err)
}

func TestParseArinOrgReply(t *testing.T) {
	data := []byte(ARIN_ORG_XML_REPLY)
	_, err := parseArinOrgReply(&data)

	assert.NoError(t, err)
}

func TestArinWhois(t *testing.T) {
	result, err := ArinCheck("8.8.8.8", "whois.arin.net")
	assert.NoError(t, err)

	assert.NotEmpty(t, result.Inetnum)
	assert.NotEmpty(t, result.Netname)
	assert.NotEmpty(t, result.Descr)
	assert.NotEmpty(t, result.Country)
	assert.NotEmpty(t, result.Organization)
}
