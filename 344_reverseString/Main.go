package main

import "fmt"

func main() {
	str1 := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(str1)
	reverseString(str1)
	fmt.Println(str1)
}

func reverseString(s []byte) {
	len := len(s)
	for i := 0; i < len/2; i++ {
		s[i], s[len-1-i] = s[len-1-i], s[i]
	}
}
