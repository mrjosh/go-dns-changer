package main

import (
	"github.com/MrJoshLab/go-dns-changer"
	"log"
)

func main() {

	gdc.SetDebugMode(true)

	if err := gdc.SetDnsServers([]string{"Wi-Fi"}, gdc.GoogleDnsServers); err != nil {

		switch err.(type) {
		case *gdc.Error:
			log.Fatal(err.(*gdc.Error).CommandOutput())
			return
		}

		log.Fatal(err)
	}

}
