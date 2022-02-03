package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAccountSuccess(t *testing.T) {
	accountID := uuid.NewString()
	accountNumber := "000000001"
	accountNumberIll := "??0000001"
	tests := []struct{
		name string
		input []string
		output *Account
	}{
		{name: "create account", input: []string{accountID, accountNumber},  output: &Account{ id: accountID, number: accountNumber }},
		{name: "create account with illegible numbers", input: []string{accountID, accountNumberIll},  output: &Account{ id: accountID, number: accountNumberIll }},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAccount(tt.input[0], tt.input[1])
			assert.Nil(t, err, tt.name)
			assert.EqualValues(t, tt.output, got, tt.name)
		})
	}
}

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
			_, err := NewAccount(tt.input[0], tt.input[1])
			assert.Equal(t, tt.output, err.Error(), tt.name)
		})
	}
}