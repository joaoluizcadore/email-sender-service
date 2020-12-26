package services

import (
	"encoding/json"
	"os"
	"sync"
)

type queue struct {
	Name string
	Host string
}

type smtp struct {
	Host     string
	Port     int
	From     string
	Password string
}

type application struct {
	LogFile          string
	TemplateDir      string
	TemplateOnError  string
	SendEmailOnError bool
}

//Config - the main configuration struct
type Config struct {
	Queue       queue
	SMTP        smtp
	Application application
}

//ReadConfig Function to read from a config file.
func (c *Config) ReadConfig(filename string) {
	file, err := os.Open(filename)
	ExitOnFail(err, "Cannot open the configuration file: "+filename)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	ExitOnFail(err, "Cannot read the config file: "+filename)

}

//Prepare - It adds all the default values when not filled
func (c *Config) Prepare() {
	if c.Application.TemplateDir == "" {
		c.Application.TemplateDir = "./templates/"
	}

	if c.Application.LogFile == "" {
		c.Application.LogFile = "application.log"
	}
}

//UpdateInstance - update the global instance
func (c Config) UpdateInstance() {
	instance = &c
}

var (
	instance *Config
)
var once sync.Once

//GetConfig - get the global configuration
func GetConfig() Config {
	if instance == nil {
		instance = &Config{}
	}
	return *instance
}
