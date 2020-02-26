package main

import (
	"fmt"
	opb "github.com/cheetahfm/operators-between-digits"
)

func main() {
	var expectedResult int64 = 200
	var digits = opb.ExpressionDigits{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	var signs = opb.ExpressionSigns{'_', '+', '-'}

	allExpressions := make(chan string)
	go opb.GenAllExpressions(allExpressions, digits, signs)

	validExpressions := make(chan string)
	go opb.CheckExpressions(allExpressions, validExpressions, expectedResult)

	fmt.Printf("Valid expressions with result %d\n", expectedResult)
	opb.PrintValidResults(validExpressions)
}
