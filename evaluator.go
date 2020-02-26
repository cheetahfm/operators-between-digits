package operators_between_digits

import (
	"log"
	"regexp"
	"strconv"
)

func getAllNumbers(expr string) []int64 {
	var result []int64

	r := regexp.MustCompile(`([0-9]+)`)
	matches := r.FindAllString(expr, -1)

	for _, m := range matches {
		num, err := strconv.ParseInt(m, 10, 64)
		if err != nil {
			log.Fatalf("Parsing number '%v' error: %v\n", m, err)
		}
		result = append(result, num)
	}

	return result
}

func getAllSigns(expr string) []string {
	r := regexp.MustCompile(`([+\\-])`)
	matches := r.FindAllString(expr, -1)

	return matches
}

func checkUnsupportedSigns(expr string) {
	re := regexp.MustCompile(`[^0-9+\\-]`)
	result := re.FindAllString(expr, -1)

	if result != nil {
		log.Fatalf("Expression '%s' has unsupported signs %v", expr, result)
	}
}

// for more complex case, parser would be much better
func CalculateExpression(expr string) int64 {
	checkUnsupportedSigns(expr)

	numbers := getAllNumbers(expr)
	signs := getAllSigns(expr)

	if len(numbers) != (len(signs) + 1) {
		log.Fatal("Incorrect input")
	}
	if len(signs) == 0 {
		return numbers[0]
	}

	result := numbers[0]
	var lastSign string

	for i := 1; i < len(numbers); i++ {
		lastSign = signs[i-1]
		switch lastSign {
		case "+":
			result += numbers[i]
		case "-":
			result -= numbers[i]
		default:
			log.Fatalf("Unknown sign: %s", lastSign)
		}
	}

	return result
}
