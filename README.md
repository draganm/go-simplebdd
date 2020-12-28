# simple BDD for Golang

This library is a simple syntactic sugar that utilizes existing Golang testing framework and add BDD style syntax on top of it.


# Motivation
I've been missing something that will keep me close to the Golang tests, but at the same time enable me to write BDD style Given/When/Then tests.
There are heavy weight approaches to this such as [Ginkgo](https://github.com/onsi/ginkgo) and I've been using it for years, but then I've realized that my tests are anything but readable for people who are not used to RSpec style of tests.

# Why not just using Golang Sub-Tests
I've been writing tests of the form

```golang
func TestAdd(t *testing.T) {

	var result int

	t.Run("When I add 2 and 3", func(t *testing.T) {
		result = 2 + 3
	})

	t.Run("Then the result should be positive", func(t *testing.T) {
		if result < 0 {
			t.Fatal("not positive")
		}
	})

	t.Run("And the result should be 5", func(t *testing.T) {
		if result != 5 {
			t.Fatal("not equal 5")
		}
	})

}
```

and it works nicely, but if for some reason the step `Then the result should be positive` fails, the step `And the result should be 5` will still be executed, which is ugly to say the least and might cost you time waiting for the steps to time out.
