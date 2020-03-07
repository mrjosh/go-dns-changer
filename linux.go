package gdc

import (
	"log"
)

func setDnsServersForLinux(interfaces, servers []string) error {

	if debug {
		log.Println("Setting DNS servers for linux")
	}

	// changing dns servers here

	return nil
}