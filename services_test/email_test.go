package services_test

import (
	"testing"

	"github.com/joaoluizcadore/email-sender-service/services"
)

func TestEmailPrepare(t *testing.T) {
	msg := `
		{
			"To": ["email01@test.com", "email02@test.com"],
			"Subject": "Test Email",
			"Template": "",
			"Parameters": {"name": "John Doe", "message": "This is a test."}
		}
	`
	email := services.Email{}
	err := email.Parse(msg)

	if err != nil {
		t.Errorf("Got an erro: %v", err)
	}

	if len(email.To) != 2 {
		t.Errorf("Expected 2 TO address and got %v", len(email.To))
	}

	if email.Subject != "Test Email" {
		t.Errorf("Expected the subject to be %v and got %v", "Test Email", email.Subject)
	}

	if len(email.Parameters) != 2 {
		t.Errorf("Expected 2 parameters and got %v", len(email.Parameters))
	}

	if email.Parameters["name"] != "John Doe" {
		t.Errorf("Expected the parameter name be John Doe and got %v", email.Parameters["name"])
	}

}
