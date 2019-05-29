package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)


func config() (string, string, string){
	type Configuration struct {
		Hostname  string
		Port      string
		Scheme    string
		Username		  string
		Password		  string
	}
	var conf Configuration
	err := envconfig.Process("vortigaunt", &conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	host := conf.Scheme+"://"+conf.Hostname+":"+conf.Port
	username := conf.Username
	password := conf.Password
	return host, username, password
}