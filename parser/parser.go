package parser

type JsonFile map[string]interface{}

type Parser interface {
	ReadFile() (JsonFile, error)
	WriteFile() error
}
