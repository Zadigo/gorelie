package tests

import (
	"strings"
	"testing"

	"github.com/Zadigo/gorelie/internal"
	"github.com/stretchr/testify/assert"
)

func TestEntry(t *testing.T) {
	testCases := []struct {
		title string
		value any
	}{
		{
			title: "encode string",
			value: "Kendall Jenner",
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.title, func(t *testing.T) {
			instance := internal.Entry{
				Key:   []byte(testcase.title),
				Value: []byte((testcase.value).(string)),
			}

			if strings.Contains(testcase.title, "encode") {
				result := instance.Encode()
				assert.NotNil(t, result)
				assert.IsType(t, result, []byte{})
			}
		})
	}
}
