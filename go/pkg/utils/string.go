package utils

import (
	"math/rand"
	"strings"
	"time"
)

// String, Luv util for strings
type String struct{}

const EMPTY string = ""
const NewLine string = "\n"

// StringReverse, which takes a string as argument and return the reversed version
func (*String) StringReverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

// StringRandomCharacter, which takes a string as argument and return a rand character from it
func (*String) StringRandomCharacter(s string) (result string) {
	if len(s) > 0 {
		rand.Seed(time.Now().UnixNano())
		chars := []rune(s)
		result = string(chars[rand.Intn(len(chars))])
	}
	return
}

// StringSplitFixedSize, split string with fixed size in golang
func (sp *String) StringSplitFixedSize(s string, size int) []string {
	if len(s) <= size {
		return []string{s}
	}
	return append([]string{string(s[0:size])}, sp.StringSplitFixedSize(s[size:], size)...)
}

// StringTail, which takes 2 strings, one as value and the other the
//contained part and return the tail part after the contained
func StringTail(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:]
}
