package main

import "fmt"

func main() {
	testString := "가갛힣"
	for i, r := range testString {
		fmt.Println(i, r, string(r))
	}
	fmt.Println(len(testString))
}
