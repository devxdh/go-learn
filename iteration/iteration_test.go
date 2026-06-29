package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"

	assertCorrectMessage(t, got, want)
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}
}
