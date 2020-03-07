package gdc

import (
	"log"
	"os/exec"
)

var (
	macOSignoreInterfaces = []string{
		"iPhone USB", "Bluetooth PAN", "Thunderbolt Bridge", "lo0", "",
	}
)

func setDnsServersForDarwin(interfaces, servers []string) error {

	if debug {
		log.Println("Setting DNS servers for darwin/macOS")
	}

	args := append([]string{"-setdnsservers"}, interfaces...)
	args  = append(args, servers...)

	cmd := exec.Command("networksetup", args...)

	if debug {
		log.Println("Setting DNS servers for interface(s) ", interfaces)
		log.Println("Running command: ", cmd.String())
	}

	result, err := cmd.Output()
	if err != nil {
		return &Error{
			error: err,
			cmdOUT: string(result),
		}
	}

	return nil
}