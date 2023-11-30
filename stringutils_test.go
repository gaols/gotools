package gotools

import (
	"testing"
)

func TestIsBlank(t *testing.T) {
	if !IsBlank("") {
		t.FailNow()
	}
	if !IsBlank(" ") {
		t.FailNow()
	}
	if !IsBlank(" \t\n") {
		t.FailNow()
	}
	if IsBlank("h") {
		t.FailNow()
	}
}

func TestIsEmpty(t *testing.T) {
	if !IsEmpty("") {
		t.FailNow()
	}
	if IsEmpty("a") {
		t.FailNow()
	}
}

func TestDefaultIfEmpty(t *testing.T) {
	if DefaultIfEmpty("", "foo") != "foo" {
		t.FailNow()
	}
}

func TestDefaultIfBlank(t *testing.T) {
	if DefaultIfBlank(" \t\r", "foo") != "foo" {
		t.FailNow()
	}
}

func TestIsNotEmpty(t *testing.T) {
	if !IsNotEmpty(" ") || !IsNotEmpty("\n") || !IsNotEmpty("\r") || !IsNotEmpty("\t") {
		t.FailNow()
	}
}

func TestIsNotBlank(t *testing.T) {
	if IsNotBlank(" ") {
		t.FailNow()
	}
}

func TestRightPad(t *testing.T) {
	if RightPad("hello", 7, ' ') != "hello  " {
		t.FailNow()
	}
	if RightPad("猴王", 7, '王') != "猴王王王王王王" {
		t.FailNow()
	}
	if RightPad("猴王", 7, ' ') != "猴王     " {
		t.FailNow()
	}
	if RightPad("猴王", 2, ' ') != "猴王" {
		t.FailNow()
	}
	if RightPad("猴王", 1, ' ') != "猴王" {
		t.FailNow()
	}
	if RightPad("猴王", 0, ' ') != "猴王" {
		t.FailNow()
	}
	if RightPad("猴王", -1, ' ') != "猴王" {
		t.FailNow()
	}
	if RightPad("ab", 2, ' ') != "ab" {
		t.FailNow()
	}
	if RightPad("ab", 1, ' ') != "ab" {
		t.FailNow()
	}
	if RightPad("ab", 0, ' ') != "ab" {
		t.FailNow()
	}
	if RightPad("ab", -1, ' ') != "ab" {
		t.FailNow()
	}
	if RightPad("hello,猴王 ", 10, '强') != "hello,猴王 强" {
		t.FailNow()
	}
}

func TestLeftPad(t *testing.T) {
	if LeftPad("hello", 7, ' ') != "hello" {
		t.FailNow()
	}
	if LeftPad("猴王", 7, '王') != "王王王王王猴王" {
		t.FailNow()
	}
	if LeftPad("猴王", 7, ' ') != "     猴王" {
		t.FailNow()
	}
	if LeftPad("猴王", 2, ' ') != "猴王" {
		t.FailNow()
	}
	if LeftPad("猴王", 1, ' ') != "猴王" {
		t.FailNow()
	}
	if LeftPad("猴王", 0, ' ') != "猴王" {
		t.FailNow()
	}
	if LeftPad("猴王", -1, ' ') != "猴王" {
		t.FailNow()
	}
	if LeftPad("ab", 2, ' ') != "ab" {
		t.FailNow()
	}
	if LeftPad("ab", 1, ' ') != "ab" {
		t.FailNow()
	}
	if LeftPad("ab", 0, ' ') != "ab" {
		t.FailNow()
	}
	if LeftPad("ab", -1, ' ') != "ab" {
		t.FailNow()
	}
	if LeftPad("hello,猴王 ", 10, '强') != "强hello,猴王 " {
		t.FailNow()
	}
}

func TestReverse(t *testing.T) {
	if Reverse("hello") != "olleh" {
		t.FailNow()
	}
	if Reverse("落霞与孤鹜齐飞") != "飞齐鹜孤与霞落" {
		t.FailNow()
	}
}

func TestReversePreservingCombiningCharacters(t *testing.T) {
	if ReversePreservingCombiningCharacters("The quick bròwn 狐 jumped over the lazy 犬") != "犬 yzal eht revo depmuj 狐 nwòrb kciuq ehT" {
		t.FailNow()
	}
}

func TestIsAnyBlank(t *testing.T) {
	if IsAnyBlank("a", "b") {
		t.FailNow()
	}
	if !IsAnyBlank("a", "") {
		t.FailNow()
	}
	if IsAnyBlank("a") {
		t.FailNow()
	}
}

func TestSubstring(t *testing.T) {
	if Substring("hello", -1, -1) != "hello" {
		t.FailNow()
	}
	if Substring("hello", 0, -1) != "hello" {
		t.FailNow()
	}
	if Substring("hello", -1, 10) != "hello" {
		t.FailNow()
	}
	if Substring("hello", 10, 11) != "" {
		t.FailNow()
	}
	if Substring("hello", 4, 3) != "" {
		t.FailNow()
	}
	if Substring("hello", 3, 4) != "l" {
		t.FailNow()
	}
	if Substring("hello", 3, 3) != "" {
		t.FailNow()
	}
	if Substring("hello", 0, 1) != "h" {
		t.FailNow()
	}
}

func TestIsEqualsAny(t *testing.T) {
	if !IsEqualsAny("", "") {
		t.FailNow()
	}
	if IsEqualsAny("") {
		t.FailNow()
	}
	if !IsEqualsAny("a", "a", "b") {
		t.FailNow()
	}
	if !IsEqualsAny("a", "a") {
		t.FailNow()
	}
	if IsEqualsAny("a", "b") {
		t.FailNow()
	}
}
