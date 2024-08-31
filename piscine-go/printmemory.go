package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func isLetter(b byte) bool {
	if (b >= 65 && b <= 90) || (b >= 97 && b <= 122) {
		return true
	}
	return false
}

func PrintMemory(arr [10]byte) {
	for _, v := range arr {
		if !isLetter(v) {
			v = '.'
		}
		if isLetter(v) {
			fmt.Println(tohex(int(v)))
		}

	}
}

func printstr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func tohex(n int) string {
	hex, rs := "0123456789abcdef", ""
	for n > 0 {
		rs = string(hex[n%16]) + rs
		n /= 16
	}
	return rs
}
