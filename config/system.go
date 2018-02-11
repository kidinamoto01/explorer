package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)
type RestServer struct {
	Port int `yaml:"port"`

}
var Config *RestServer

func LoadConfiguration(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}
	//var config Configuration
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}
	//Config = &config
	log.Printf("config load succeessfully:%v", Config)
	return nil
}

func InitConfig(){

	//init config
	if err := LoadConfiguration("./config.yml"); err != nil {
		log.Print("config error")
		return
	}
}