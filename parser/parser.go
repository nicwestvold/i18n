package parser

type Parser interface {
	ReadFile() ([]Input, error)
	Convert([]Input) ([]Output, error)
	WriteFile([]Output) error
}

type Input struct {
	Locale string `json:"locale,omitempty"`
	Key    string `json:"key,omitempty"`
	Value  string `json:"value,omitempty"`
}

type Output struct {
	Key    string            `json:"key,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
