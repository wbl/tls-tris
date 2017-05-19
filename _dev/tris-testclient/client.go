package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
)

var tlsVersionToName = map[string]uint16{
	"tls10": tls.VersionTLS10,
	"tls11": tls.VersionTLS11,
	"tls12": tls.VersionTLS12,
	"tls13": tls.VersionTLS13,
}

//Usage client args host:port
func main() {
	var version string
	var addr string
	var enableEMS bool
	var config tls.Config
	flag.StringVar(&version, "version", "tls12", "Version of TLS to use")
	flag.BoolVar(&enableEMS, "m", false, "Enable EMS")
	flag.Parse()
	config.MinVersion = tlsVersionToName[version]
	config.MaxVersion = tlsVersionToName[version]
	config.InsecureSkipVerify = true
	config.DisableExtendedMasterSecret = !enableEMS
	addr = flag.Arg(0)
	conn, err := tls.Dial("tcp", addr, &config)
	if err != nil {
		fmt.Println("Error %s", err)
		os.Exit(1)
	}
	conn.Close()
	os.Exit(0)
}
