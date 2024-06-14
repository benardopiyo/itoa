package main

import "github.com/01-edu/z01"

// Itoa converts an integer to its string representation.
func Itoa(n int) string {
	// Handle zero explicitly
	if n == 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return "0"
	}

	// Determine if the number is negative
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	var digits []rune
	for n > 0 {
		digits = append(digits, rune('0'+n%10))
		n /= 10
	}

	// If the number was negative, add the minus sign
	if negative {
		digits = append(digits, '-')
	}

	// Output the digits in reverse order to form the correct number string
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
	z01.PrintRune('\n')

	// Convert digits slice to string (using a slice of runes here for simplicity)
	// This conversion is mainly for illustrative purposes; in actual use, we would not typically convert back from slice to string.
	result := ""
	for i := len(digits) - 1; i >= 0; i-- {
		result += string(digits[i])
	}
	return result
}

func main() {
	// Test cases for Itoa
	z01.PrintRune([]rune(Itoa(12345))[0])
	z01.PrintRune('\n')
	z01.PrintRune([]rune(Itoa(0))[0])
	z01.PrintRune('\n')
	z01.PrintRune([]rune(Itoa(-1234))[0])
	z01.PrintRune('\n')
	z01.PrintRune([]rune(Itoa(987654321))[0])
	z01.PrintRune('\n')
}
