package main

import "testing"

/*
1. Write a test
2. Make the compiler pass
3. Run the test, see that it fails and check the error message is meaningful
4. Write enough code to make the test pass
5. Refactor
*/

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// func TestHelloFriend(t *testing.T) {
// 	got := HelloFriend("Joe")
// 	want := "Hello, Joe"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// subtests
// A benefit of this approach is you can set up shared code that can be used in the other tests.
func TestHelloFriend2(t *testing.T) {
	t.Run("saying hello to a friend", func(t *testing.T) {
		got := HelloFriend("Joe", "")
		want := "Hello, Joe"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := HelloFriend("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := HelloFriend("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := HelloFriend("Elodie", "Portuguese")
		want := "Olá, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/*
The TDD process and why the steps are important
1. Write a failing test and see it fail so we know we have written a relevant test for our requirements and seen that it produces an easy to understand description of the failure
2. Writing the smallest amount of code to make it pass so we know we have working software
3. Then refactor, backed with the safety of our tests to ensure we have well-crafted code that is easy to work with
*/
