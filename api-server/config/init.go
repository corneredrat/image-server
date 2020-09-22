package config

import (
	"os"
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
	Kafka			struct {
		URL			string `yaml: "url"`
		PORT		string	`yaml:"port"`
	}
	ImagePath		string `yaml:"imageDir"`
} 

var Options Config

func Load() error {
	
	// read config data from file.
	binData, err := ioutil.ReadFile("config/config.yaml")
	if nil != err {
		msg := fmt.Sprintf("unable to read config file 'config.yaml' : %s", err.Error())
		log.Fatal(msg)
		return errors.New("failed to load config.")
	}
	// load config into the struct variable
	err = yaml.Unmarshal(binData,&Options)
	if nil != err {
		msg := fmt.Sprintf("unable to parse contents of config file 'config.yaml' : %s", err.Error())
		log.Fatal(msg)
		return errors.New("failed to load config.")
	}
	// override defaults with environment variables
	// Mongo URL
	
	if mongoURL, ok := os.LookupEnv("MONGO_URL"); ok {
		Options.Database.URL = mongoURL
	}
	if mongoPORT, ok := os.LookupEnv("MONGO_PORT"); ok {
		Options.Database.PORT = mongoPORT
	}
	if kafkaURL, ok := os.LookupEnv("KAFKA_URL"); ok {
		Options.Kafka.URL = kafkaURL
	}
	if kafkaPORT, ok := os.LookupEnv("KAFKA_PORT"); ok {
		Options.Kafka.PORT = kafkaPORT
	}
	msg := fmt.Sprintf("Listen Port : %v",Options.Server.ListenPort)
	log.Info(msg)
	msg = fmt.Sprintf("Listen Addr : %v",Options.Server.BindAddress)
	log.Info(msg)
	msg = fmt.Sprintf("MongoDB URL : %v",Options.Database.URL)
	log.Info(msg)
	msg = fmt.Sprintf("MongoDB PORT: %v",Options.Database.PORT)
	log.Info(msg)
	msg = fmt.Sprintf("Kafka URL   : %v",Options.Kafka.URL)
	log.Info(msg)
	msg = fmt.Sprintf("Kafka PORT  : %v",Options.Kafka.PORT)
	log.Info(msg)
	return nil
}
