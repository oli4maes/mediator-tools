package filegeneration

import (
	"errors"
	"os"
	"text/template"
)

const (
	RequestFileType  = iota
	ResponseFileType = iota
	HandlerFileType  = iota
)

type Feature struct {
	Name string
}

func GenerateFiles(feature string, fileType int) error {
	f := Feature{
		Name: feature,
	}

	var tmplPath string
	var fileName string

	switch fileType {
	case RequestFileType:
		tmplPath = "templates/request.tmpl"
		fileName = "request.go"
	case ResponseFileType:
		tmplPath = "templates/response.tmpl"
		fileName = "response.go"
	case HandlerFileType:
		tmplPath = "templates/handler.tmpl"
		fileName = "handler.go"
	default:
		return errors.New("unsupported file type")
	}

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()
	err = tmpl.Execute(file, f)
	if err != nil {
		return err
	}

	return nil
}
