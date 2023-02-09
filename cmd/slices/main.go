package main

import "fmt"

func main() {
	s := make([]byte, 2, 4)

	a0 := [0]byte(s)
	fmt.Println(a0)
	a1 := [1]byte(s[1:]) // a1[0] == s[1]
	fmt.Println(a1)
	a2 := [2]byte(s) // a2[0] == s[0]
	fmt.Println(a2)
	func() {
		defer fmt.Println(recover())
		a4 := [4]byte(s) // panics: len([4]byte) > len(s)
		fmt.Println(a4)
	}()

	s0 := (*[0]byte)(s) // s0 != nil
	fmt.Println(s0)
	s1 := (*[1]byte)(s[1:]) // &s1[0] == &s[1]
	fmt.Println(s1)
	s2 := (*[2]byte)(s) // &s2[0] == &s[0]
	fmt.Println(s2)
	func() {
		defer fmt.Println(recover())
		s4 := (*[4]byte)(s) // panics: len([4]byte) > len(s)
		fmt.Println(s4)
	}()

	var t []string
	t0 := [0]string(t) // ok for nil slice t
	fmt.Println(t0)
	t1 := (*[0]string)(t) // t1 == nil
	fmt.Println(t1)
	func() {
		defer fmt.Println(recover())
		t2 := (*[1]string)(t) // panics: len([1]string) > len(t)
		fmt.Println(t2)
	}()

	u := make([]byte, 0)
	u0 := (*[0]byte)(u) // u0 != nil
	fmt.Println(u0)
}
