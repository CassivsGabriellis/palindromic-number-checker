package main

import "fmt"

func isPalindrome(num int) bool {
	if num < 0 {
		num = -num // make the number positive
	}
	if num < 10 {
		return true // a single digit is always a palindrome
	}
	if num%10 == 0 {
		return false // if the last digit is 0, it cannot be a palindrome
	}

	reverse := 0
	temp := num

	for temp > reverse {
		reverse = reverse*10 + temp%10
		temp /= 10
	}

	// When the number has an odd number of digits, we can remove the middle digit
	// from 'reverse' and directly compare it with 'temp' to check if it's a palindrome.
	return temp == reverse || temp == reverse/10
}

func main() {
	num := 45
	if isPalindrome(num) {
		fmt.Println(num, "is a palindrome number.")
	} else {
		fmt.Println(num, "is not a palindrome number.")
	}
}
