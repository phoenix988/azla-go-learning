package cmd

import (
	"flag"
)

// Dynamic config that you can change
type AzlaConfig struct {
    Port string // Portt number of server
    DataPath string
}

func ParseFlag()AzlaConfig {
    var azlaConfig AzlaConfig
	flag.StringVar(&azlaConfig.Port, "port", "8080", "Port Nummer to listen on")

	flag.Parse()

    return azlaConfig
}

