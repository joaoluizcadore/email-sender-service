package services

import "log"

//ExitOnFail Exit when there is an error!
func ExitOnFail(err error, msg string) {
	if err != nil {
		if msg != "" {
			log.Println(msg)
		}
		log.Fatalf("Exiting. %v", err)
	}
}
