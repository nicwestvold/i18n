package real

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/nicwestvold/i18n/parser"
	"github.com/pkg/errors"
)

type Real struct {
	Filename   string
	OutputFile string
}

func New(filename, output string) (*Real, error) {
	if filename == "" {
		return nil, errors.Errorf("invalid filename '%s' provided", filename)
	}

	if output == "" {
		return nil, errors.Errorf("invalid output filename '%s' provided", output)
	}

	return &Real{
		Filename:   filename,
		OutputFile: output,
	}, nil
}

func (p *Real) ReadFile() ([]parser.Input, error) {
	var data []parser.Input

	// read the contents of the default file
	content, err := ioutil.ReadFile(p.Filename)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read JSON file")
	}

	// unmarshal the JSON
	if err := json.Unmarshal(content, &data); err != nil {
		return nil, errors.Wrap(err, "failed to decode JSON")
	}

	return data, nil
}

func (p *Real) Convert(input []parser.Input) ([]parser.Output, error) {

	parsed := make(map[string]parser.Output)
	for _, in := range input {
		if _, ok := parsed[in.Key]; !ok {
			parsed[in.Key] = parser.Output{}
		}

		for locale := range parsed[in.Key].Values {
			if locale == in.Locale {
				return nil, errors.Errorf("duplicate entry found for '%s' key under '%s' locale", in.Key, locale)
			}
		}

		values := map[string]string{}
		existing, ok := parsed[in.Key]
		if ok && existing.Values != nil {
			values = existing.Values
		}

		values[in.Locale] = in.Value
		parsed[in.Key] = parser.Output{
			Key:    in.Key,
			Values: values,
		}
	}

	var out []parser.Output
	for _, p := range parsed {
		out = append(out, p)
	}

	sort.Slice(out, func(i int, j int) bool {
		return strings.ToLower(out[i].Key) < strings.ToLower(out[j].Key)
	})

	return out, nil
}

func (p *Real) WriteFile(data []parser.Output) error {

	file, err := os.Create(p.OutputFile)
	if err != nil {
		return errors.Wrap(err, "failed to create output file")
	}

	enc := json.NewEncoder(file)
	enc.SetIndent("", "\t")
	if err := enc.Encode(data); err != nil {
		// write the new JSON to the file
		return errors.Wrap(err, "failed to write JSON to file")
	}

	return nil
}
