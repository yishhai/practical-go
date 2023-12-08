package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G😃", 6)

	// code points = runes = unicode character
	s := "Go"

	for i, r := range s {
		fmt.Println("index of string char:", i, "unicode code point:", r)
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r) // rune (int32)
		}
	}

	fmt.Printf("%c of type %T\n", s[0], s[0]) // byte (uint8)

	/*
		when strings of chars are looped over, each individual type is treated as a rune (int32) by Go,
		but when indexed, Go treats them as byte (uint8)

	*/

	d := len("G😃")                    // bytes (uint8)
	v := utf8.RuneCountInString("G😃") // rune (int32)
	fmt.Printf("d: %v\n", d)          // 5
	fmt.Printf("v: %v\n", v)          // 2

	res := isPalindrome("madam")
	fmt.Printf("is palindrome?: %v\n", res)

	res2 := isPalindrome2("madam")
	fmt.Printf("is palindrome?: %v\n", res2)
}

// Check if a word is the same when printed backwards (if it's a palindrome)
func isPalindrome(word string) bool {
	start := 0
	end := utf8.RuneCountInString(word) - 1

	for i := start; i < end; i++ {

		if word[end] != word[start] {
			return false
		}

		start++
		end--
	}

	return true
}

// Check if a word is the same when printed backwards (if it's a palindrome)
func isPalindrome2(word string) bool {
	rs := []rune(word)
	for i := 0; i < len(word)/2; i++ {
		if rs[i] != rs[len(rs)-i-1] {
			return false
		}

	}
	return true
}

func banner(text string, width int) int {
	// This have a bug because len() returns the number of bytes in a string
	// and not the rune count when used for string.
	// padding := (width - len(text)) / 2
	//To get the length in terms of rune count, we use
	// Detailed explanation at the end of this file
	padding := (width - utf8.RuneCountInString(text)) / 2

	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}

	fmt.Println(text)

	for i := 0; i < width; i++ {
		fmt.Print(".")
	}

	fmt.Println()

	return padding
}

/*

The key issue in the `banner` function within your Go code relates to how string length is calculated and the difference between bytes and runes, especially when dealing with Unicode characters like emojis.

Here's the breakdown of what's happening:

1. **String Length with `len()`**:
   - In Go, the `len()` function, when used with a string, returns the number of bytes in the string, not the number of characters (runes).
   - This is straightforward for ASCII characters, where each character is exactly one byte. However, it becomes problematic with Unicode characters, particularly those outside the ASCII range (like emojis), as they are often represented by multiple bytes.

2. **UTF-8 Encoding and Emojis**:
   - Go strings are UTF-8 encoded. UTF-8 is a variable-length encoding system, meaning that different characters can use a different number of bytes.
   - Most standard ASCII characters use 1 byte, but emojis can use several bytes. For example, the smiling face emoji (`😃`) is represented by 4 bytes in UTF-8.

3. **Issue with the Emoji in `banner("G😃", 6)`**:
   - When you call `banner("G😃", 6)`, the string "G😃" consists of 1 byte for 'G' and 4 bytes for the emoji.
   - The `len()` function would return `5` (total bytes), which is incorrect in terms of character count. It should be `2` characters ('G' and the emoji).
   - This discrepancy leads to incorrect calculation of padding if you use `len()` directly.

4. **Correct Approach with `utf8.RuneCountInString()`**:
   - To accurately count the number of characters (runes) in a string, especially when it includes multi-byte characters like emojis, you should use `utf8.RuneCountInString()`.
   - This function correctly counts the number of runes (characters) regardless of their byte size.
   - In the case of "G😃", `utf8.RuneCountInString("G😃")` correctly returns `2`, ensuring the proper calculation for padding in your banner.

In summary, the core reason behind using `utf8.RuneCountInString()` instead of `len()` in the `banner` function is to accurately count the number of characters in strings that may contain multi-byte Unicode characters, ensuring that operations based on character count (like centering text) are done correctly.


*/
