package cmd

import (
	"flag"
	"fmt"
	"os"
)

// Dynamic config that you can change
type AzlaConfig struct {
	Port     string // Portt number of server
	DataPath string
	Help     bool
}

var azlaConfig AzlaConfig

// Parse command line arguments
func ParseFlag() AzlaConfig {
	flag.StringVar(&azlaConfig.Port, "p", "8080", "Port Nummer to listen on")
	flag.StringVar(&azlaConfig.Port, "port", "8080", "Port Nummer to listen on")
	flag.StringVar(&azlaConfig.DataPath, "data", "data/", "Data Path")
	flag.BoolVar(&azlaConfig.Help, "help", false, "Print Help Information")

	flag.Parse()

	if azlaConfig.Help {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	return azlaConfig
}


// Retrieve the data
func GetFlag() AzlaConfig {
    return azlaConfig
}
