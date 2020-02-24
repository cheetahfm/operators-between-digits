package operators_between_digits

import (
	"fmt"
	"runtime"
	"sync"
)

func maxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

// reads from expressions channel,
// sends expressions to evaluator and
// returns only those with the expected result
// via validExpressions channel
func CheckExpressions(expressions <-chan string, validExpressions chan<- string,
	expectedResult int) {
	numThreads := maxParallelism()
	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				expr, ok := <-expressions
				if !ok {
					return
				}
				exprResult := CalculateExpression(expr)
				if exprResult == expectedResult {
					validExpressions <- expr
				}
			}
		}()
	}
	wg.Wait()
	close(validExpressions)
}

func PrintValidResults(validExpressions chan string) {
	for {
		e, ok := <-validExpressions
		if !ok {
			break
		}
		fmt.Println(e)
	}
}
