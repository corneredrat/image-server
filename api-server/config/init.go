package config

import (
	"fmt"
	"errors"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Server	struct {
		ListenPort	string	`yaml:"listenPort"`
		BindAddress	string	`yaml:"bindAddress"`
	}
	Database 		struct {
		URL			string	`yaml:"url"`
		PORT		string	`yaml:"port"`
	}
	imageDir		string `yaml:"imageDir"`
} 

func (c *Config) load() error {
	
	// read config data from file.
	binData, err := ioutil.ReadFile("config/config.yaml")
	if nil != err {
		msg := fmt.Sprintf("unable to read config file 'config.yaml' : %s", err.Error())
		log.Fatal(msg)
		return errors.New("failed to load config.")
	}
	// load config into the struct variable
	err = yaml.Unmarshal(binData,&c)
	if nil != err {
		msg := fmt.Sprintf("unable to parse contents of config file 'config.yaml' : %s", err.Error())
		log.Fatal(msg)
		return errors.New("failed to load config.")
	}
	return nil
}
