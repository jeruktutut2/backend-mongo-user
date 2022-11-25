package configuration

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	WebServer WebServer `yaml:"webserver"`
	Mongo     Mongo     `yaml:"mongo"`
}

type WebServer struct {
	Port    int `yaml:"port"`
	Timeout int `yaml:"timeout"`
}

type Mongo struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Port     int    `yaml:"port"`
}

func NewConfiguration() *Configuration {
	var configuration Configuration
	file, err := ioutil.ReadFile("./configuration/configuration.yaml")
	if err != nil {
		log.Fatalf("File configuration error: %v", err)
	}
	err = yaml.Unmarshal(file, &configuration)
	if err != nil {
		log.Fatalf("Unmarshal file configuration error: %v", err)
	}
	return &configuration
}
