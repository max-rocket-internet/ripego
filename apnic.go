package ripego

func ApnicCheck(search, whoisServer string) (*WhoisInfo, error) {
	whoisData, err := getTcpContent(search, whoisServer)

	if err != nil {
		return nil, err
	}

	wi := WhoisInfo{
		Inetnum:      parseRPSLValue(whoisData, "inetnum", "inetnum"),
		Netname:      parseRPSLValue(whoisData, "inetnum", "netname"),
		AdminC:       parseRPSLValue(whoisData, "inetnum", "admin-c"),
		Country:      parseRPSLValue(whoisData, "inetnum", "country"),
		Descr:        parseRPSLValue(whoisData, "inetnum", "descr"),
		LastModified: parseRPSLValue(whoisData, "inetnum", "changed"),
		MntBy:        parseRPSLValue(whoisData, "inetnum", "mnt-by"),
		MntLower:     parseRPSLValue(whoisData, "inetnum", "mnt-lower"),
		MntRoutes:    parseRPSLValue(whoisData, "inetnum", "mnt-routes"),
		Source:       parseRPSLValue(whoisData, "inetnum", "source"),
		TechC:        parseRPSLValue(whoisData, "inetnum", "tech-c"),
		Organization: parseRPSLValue(whoisData, "irt", "irt"),

		Person: WhoisPerson{
			Name:         parseRPSLValue(whoisData, "role", "role"),
			AbuseMailbox: parseRPSLValue(whoisData, "irt", "abuse-mailbox"),
			Address:      parseRPSLValue(whoisData, "role", "address"),
			LastModified: parseRPSLValue(whoisData, "role", "changed"),
			MntBy:        parseRPSLValue(whoisData, "role", "mnt-by"),
			NicHdl:       parseRPSLValue(whoisData, "role", "nic-hdl"),
			Phone:        parseRPSLValue(whoisData, "role", "phone"),
			Source:       parseRPSLValue(whoisData, "role", "source"),
		},
	}

	return &wi, err
}

func ApnicCheck6(search, whoisServer string) (*WhoisInfo, error) {
	whoisData, err := getTcpContent(search, whoisServer)

	if err != nil {
		return nil, err
	}

	wi := WhoisInfo{
		Inetnum:      parseRPSLValue(whoisData, "inet6num", "inet6num"),
		Netname:      parseRPSLValue(whoisData, "inet6num", "netname"),
		AdminC:       parseRPSLValue(whoisData, "inet6num", "admin-c"),
		Country:      parseRPSLValue(whoisData, "inet6num", "country"),
		Descr:        parseRPSLValue(whoisData, "inet6num", "descr"),
		LastModified: parseRPSLValue(whoisData, "inet6num", "changed"),
		MntBy:        parseRPSLValue(whoisData, "inet6num", "mnt-by"),
		MntLower:     parseRPSLValue(whoisData, "inet6num", "mnt-lower"),
		MntRoutes:    parseRPSLValue(whoisData, "inet6num", "mnt-routes"),
		Source:       parseRPSLValue(whoisData, "inet6num", "source"),
		TechC:        parseRPSLValue(whoisData, "inet6num", "tech-c"),
		Organization: parseRPSLValue(whoisData, "irt", "irt"),

		Person: WhoisPerson{
			Name:         parseRPSLValue(whoisData, "role", "role"),
			AbuseMailbox: parseRPSLValue(whoisData, "irt", "abuse-mailbox"),
			Address:      parseRPSLValue(whoisData, "role", "address"),
			LastModified: parseRPSLValue(whoisData, "role", "changed"),
			MntBy:        parseRPSLValue(whoisData, "role", "mnt-by"),
			NicHdl:       parseRPSLValue(whoisData, "role", "nic-hdl"),
			Phone:        parseRPSLValue(whoisData, "role", "phone"),
			Source:       parseRPSLValue(whoisData, "role", "source"),
		},
	}

	return &wi, err
}
