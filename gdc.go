package gdc

import (
	"errors"
	"log"
	"runtime"
)

var debug = false

func SetDebugMode(mode bool) {
	debug = mode
}

func SetDnsServers(interfaces, servers []string) error {

	if len(servers) < 2 {
		return errors.New("you must include two DNS server addresses")
	}

	if debug {
		log.Println("Setting DNS servers: ", servers)
	}

	switch runtime.GOOS {
	case "darwin":
		return setDnsServersForDarwin(interfaces, servers)
	case "linux":
		return setDnsServersForLinux(interfaces, servers)
	case "windows":
		return setDnsServersForWindows(interfaces, servers)
	default:
		return errors.New("unsupported platform")
	}
}
