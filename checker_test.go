package operators_between_digits_test

import (
	opb "github.com/cheetahfm/operators-between-digits"
	"reflect"
	"testing"
)

func TestCheckExpressions(t *testing.T) {
	testCases := []struct {
		expressions      []string
		expectedResult   int
		validExpressions map[string]struct{}
	}{
		{
			[]string{
				"0-5",
				"12-6",
				"12-7",
				"12-7+0",
				"12-7-0",
			},
			5,
			map[string]struct{}{
				"12-7":   {},
				"12-7+0": {},
				"12-7-0": {},
			},
		},
	}
	_ = testCases

	for _, tc := range testCases {
		allExpressions := make(chan string)

		go func() {
			for _, expr := range tc.expressions {
				allExpressions <- expr
			}
			close(allExpressions)
		}()

		validExpressions := make(chan string)

		go opb.CheckExpressions(allExpressions, validExpressions, tc.expectedResult)

		for ve := range validExpressions {
			if _, present := tc.validExpressions[ve]; present {
				delete(tc.validExpressions, ve)
			} else {
				t.Errorf("Unexpected valid expression %s\n", ve)
			}
		}
		if len(tc.validExpressions) != 0 {
			keys := reflect.ValueOf(tc.validExpressions).MapKeys()

			t.Errorf("Valid expressions %v missing\n", keys)
		}
	}
}
