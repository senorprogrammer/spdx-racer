package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_isLicensableFile(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "with blank extension",
			path:     "cats",
			expected: false,
		},
		{
			name:     "with licensable extension",
			path:     ".go",
			expected: true,
		},
		{
			name:     "with unlicensable extension",
			path:     ".rs",
			expected: false,
		},
		{
			name:     "with blank name",
			path:     ".go",
			expected: true,
		},
		{
			name:     "with licensable name",
			path:     "main.go",
			expected: true,
		},
		{
			name:     "with unlicensable name",
			path:     "main_mocks.go",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := NewSPDXWriter()
			writer.fileTypes = []string{"go"}
			writer.license = "MPL-2.0"

			actual := writer.isLicensableFile(tt.path)

			assert.Equal(t, actual, tt.expected)
		})
	}
}
