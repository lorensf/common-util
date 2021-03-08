package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTemplate(t *testing.T) {
	testCases := []map[string]interface{}{
		{
			"name":    "test print with data",
			"content": "test print {{ .Name }}",
			"data":    map[string]interface{}{"Name": "testingname"},
			"result":  "test print testingname",
		},
		{
			"name":    "test print with no data",
			"content": "test print data",
			"data":    map[string]interface{}{},
			"result":  "test print data",
		},
		{
			"name":    "test print error",
			"content": "test print {{ Name }}",
			"data":    map[string]interface{}{},
			"result":  "",
		},
	}

	for _, element := range testCases {
		actual, err := ParseTemplate(
			element["name"].(string),
			element["content"].(string),
			element["data"].(map[string]interface{}),
		)
		if err != nil {
			return
		}

		assert.Equal(t, element["result"].(string), actual)
	}
}
