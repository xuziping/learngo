package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupport")
	}
}

func apply(op func(int, int) int, a int, b int) int {
	p := reflect.ValueOf(op).Pointer()
	name := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s", name)
	return op(a, b)
}

func div(a, b int) (q, r int) {
	q = a / b
	return q, a % b
}

func pow(a, b  int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	if result, err := eval(3,4,"+"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 4)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1,4))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap (a, b *int) {
	*b, *a = *a, *b
}