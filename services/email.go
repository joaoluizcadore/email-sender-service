package services

import (
	"encoding/json"
	"fmt"
	"log"
)

//MessageHandlerType - Function type
type MessageHandlerType func([]byte)

//Email - Struct of an Email Message
type Email struct {
	To         []string
	Parameters map[string]string
	Template   string
	Subject    string
	Body       string
	IsHTML     bool
}

//Parse - Parse a json string to the email object and validate it
func (e *Email) Parse(body string) error {
	err := json.Unmarshal([]byte(body), &e)
	if err != nil {
		log.Printf("Cannot read the message [%v] -> %v", body, err)
		return fmt.Errorf("Invalid email message: %v", err)
	}

	if len(e.To) == 0 {
		return fmt.Errorf("There is no destination address")
	}
	if e.Template != "" {
		tempServices := TemplateServicesImp{}
		e.Body, err = ProcessTemplate(tempServices, e.Template, e.Parameters)
		if err != nil {
			return err
		}
	}
	return nil
}

func newEmail(config Config, body []byte) (*Email, error) {
	email := &Email{}
	err := email.Parse(string(body))
	if err != nil {
		return nil, err
	}
	return email, nil
}

//StartService - It will literally enable the service!
func StartService() {
	StartQueue(func(body []byte) {
		go MessageHandler(body)
	})
}

//MessageHandler - Process the e-mail from the queue
func MessageHandler(body []byte) {
	config := GetConfig()
	log.Printf("Received message -> %d bytes\n", len(body))
	email, err := newEmail(config, body)
	if err != nil {
		log.Printf("Cannot process the message: %s\n", err)
	} else {
		log.Printf("Sending e-mail [%v] to [%v]", email.Subject, email.To)
		err := SendEmail(config, *email)
		if err != nil {
			log.Printf("ERROR - Cannot send the e-mail: %v\n", err)
		} else {
			log.Printf("E-mail has been sent!")
		}
	}
}
