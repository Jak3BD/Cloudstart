package utils_test

import (
	"cloudstart/utils"
	"fmt"
	"testing"
)

func TestPrintDefault(t *testing.T) {
	tests := []struct {
		message    string
		defaultVal any
		want       string
	}{
		{"Hello", nil, "Hello [<nil>]"},
		{"Status", 200, "Status [200]"},
		{"User", "Alice", "User [Alice]"},
		{"Balance", 99.99, "Balance [99.99]"},
		{"Active", true, "Active [true]"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%T", test.defaultVal), func(t *testing.T) {
			out := utils.PrintDefault(test.message, test.defaultVal)
			if out != test.want {
				t.Errorf("'%s' != '%s'", out, test.want)
			}
		})
	}
}
