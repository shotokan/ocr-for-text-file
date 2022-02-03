package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	tests := []struct{
		name string
		input string
		output string
	}{
		{name: "valid checksum", input: "000000000",  output: "OK"},
		{name: "invalid checksum", input: "000000001",  output: "ERR"},
		{name: "checksum cannot be performed when digits were not recognized", input: "000?0000?",  output: "ILL"},
		{name: "checksum cannot be performed when invalid account", input: "000",  output: "ILL"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkSum(tt.input)
			assert.Equal(t, tt.output, got, tt.name)
		})
	}
}