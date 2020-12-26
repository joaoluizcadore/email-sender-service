package services_test

import (
	"testing"

	"github.com/joaoluizcadore/email-sender-service/services"
)

func TestProcessTemplate(t *testing.T) {
	mock := new(services.TemplateServicesMock)
	params := make(map[string]string)
	params["Param1"] = "AB"
	params["Param2"] = "CD"
	params["Param3"] = "EF"
	templStr, err := services.ProcessTemplate(mock, "file.html", params)
	expected := "AB;CD;EF"
	if templStr != expected {
		t.Errorf("Invalid Template Render, expected %v and got %v", expected, templStr)
	}
	if err != nil {
		t.Errorf("The template got error: %v", err)
	}
}
