package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	Database struct {
		Endpoint string
		Username string
		Password string
		Port     int
		DB       string
	}
}

func newConfig(endpoint, username, password string, port int, db string) *Config {
	var config Config
	config.Database.Endpoint = endpoint
	config.Database.Username = username
	config.Database.Password = password
	config.Database.Port = port
	config.Database.DB = db
	return &config
}

func GenarateConfig() error {
	const tmpFile = "/tmp/sample.json"
	conf := newConfig("127.0.0.1", "root", "yourpasswordhere", 3306, "muydb")
	jsondata, err := json.MarshalIndent(conf, "", " ")
	if err != nil {
		return err
	}

	file, err := os.OpenFile(tmpFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(jsondata)
	if err != nil {
		return err
	} else {
		fmt.Printf("Sample configuration has been created in: %v\n", tmpFile)
		return nil
	}
}

func ParseConfig(filePath string) (endpoint, username, password string, port int, db string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	return config.Database.Endpoint, config.Database.Username, config.Database.Password, config.Database.Port, config.Database.DB

}

func CheckConfigFile(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			return false
		}
	} else {
		return true
	}
}
