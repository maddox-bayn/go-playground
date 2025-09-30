package main

import "testing"

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
