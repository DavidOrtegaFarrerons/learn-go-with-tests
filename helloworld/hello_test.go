package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people with name parameter", func(t *testing.T) {
		got := Hello("David", "")
		want := "Hello, David"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying 'Hello World' when parameter name is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("David", "Spanish")
		want := "Hola, David"

		assertCorrectMessage(t, got, want)
	})
	t.Run("In French", func(t *testing.T) {
		got := Hello("David", "French")
		want := "Bonjour, David"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	//Used to tell the test suite that this method is a helper,
	// so when it fails the line will be the function call rather than this helper
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
