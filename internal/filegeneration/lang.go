package filegeneration

import "errors"

const (
	Golang = iota
	CSharp = iota
)

func GetLangType(lang string) (int, error) {
	switch lang {
	case "go":
		return Golang, nil
	case "cs":
		return CSharp, nil
	default:
		return 0, errors.New("unsupported language")
	}
}
