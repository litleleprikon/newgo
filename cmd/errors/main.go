package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := fmt.Errorf("first sample error")
	err2 := fmt.Errorf("second sample error")
	err3 := fmt.Errorf("third sample error")

	complexErr1 := fmt.Errorf("first complex error %w, %w", err1, err2)
	complexErr2 := fmt.Errorf("first complex error %w, %w", complexErr1, err3)

	fmt.Println("err1 in complexErr1: ", errors.Is(complexErr1, err1))
	fmt.Println("err2 in complexErr1: ", errors.Is(complexErr1, err2))
	fmt.Println("complexErr1 in complexErr2: ", errors.Is(complexErr2, complexErr1))
	fmt.Println("err3 in complexErr2: ", errors.Is(complexErr2, err3))
	fmt.Println("err3 in complexErr1: ", errors.Is(complexErr1, err3))
}
