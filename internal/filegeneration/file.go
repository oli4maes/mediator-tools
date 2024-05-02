package filegeneration

import (
	"errors"
	"fmt"
	"os"
	"text/template"
)

const (
	RequestFileType  = iota
	ResponseFileType = iota
	HandlerFileType  = iota
)

type File struct {
	FeatureName string
	FileType    int
	TmplPath    string
	FileName    string
}

func GenerateFiles(feature string, fileType int, langId int) error {
	var err error

	f := File{
		FeatureName: feature,
		FileType:    fileType,
	}

	if langId == Golang {
		f.TmplPath, f.FileName, err = getGoTemplate(f)
	}
	if langId == CSharp {
		f.TmplPath, f.FileName, err = getCSharpTemplate(f)
	}
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(f.TmplPath)
	if err != nil {
		return err
	}

	file, err := os.Create(f.FileName)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	err = tmpl.Execute(file, f)
	if err != nil {
		return err
	}

	return nil
}

func getGoTemplate(file File) (tmplPath, fileName string, err error) {
	switch file.FileType {
	case RequestFileType:
		tmplPath = "templates/go/request.tmpl"
		fileName = "request.go"
	case ResponseFileType:
		tmplPath = "templates/go/response.tmpl"
		fileName = "response.go"
	case HandlerFileType:
		tmplPath = "templates/go/handler.tmpl"
		fileName = "handler.go"
	default:
		return "", "", errors.New("unsupported file type")
	}

	return tmplPath, fileName, nil
}

func getCSharpTemplate(file File) (tmplPath, fileName string, err error) {
	switch file.FileType {
	case RequestFileType:
		tmplPath = "templates/cs/request.tmpl"
		fileName = fmt.Sprintf("%sRequest.cs", file.FeatureName)
	case ResponseFileType:
		tmplPath = "templates/cs/response.tmpl"
		fileName = fmt.Sprintf("%sResponse.cs", file.FeatureName)
	case HandlerFileType:
		tmplPath = "templates/cs/handler.tmpl"
		fileName = fmt.Sprintf("%sHandler.cs", file.FeatureName)
	default:
		return "", "", errors.New("unsupported file type")
	}

	return tmplPath, fileName, nil
}
