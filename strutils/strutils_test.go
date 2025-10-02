package strutils

import (
	"testing"
)

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalmap(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func TestReverse(t *testing.T) {
	got := Reverse("hello")
	want := "olleh"
	if got != want {
		t.Errorf("Reverse(\"hello\") = %q; want %q", got, want)
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := map[string]bool{
		"racecar": true,
		"madam":   true,
		"hello":   false,
	}
	for input, want := range cases {
		if got := IsPalindrome(input); got != want {
			t.Errorf("IsPalindrome(%q) = %v; want %v", input, got, want)
		}
	}
}

func TestCharCount(t *testing.T) {
	got := CharCount("hello")
	want := map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("CharCount(\"hello\")[%q] = %d; want %d", k, got[k], v)
		}
	}
}

func TestRuneCount(t *testing.T) {
	got := RuneCount("🙂🙂")
	want := 2
	if got != want {
		t.Errorf("RuneCount(%q) = %d; want %d", "🙂🙂", got, want)
	}
}

func TestFristRune(t *testing.T) {
	got := FirstRune("hello")
	want := 'h'
	if got != want {
		t.Errorf("FirstRune(%q) = %q; want %q", "hello", string(got), string(want))
	}
}

func TestLastRune(t *testing.T) {
	got := LastRune("Maddox")
	want := 'x'
	if got != want {
		t.Errorf("LastRune(%q)) = %q; want %q", "Maddox", string(got), string(want))
	}
}

func TestNthRune(t *testing.T) {
	got := NthRune("hello", 1)
	want := 'e'
	if got != want {
		t.Errorf("NthRune(%q, %d) = %q; want %q", "hello", 1, string(got), string(want))
	}
}

func TestSubstringRunes(t *testing.T) {
	got := SubstringRunes("hello", 1, 4)
	want := "ell"
	if got != want {
		t.Errorf("SubstringRunes(%q, %d, %d) = %q; want %q", "hello", 1, 4, got, want)
	}
}

func TestToUpper(t *testing.T) {
	got := ToUpper("hello")
	want := "Hello"
	if got != want {
		t.Errorf("ToUpper(%q)= %q want %q", "hello", string(got), string(want))
	}
}

func TestTwoSum(t *testing.T) {
	got := TwoSum([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}
	if !equalSlice(got, want) {
		t.Errorf("TwoSum() = %v; want %v", got, want)
	}
}

func TestFrequencyCounter(t *testing.T) {
	got := FrequencyCounter("hello")
	want := map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1}
	if !equalmap(got, want) {
		t.Errorf("FrequencyCounter() = %v, want %v", got, want)
	}
}
