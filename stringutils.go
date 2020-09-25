package gotools

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// IsBlank Checks if a string is whitespace, empty ("").
func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// IsNotBlank Checks if a string is not empty (""), not null and not whitespace only.
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

// IsEmpty Checks if a string is whitespace, empty ("").
func IsEmpty(str string) bool {
	return len(str) == 0
}

// IsNotEmpty Checks if a string is not empty ("").
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// IsAnyBlank checks if any one of the CharSequences are blank ("") or whitespace only.
func IsAnyBlank(strList ...string) bool {
	if len(strList) == 0 {
		panic("no param found")
	}

	for _, v := range strList {
		if IsBlank(v) {
			return true
		}
	}
	return false
}

// Trim returns a slice of the string s, with all leading and trailing white space removed, as defined by Unicode.
func Trim(str string) string {
	return strings.TrimSpace(str)
}

// DefaultIfEmpty Returns either the passed in string, or if the string is
// empty (""), the value of default string.
func DefaultIfEmpty(str, defaultStr string) string {
	if IsEmpty(str) {
		return defaultStr
	}

	return str
}

// DefaultIfBlank Returns either the passed in string, or if the string is
// whitespace, empty (""), the value of default string.
func DefaultIfBlank(str, defaultStr string) string {
	if IsBlank(str) {
		return defaultStr
	}

	return str
}

// LeftPad pad a String with a specified character on the left.
// WARNING: string to pad should be uft8-encoded!
//
// goutils.LeftPad("", 3, 'z')     = "zzz"
// goutils.LeftPad("bat", 3, 'z')  = "bat"
// goutils.LeftPad("bat", 5, 'z')  = "zzbat"
// goutils.LeftPad("bat", 1, 'z')  = "bat"
// goutils.LeftPad("bat", -1, 'z') = "bat"
func LeftPad(str string, size int, padChar rune) string {
	return calcPadStr(str, size, padChar) + str
}

// RightPad right pad a String with a specified character.
// WARNING: string to pad should be uft8-encoded!
//
// goutils.RightPad("", 3, 'z')     = "zzz"
// goutils.RightPad("bat", 3, 'z')  = "bat"
// goutils.RightPad("bat", 5, 'z')  = "batzz"
// goutils.RightPad("bat", 1, 'z')  = "bat"
// goutils.RightPad("bat", -1, 'z') = "bat"
func RightPad(str string, size int, padChar rune) string {
	return str + calcPadStr(str, size, padChar)
}

func calcPadStr(str string, size int, padChar rune) string {
	count := utf8.RuneCountInString(str)
	if size <= count {
		return ""
	}
	return strings.Repeat(string(padChar), size-count)
}

// Reverse reverse a string.
// WARNING: This does not work with combining characters. check
// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go for more stories.
//
// goutils.Reverse("hello") = "olleh"
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ReversePreservingCombiningCharacters reverse a string preserving combining characters.
// The implementation is copied from http://rosettacode.org/wiki/Reverse_a_string#Go
// goutils.ReversePreservingCombiningCharacters("The quick bròwn 狐 jumped over the lazy 犬") = "犬 yzal eht revo depmuj 狐 nwòrb kciuq ehT"
func ReversePreservingCombiningCharacters(s string) string {
	if s == "" {
		return ""
	}
	p := []rune(s)
	r := make([]rune, len(p))
	start := len(r)
	for i := 0; i < len(p); {
		// quietly skip invalid UTF-8
		if p[i] == utf8.RuneError {
			i++
			continue
		}
		j := i + 1
		for j < len(p) && (unicode.Is(unicode.Mn, p[j]) ||
			unicode.Is(unicode.Me, p[j]) || unicode.Is(unicode.Mc, p[j])) {
			j++
		}
		for k := j - 1; k >= i; k-- {
			start--
			r[start] = p[k]
		}
		i = j
	}
	return string(r[start:])
}

// Substring Returns a substring of str in range(i, j).
func Substring(str string, i, j int) string {
	runes := []rune(str)
	if i >= len(str) {
		return ""
	}
	if j > len(str) {
		j = len(str)
	}

	ld := i >= 0
	rd := j >= 0
	if ld && rd {
		if j <= i {
			return ""
		}
		return string(runes[i:j])
	}

	if ld {
		return string(runes[i:])
	} else if rd {
		return string(runes[:j])
	}

	return str
}

// IsEqualsAny tests whether a string equals any string provided.
func IsEqualsAny(val string, vals ...string) bool {
	for _, v := range vals {
		if val == v {
			return true
		}
	}
	return false
}
