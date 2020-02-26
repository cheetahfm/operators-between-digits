package operators_between_digits

import (
	"bytes"
	"log"
	"strconv"
)

type Sign rune

func (s Sign) String(printNothingSign bool) string {
	if s == '_' {
		if !printNothingSign {
			return ""
		}
	}
	return string(s)
}

type ExpressionDigits []int64
type ExpressionSigns []Sign

type Expression struct {
	Signs  ExpressionSigns
	Digits ExpressionDigits
}

func (e Expression) String(printNumbers bool, printNothingSign bool) string {
	if len(e.Digits) != (len(e.Signs) + 1) {
		log.Printf("expression: %v\n", e.Signs)
		log.Printf("digits: %v\n", e.Digits)
		log.Fatal("Expression and digits size do not match")
	}

	var buffer bytes.Buffer

	for i := 0; i < len(e.Signs); i++ {
		if printNumbers {
			buffer.WriteString(strconv.FormatInt(e.Digits[i], 10))
		}
		buffer.WriteString(e.Signs[i].String(printNothingSign))
	}
	if printNumbers {
		buffer.WriteString(strconv.FormatInt(e.Digits[len(e.Signs)], 10))
	}
	return buffer.String()
}

func GenAllExpressions(allExpressions chan<- string, digits ExpressionDigits,
	signOptions ExpressionSigns) {
	expr := Expression{
		Signs:  make(ExpressionSigns, len(digits)-1),
		Digits: digits,
	}
	genExpression(allExpressions, expr, 0, signOptions)
	close(allExpressions)
}

func genExpression(results chan<- string, expr Expression, idx uint8, signOptions ExpressionSigns) {
	for _, s := range signOptions {
		expr.Signs[idx] = s
		if int(idx) == (len(expr.Signs) - 1) {
			preparedExpression := expr.String(true, false)
			results <- preparedExpression
		} else {
			genExpression(results, expr, idx+1, signOptions)
		}
	}
}
