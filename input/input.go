package input

import (
	"flag"
	"fmt"
	"strings"

	"github.com/MHSaeedkia/web-crawler/pkg/config"
)

func GetInput() (filepath string) {
	generateConfig := flag.String("generate-config", "", "Set to true to generate sample congiguration file in /tmp/sample.json")
	configFile := flag.String("config", "", "Path to the config file. use -generate-config to generate sample config")
	flag.Parse()
	if *generateConfig == "true" {
		err := config.GenarateConfig()
		if err != nil {
			fmt.Println(err.Error())
			return ""
		}
	} else if *configFile == "" {
		fmt.Println("Please specify config file. use --help for more info")
	} else {
		path := strings.TrimSuffix(*configFile, "\n")
		status := config.CheckConfigFile(path)
		if status {
			return path
		} else {
			fmt.Printf("Config file does not exist %v\n", path)
			return ""
		}
	}
	return filepath
}
