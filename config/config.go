package config

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

type ConfigList struct {
	TargetUrl string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("faild to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		TargetUrl: cfg.Section("web").Key("target_url").String(),
	}
}
