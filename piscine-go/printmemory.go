package piscine

import (
	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	rs := ""
	for i, v := range arr {
		str := tohex(v)
		printstr(str)
		if v < 33 || v > 126 {
			rs += "."
		} else {
			rs += string(v)
		}
		if (i+1)%4 != 0 && i != len(arr)-1 {
			printstr(" ")
		} else {
			printstr("\n")
		}
	}
	rs += "\n"
	printstr(rs)
}

func printstr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func tohex(n byte) string {
	hex := "0123456789abcdef"
	return string(hex[n/16]) + string(hex[n%16])
}
