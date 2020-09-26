package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "alony"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello(%q) = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestHellos(t *testing.T) {
	name := "alony"
	want := regexp.MustCompile(`\b` + name + `\b`)
	ans, err := Hellos(name)
	if !want.MatchString(ans[name]) || err != nil {
		t.Fatalf(`Hello(%q) = %q, %v, want match for %#q, nil`, name, ans[name], err, want)
	}
}
