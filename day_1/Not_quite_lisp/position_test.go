package Not_quite_lisp

import "testing"

func TestPosition(t *testing.T) {
	got := Position(")")
	want := 1

	if got != want {
		t.Errorf("got %q want %d", got, want)
	}

	got = Position("()())")
	want = 5

	if got != want {
		t.Errorf("got %q want %d", got, want)
	}

	got = Position("")
	want = -1

	if got != want {
		t.Errorf("got %q want %d", got, want)
	}
}
