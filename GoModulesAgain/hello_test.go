package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, World."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %v, want = %v", got, want)
	}
}
