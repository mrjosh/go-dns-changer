### This project is under construction.

# Golang DNS Changer
This is a DNS changer library written in Golang.

## Notes
- IMPORTANT: Requires root/sudo/admin privileges to adjust any settings.
- Tested on:
	- macOS (Sierra, High Sierra, mojave, catalina)

## Installing
```bash
$ go get github.com/MrJoshLab/go-dns-changer
```

## Usage
Standard usage:
```go
package main

import (
    "github.com/MrJoshLab/go-dns-changer"
    "log"
)

func main() {

    interfaces := []string{"Wi-Fi"}
    dnsServers := []string{"8.8.8.8", "8.8.4.4"}    

    // you can debug with SetDebugMode to true!
    gdc.SetDebugMode(true)

    if err := gdc.SetDnsServers(interfaces, dnsServers); err != nil {
    
        switch err.(type) {
        case *gdc.Error:
            log.Fatal(err.(*gdc.Error).CommandOutput())
            return
        }
        
        log.Fatal(err)
    }

}
```

## Contributing
Thank you for considering contributing to the this project!

## License
The go-dns-changer is open-source software licensed under the MIT license.