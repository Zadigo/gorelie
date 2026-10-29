package tests

import (
	"testing"

	"github.com/Zadigo/gorelie/internal"
	"github.com/stretchr/testify/assert"
)

func TestSchemaNewRow(t *testing.T) {
	testcases := []struct {
		title   string
		columns []internal.TableColumn
	}{
		{
			title:   "schema no columns",
			columns: []internal.TableColumn{},
		},
		{
			title: "schema with columns",
			columns: []internal.TableColumn{
				{
					Name: "firstname",
					Type: internal.TypeStr,
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			instance := internal.Schema{}
			result := instance.NewRow()

			assert.NotNil(t, result)
			assert.IsType(t, internal.Row{}, result)
		})
	}
}
