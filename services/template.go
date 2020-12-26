package services

import (
	"bytes"
	"fmt"
	"html/template"
)

//TemplateServices - It will wrap the html/template used functions
type TemplateServices interface {
	ParseFile(filename string) (*template.Template, error)
}

//TemplateServicesMock - Mock our template service
type TemplateServicesMock struct{}

//ParseFile - Mock the ParseFile funciton!
func (t TemplateServicesMock) ParseFile(filename string) (*template.Template, error) {
	const tpl = `{{.Param1}};{{.Param2}};{{.Param3}}`
	return template.New("templ").Parse(tpl)
}

//TemplateServicesImp - Our implementation
type TemplateServicesImp struct{}

//ParseFile - Get a file and return the template
func (t TemplateServicesImp) ParseFile(filename string) (*template.Template, error) {
	return template.ParseFiles(filename)
}

//ProcessTemplate - return the template already processed
func ProcessTemplate(tempServices TemplateServices, file string, params map[string]string) (string, error) {
	templ, err := tempServices.ParseFile(GetConfig().Application.TemplateDir + "/" + file)
	if err != nil {
		return "", fmt.Errorf("Cannot parse the template: %v", err)
	}
	buf := new(bytes.Buffer)
	templ.Execute(buf, params)
	return buf.String(), nil
}
