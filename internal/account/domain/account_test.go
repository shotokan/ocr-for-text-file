package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestAccountFail(t *testing.T) {
	accountID := uuid.NewString()
	tests := []struct{
		name string
		input []string
		output string
	}{
		{name: "empty id", input: []string{"", "000000001"},  output: "empty account id"},
		{name: "invalid id", input: []string{"notuuid", "000000001"},  output: "not valid account id"},
		{name: "empty number", input: []string{accountID, "",}, output: "empty account number"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAccount(tt.input[0], tt.input[1])
			if err.Error() != tt.output {
				t.Errorf("NewAccount() = %v, want %v", got, tt.output)
			}
		})
	}
}