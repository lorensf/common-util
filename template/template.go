package template

import (
	"bytes"
	"html/template"
)

// ParseTemplate parse template with data
func ParseTemplate(name string, content string, data map[string]interface{}) (string, error) {
	var r bytes.Buffer

	t, err := template.New(name).Parse(content)
	if err != nil {
		return "", err
	}

	if err = t.Execute(&r, data); err != nil {
		return "", err
	}

	return r.String(), nil
}
