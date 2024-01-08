package ripego

import (
	"encoding/xml"
)

type OrganizationCountry struct {
	Alpha2 string `xml:"code2"`
	Alpha3 string `xml:"code3"`
	Name   string `xml:"name"`
}

type Organization struct {
	RegistrationDate string              `xml:"registrationDate"`
	City             string              `xml:"city"`
	Country          OrganizationCountry `xml:"iso3166-1"`
	Handle           string              `xml:"handle"`
	Name             string              `xml:"name"`
	PostalCode       string              `xml:"postalCode"`
}

type NetOrg struct {
	Code string `xml:"handle,attr"`
	Name string `xml:"name,attr"`
}

type NetBlocks struct {
	NetBlock NetBlock `xml:"netBlock"`
}

type NetBlock struct {
	Description string `xml:"description"`
	CidrLength  string `xml:"cidrLength"`
	Type        string `xml:"type"`
}

type Net struct {
	Name         string    `xml:"name"`
	EndAddress   string    `xml:"endAddress"`
	StartAddress string    `xml:"startAddress"`
	OrgInfo      NetOrg    `xml:"orgRef"`
	NetBlocks    NetBlocks `xml:"netBlocks"`
	LastModified string    `xml:"updateDate"`
	Created      string    `xml:"registrationDate"`
}

func ArinCheck(search string, whoisServer string) (*WhoisInfo, error) {
	ipData, err := httpsGet(whoisServer + "/rest/ip/" + search)
	if err != nil {
		return nil, err
	}

	result, err := parseArinIpReply(ipData)
	if err != nil {
		return nil, err
	}

	whois := WhoisInfo{
		Noc:          "ARIN",
		Inetnum:      result.StartAddress + " - " + result.EndAddress,
		Netname:      result.Name,
		Organization: result.OrgInfo.Name,
		Created:      result.Created,
		LastModified: result.LastModified,
		Descr:        result.NetBlocks.NetBlock.Description,
	}

	if result.OrgInfo.Code != "" {
		orgData, err := httpsGet(whoisServer + "/rest/org/" + result.OrgInfo.Code)
		if err != nil {
			return nil, err
		}

		orgParsed, err := parseArinOrgReply(orgData)
		if err != nil {
			return nil, err
		}

		whois.Country = orgParsed.Country.Alpha2
	}

	return &whois, nil
}

func parseArinIpReply(r *[]byte) (*Net, error) {
	v := Net{}

	err := xml.Unmarshal(*r, &v)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func parseArinOrgReply(r *[]byte) (*Organization, error) {
	v := Organization{}

	err := xml.Unmarshal(*r, &v)
	if err != nil {
		return nil, err
	}

	return &v, nil
}
