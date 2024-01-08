# ripego

This is a go package for doing [WHOIS](https://en.wikipedia.org/wiki/WHOIS) lookups. These lookups can return information such as company name, contact name, physical address, country, [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) range (and more) for a given IP address.

## Install

```shell
go get github.com/max-rocket-internet/ripego
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/max-rocket-internet/ripego"
)

func main() {
	result, _ := ripego.IPLookup("1.2.3.4")

	fmt.Println("Inetnum: " + result.Inetnum)
	fmt.Println("Organization: " + result.Organization)
	fmt.Println("Netname: " + result.Netname)
	fmt.Println("Country: " + result.Country)
}
```

Output:

```console
Inetnum: 1.2.3.0 - 1.2.3.255
Organization: IRT-APNICRANDNET-AU
Netname: Debogon-prefix
Country: AU
```

## Contributing

This repo is a copy of [digineo/ripego](https://github.com/digineo/ripego/tree/master) which was a fork of [alxark/ripego](https://github.com/alxark/ripego) which was a fork of [c1982/ripego](https://github.com/c1982/ripego). The reason is that these older repos have various issues such as having issues disabled or are not simply not maintained.

Pull requests welcome 💙

To run all tests:

```console
go test ./...
```
