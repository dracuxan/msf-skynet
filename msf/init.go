package msf

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MsfHost string `yaml:"msfhost"`
	MsfPass string `yaml:"msfpass"`
	MsfUser string `yaml:"msfuser"`
}

var (
	user = os.Getenv("USER")
	path = fmt.Sprintf("/home/%s/.msf-skynet.yml", user)
)

func (c *Config) New() {
	f, err := os.Open(path)
	if err != nil {
		log.Panicf("Unable to open config file at %s\n", path)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		log.Panicf("Unable to decode yaml data. Check config file for example config")
	}
	if c.MsfHost == "" || c.MsfPass == "" || c.MsfUser == "" {
		log.Fatalln("Missing env variables MSFHOST or MSFPASS or MSFUSER")
	}
}
