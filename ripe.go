package ripego

func RipeCheck4(search, whoisServer string) (*WhoisInfo, error) {
	whoisData, err := getTcpContent(search, whoisServer)

	if err != nil {
		return nil, err
	}

	wi := WhoisInfo{
		Noc:          "RIPE",
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
		Status:       parseRPSLValue(whoisData, "inetnum", "status"),

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

func RipeCheck6(search, whoisServer string) (*WhoisInfo, error) {
	whoisData, err := getTcpContent(search, whoisServer)

	if err != nil {
		return nil, err
	}

	wi := WhoisInfo{
		Inetnum:      parseRPSLValue(whoisData, "inet6num", "inet6num"),
		Netname:      parseRPSLValue(whoisData, "inet6num", "netname"),
		AdminC:       parseRPSLValue(whoisData, "inet6num", "admin-c"),
		Country:      parseRPSLValue(whoisData, "inet6num", "country"),
		Created:      parseRPSLValue(whoisData, "inet6num", "created"),
		Descr:        parseRPSLValue(whoisData, "inet6num", "descr"),
		LastModified: parseRPSLValue(whoisData, "inet6num", "last-modified"),
		MntBy:        parseRPSLValue(whoisData, "inet6num", "mnt-by"),
		MntLower:     parseRPSLValue(whoisData, "inet6num", "mnt-lower"),
		MntRoutes:    parseRPSLValue(whoisData, "inetnum", "mnt-routes"),
		Source:       parseRPSLValue(whoisData, "inet6num", "source"),
		TechC:        parseRPSLValue(whoisData, "inet6num", "tech-c"),
		Organization: parseRPSLValue(whoisData, "inet6num", "org"),
		Status:       parseRPSLValue(whoisData, "inet6num", "status"),

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
			Origin:       parseRPSLValue(whoisData, "route6", "origin"),
			Created:      parseRPSLValue(whoisData, "route6", "created"),
			Descr:        parseRPSLValue(whoisData, "route6", "descr"),
			LastModified: parseRPSLValue(whoisData, "route6", "last-modified"),
			Route:        parseRPSLValue(whoisData, "route6", "route6"),
			Source:       parseRPSLValue(whoisData, "route6", "source"),
		},
	}

	return &wi, err
}
