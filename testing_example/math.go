package testing_example

import "fmt"

// IntMin returns the minimum of three integers
func IntMin(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// 40, 56, 50, 10, 4, 76, 3, 1, 10, 1000, 500

// SendMail uses Canada Post to send letters to random people
func SendMail() {
	fmt.Println("sent")
}
