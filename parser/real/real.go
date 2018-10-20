package real

import (
	"encoding/json"
	"github.com/nicwestvold/i18n/parser"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

type Real struct {
	Filename string
}

func New(filename string) (*Real, error) {
	if filename == "" {
		return nil, errors.Errorf("invalid filename '%s' provided", filename)
	}

	return &Real{
		Filename: filename,
	}, nil
}

func (p *Real) ReadFile() (parser.JsonFile, error) {
	var data map[string]interface{}

	// read the contents of the default file
	content, err := ioutil.ReadFile("en-US.json")
	if err != nil {
		return nil, errors.Errorf("error reading file: %v", err)
	}

	// unmarshal the JSON
	if err := json.Unmarshal(content, &data); err != nil {
		return nil, errors.Errorf("error unmarshalling JSON: %v", err)
	}

	return data, nil
}

func (p *Real) WriteFile(filename string, data parser.JsonFile) error {
	if filename == "" {
		return errors.Errorf("invalid filename '%s' provided", filename)
	}
	jsonStr, err := json.MarshalIndent(&data, "", "  ")
	if err != nil {
		return errors.Errorf("error marshalling JSON: %v", err)
	}
	// write the new JSON to the file
	err = ioutil.WriteFile(filename, jsonStr, os.ModeAppend)
	if err != nil {
		return errors.Errorf("error writing JSON file: %v", err)
	}
	return nil
}
