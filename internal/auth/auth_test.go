package auth

import (
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		Name    string
		Input   []string
		Output  string
		IsError bool
	}{
		{
			Name:    "Valid authorization",
			Input:   []string{"ApiKey 123"},
			Output:  "123",
			IsError: false,
		},
		{
			Name:    "Invalid Length",
			Input:   []string{"ApiKey 123 456"},
			Output:  "123",
			IsError: false,
		},
		{
			Name:    "Invalid Length",
			Input:   []string{"ApiKey"},
			Output:  "",
			IsError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			header := make(map[string][]string)
			header["Authorization"] = test.Input

			output, err := GetAPIKey(header)
			if err != nil && !test.IsError {
				t.Errorf("Error while running test %s", err)
				return
			}

			if output != test.Output {
				t.Errorf("Error while running test input:%s output:%s", test.Input[0], test.Output)
				return
			}
		})
	}
}
