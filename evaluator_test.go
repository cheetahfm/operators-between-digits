package operators_between_digits_test

import (
	opb "github.com/cheetahfm/operators-between-digits"
	"testing"
)

func TestCalculateExpression(t *testing.T) {
	testCases := []struct {
		expression string
		answer     int64
	}{
		{"1+2", 3},
		{"12-42", -30},
		{"324534", 324534},
		{"12+0", 12},
		{"12-0", 12},
		{"9876543210", 9876543210},
	}

	for _, tc := range testCases {
		result := opb.CalculateExpression(tc.expression)
		if result != tc.answer {
			t.Errorf("Expected %s=%d, but got %d", tc.expression, tc.answer, result)
		}
	}
}
