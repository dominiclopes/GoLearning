package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)

	message, err := Hello(name)
	if err != nil || !want.MatchString(message) {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if err == nil || message != "" {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
	}

}
