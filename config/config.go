package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DBName    string
	LogFile   string
}

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DBName:    cfg.Section("db").Key("dbname").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
