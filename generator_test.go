package operators_between_digits_test

import (
	"bytes"
	opb "github.com/cheetahfm/operators-between-digits"
	"reflect"
	"sync"
	"testing"
)

func TestGenAllExpressions(t *testing.T) {
	testCases := []struct {
		digits      opb.ExpressionDigits
		signs       opb.ExpressionSigns
		expressions map[string]struct{}
	}{
		{
			opb.ExpressionDigits{1, 2, 3},
			[]opb.Sign{'_', '+', '-'},
			map[string]struct{}{
				"123":   {},
				"12+3":  {},
				"1+2+3": {},
				"1-2+3": {},
				"12-3":  {},
				"1+2-3": {},
				"1-2-3": {},
				"1+23":  {},
				"1-23":  {},
			},
		},
	}

	for _, tc := range testCases {
		var wg sync.WaitGroup
		wg.Add(1)
		allExpressions := make(chan string)

		go opb.GenAllExpressions(allExpressions, tc.digits, tc.signs)

		printSigns := func(signs opb.ExpressionSigns) string {
			var buffer bytes.Buffer
			buffer.WriteString("'")
			for _, s := range signs {
				buffer.WriteString(string(s))
			}
			buffer.WriteString("'")
			return buffer.String()
		}

		go func() {
			defer wg.Done()

			for {
				expr, ok := <-allExpressions
				if !ok {
					break
				}
				_, present := tc.expressions[expr]
				if present {
					delete(tc.expressions, expr)
				} else {
					t.Errorf("Unexpected expression '%s' generated for signs %s and digits %v\n",
						expr, printSigns(tc.signs), tc.digits)
				}
			}
			if len(tc.expressions) != 0 {
				keys := reflect.ValueOf(tc.expressions).MapKeys()
				t.Errorf("Expressions %v have not been generated for signs %v and digits %v\n",
					keys, printSigns(tc.signs), tc.digits)
			}
		}()
		wg.Wait()
	}
}
