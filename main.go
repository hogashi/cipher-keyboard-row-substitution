package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 1234567890
// qwertyuiop
// asdfghjkl;
// zxcvbnm,./
var keyboardNextChar = map[rune]rune{
	'1': 'q', '2': 'w', '3': 'e', '4': 'r', '5': 't', '6': 'y', '7': 'u', '8': 'i', '9': 'o', '0': 'p',
	'q': 'a', 'w': 's', 'e': 'd', 'r': 'f', 't': 'g', 'y': 'h', 'u': 'j', 'i': 'k', 'o': 'l', 'p': ';',
	'a': 'z', 's': 'x', 'd': 'c', 'f': 'v', 'g': 'b', 'h': 'n', 'j': 'm', 'k': ',', 'l': '.', ';': '/',
	'z': '1', 'x': '2', 'c': '3', 'v': '4', 'b': '5', 'n': '6', 'm': '7', ',': '8', '.': '9', '/': '0',
}

func substitution(r rune) rune {
	return keyboardNextChar[r]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(strings.Map(substitution, text))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
