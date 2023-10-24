package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("名前が入った時", func(t *testing.T) {
		got := Hello("ryo")
		want := "Hello ryo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("名前が空のとき", func(t *testing.T) {
		got := Hello("")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})

}
