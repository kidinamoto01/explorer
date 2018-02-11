package config


import (
	"testing"
	"log"
)

func TestLoadConfiguration(t *testing.T) {
	LoadConfiguration("../config.yml")
	log.Printf("config load succeessfully:%v", Config.Port)
}
