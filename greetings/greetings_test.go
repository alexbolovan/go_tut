package greetings

import (
	"regexp"
	"testing"
)

// Note: Test function must have format Test"FunctionName"

// Testing general case of name being not empty
// Verfying whether name is contained in output string
func TestHelloName(t *testing.T) {
	name := "Alex"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Alex")

	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Alex") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// Testing edgecase of name being empty
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
