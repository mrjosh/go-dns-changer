package gdc

import (
	"log"
)

func setDnsServersForWindows(interfaces, servers []string) error {

	if debug {
		log.Println("Setting DNS servers for windows")
	}

	// changing dns servers here

	return nil
}