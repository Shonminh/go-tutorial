package main

import "fmt"
func main() {

	// 比较int
	x := LessAny(12, 122)
	fmt.Println(x)

	// 比较Int 和int
	x = LessAny(Int(12), 123312)
	fmt.Println(x)

	// 比较float64
	x = LessAny(float64(12), 3.3)
	fmt.Println(x)

	// 指定预声明类型为float64
	lessFloat := LessAny[float64]
	x = lessFloat(123, 1231.0)
	fmt.Println(x)
}

type Int int

func Less(a, b int) bool {
	return a < b
}

func LessFloat64(a, b float64) bool {
	return a < b
}

func LessInterface(a, b interface{}) bool {
	switch a.(type) {
	case int:
		return a.(int) < b.(int)
	case float64:
		return a.(float64) < b.(float64)
	}
	return false
}

func LessAny[K  ~int | ~float64](a, b K) bool {
	return a < b
}
