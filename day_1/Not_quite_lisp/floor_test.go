package Not_quite_lisp

import "testing"

func TestFloor(t *testing.T) {
	got := Floor("(())")
	want := 0

	if got != want {
		t.Errorf("got %q want %d", got, want)
	}

	got = Floor("())")
	want = -1

	if got != want {
		t.Errorf("got %q want %d", got, want)
	}

	got = Floor("(((")
	want = 3

	if got != want {
		t.Errorf("got %q want %d", got, want)
	}

	got = Floor("(()(()(")
	want = 3

	if got != want {
		t.Errorf("got %q want %d", got, want)
	}

	got = Floor("))(((((")
	want = 3

	if got != want {
		t.Errorf("got %q want %d", got, want)
	}

	got = Floor(")())())")
	want = -3

	if got != want {
		t.Errorf("got %q want %d", got, want)
	}
}
