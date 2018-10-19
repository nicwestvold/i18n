package Fake

type Fake struct {
	Error  error
	Output map[string]interface{}
}

func (p *Fake) ReadFile() (map[string]interface{}, error) {
	return p.Output, p.Error
}

func (p *Fake) WriteFile(data map[string]interface{}) error {
	return p.Error
}
