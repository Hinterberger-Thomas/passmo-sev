package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type MySQL struct {
	DbName   string `yaml:"dbName"`
	DbPass   string `yaml:"dbPass"`
	UserName string `yaml:"userName"`
	Port     uint32 `yaml:"port"`
	IP       string `yaml:"ip"`
	Protocol string `yaml:"protocol"`
}

func GetConfMySQL() MySQL {
	var c MySQL
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
