package main

import "fmt"

func main() {
	res, err := divide(100, 0)
	if err == nil {
		fmt.Println("result:", res)
	}
	fmt.Println("error:", err)
}

type MyError struct {
	Message   string
	ErrorCode int
}

func (err *MyError) Error() string {
	return fmt.Sprintf("Error:%s ErrorCode:%d\n", err.Message, err.ErrorCode)
}

func divide(a, b int) (float32, error) {
	if b == 0 {
		return 0, &MyError{
			"Failed to divide by zero",
			429,
		}
	}
	return float32(a / b), nil
}
