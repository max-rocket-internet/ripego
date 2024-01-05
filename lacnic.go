package ripego

func LacnicCheck(search, whoisServer string) (*WhoisInfo, error) {
	whoisData, err := getTcpContent(search, whoisServer)
	if err != nil {
		return nil, err
	}

	wi := WhoisInfo{
		Inetnum:      parseRPSLValue(whoisData, "inetnum", "inetnum"),
		Status:       parseRPSLValue(whoisData, "inetnum", "status"),
		Netname:      parseRPSLValue(whoisData, "inetnum", "ownerid"),
		AdminC:       parseRPSLValue(whoisData, "inetnum", "owner-c"),
		Country:      parseRPSLValue(whoisData, "inetnum", "country"),
		Descr:        parseRPSLValue(whoisData, "inetnum", "owner"),
		LastModified: parseRPSLValue(whoisData, "inetnum", "changed"),
		MntBy:        parseRPSLValue(whoisData, "inetnum", "mnt-by"),
		MntLower:     parseRPSLValue(whoisData, "inetnum", "mnt-lower"),
		MntRoutes:    parseRPSLValue(whoisData, "inetnum", "mnt-routes"),
		Source:       parseRPSLValue(whoisData, "inetnum", "source"),
		TechC:        parseRPSLValue(whoisData, "inetnum", "tech-c"),
		Organization: parseRPSLValue(whoisData, "inetnum", "owner"),

		Person: WhoisPerson{
			Name:         parseRPSLValue(whoisData, "nic-hdl", "nic-hdl"),
			AbuseMailbox: parseRPSLValue(whoisData, "nic-hdl", "e-mail"),
			Address:      parseRPSLValue(whoisData, "nic-hdl", "address"),
			LastModified: parseRPSLValue(whoisData, "nic-hdl", "changed"),
			NicHdl:       parseRPSLValue(whoisData, "role", "nic-hdl"),
			Phone:        parseRPSLValue(whoisData, "role", "phone"),
			Source:       parseRPSLValue(whoisData, "p.Co", "source"),
		},
	}

	return &wi, err
}
