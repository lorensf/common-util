package timefmt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeInLocalTimezone(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	testCases := []map[string]time.Time{
		{
			"case":   time.Date(1970, 01, 01, 2, 0, 0, 0, time.UTC),
			"result": time.Date(1970, 01, 01, 9, 0, 0, 0, loc),
		},
		{
			"case":   time.Date(1970, 01, 01, 11, 0, 0, 0, time.UTC),
			"result": time.Date(1970, 01, 01, 18, 0, 0, 0, loc),
		},
		{
			"case":   time.Date(1970, 01, 01, 20, 0, 0, 0, time.UTC),
			"result": time.Date(1970, 01, 02, 3, 0, 0, 0, loc),
		},
	}

	for _, element := range testCases {
		assert.Equal(t, element["result"], TimeInLocalTimezone(element["case"]))
	}
}
