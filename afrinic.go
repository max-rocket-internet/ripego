package ripego

type afrinic struct {
}

// AfrinicCheck function
func AfrinicCheck(search, whoisServer string) (*WhoisInfo, error) {
	whoisData, err := getTcpContent(search, whoisServer)

	if err != nil {
		return nil, err
	}

	wi := WhoisInfo{
		Inetnum:      parseRPSLValue(whoisData, "inetnum", "inetnum"),
		Netname:      parseRPSLValue(whoisData, "inetnum", "netname"),
		AdminC:       parseRPSLValue(whoisData, "inetnum", "admin-c"),
		Country:      parseRPSLValue(whoisData, "inetnum", "country"),
		Created:      parseRPSLValue(whoisData, "inetnum", "created"),
		Descr:        parseRPSLValue(whoisData, "inetnum", "descr"),
		LastModified: parseRPSLValue(whoisData, "inetnum", "last-modified"),
		MntBy:        parseRPSLValue(whoisData, "inetnum", "mnt-by"),
		MntLower:     parseRPSLValue(whoisData, "inetnum", "mnt-lower"),
		MntRoutes:    parseRPSLValue(whoisData, "inetnum", "mnt-routes"),
		Source:       parseRPSLValue(whoisData, "inetnum", "source"),
		TechC:        parseRPSLValue(whoisData, "inetnum", "tech-c"),
		Organization: parseRPSLValue(whoisData, "inetnum", "org"),
		Person: WhoisPerson{
			Name:         parseRPSLValue(whoisData, "person", "person"),
			AbuseMailbox: parseRPSLValue(whoisData, "person", "abuse-mailbox"),
			Address:      parseRPSLValue(whoisData, "person", "address"),
			Created:      parseRPSLValue(whoisData, "person", "created"),
			LastModified: parseRPSLValue(whoisData, "person", "last-modified"),
			MntBy:        parseRPSLValue(whoisData, "person", "mnt-by"),
			NicHdl:       parseRPSLValue(whoisData, "person", "nic-hdl"),
			Phone:        parseRPSLValue(whoisData, "person", "phone"),
			Source:       parseRPSLValue(whoisData, "person", "source"),
		},
		Route: WhoisRoute{
			Origin:       parseRPSLValue(whoisData, "route", "origin"),
			Created:      parseRPSLValue(whoisData, "route", "created"),
			Descr:        parseRPSLValue(whoisData, "route", "descr"),
			LastModified: parseRPSLValue(whoisData, "route", "last-modified"),
			Route:        parseRPSLValue(whoisData, "route", "route"),
			Source:       parseRPSLValue(whoisData, "route", "source"),
		},
	}

	return &wi, err
}
