package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/joaoluizcadore/email-sender-service/services"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "configFile", "config.json", "Inform the configuration file.")
	flag.Parse()
}

func main() {
	config := &services.Config{}
	config.ReadConfig(configFile)
	config.Prepare()
	config.UpdateInstance()

	file, err := os.OpenFile(config.Application.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	services.ExitOnFail(err, "Cannot start the log file: "+config.Application.LogFile)
	defer file.Close()

	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)

	log.Println("Starting the application... ")
	log.Println("Config File: " + configFile)
	log.Printf("TemplateDir: %v\n", config.Application.TemplateDir)
	log.Printf("Queue name: %v\n", config.Queue.Name)

	services.StartService()

}
