package simplebdd_test

import (
	"testing"

	"github.com/draganm/go-simplebdd"
)

func TestScenario(t *testing.T) {
	simplebdd.Scenario(
		t,
	).When("step executes", func(t *testing.T) {
		t.Fail()
	}).Then("everything should be ok", func(t *testing.T) {

	})

}

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
