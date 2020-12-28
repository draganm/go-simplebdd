package simplebdd

import (
	"fmt"
	"testing"
)

type NextStep interface {
	Given(description string, sf func(*testing.T)) NextStep
	When(description string, sf func(*testing.T)) NextStep
	Then(description string, sf func(*testing.T)) NextStep
	And(description string, sf func(*testing.T)) NextStep
}

type noopStep struct{}

func (ns noopStep) Given(description string, sf func(*testing.T)) NextStep {
	return ns
}

func (ns noopStep) When(description string, sf func(*testing.T)) NextStep {
	return ns
}

func (ns noopStep) Then(description string, sf func(*testing.T)) NextStep {
	return ns
}

func (ns noopStep) And(description string, sf func(*testing.T)) NextStep {
	return ns
}

type SimpleScenario struct {
	t *testing.T
}

func Scenario(t *testing.T) NextStep {
	return &SimpleScenario{
		t: t,
	}
}

func (s *SimpleScenario) step(description string, tf func(t *testing.T)) NextStep {
	failed := false
	s.t.Run(fmt.Sprint(description), func(t *testing.T) {
		defer func() {
			pn := recover()
			if pn != nil {
				failed = true
				panic(pn)
			}
		}()
		tf(t)
		if t.Failed() {
			failed = true
		}
	})

	if failed {
		return noopStep{}
	}

	return s
}

func (s *SimpleScenario) Given(description string, tf func(t *testing.T)) NextStep {
	return s.step(fmt.Sprintf("Given %s", description), tf)
}

func (s *SimpleScenario) When(description string, tf func(t *testing.T)) NextStep {
	return s.step(fmt.Sprintf("When %s", description), tf)
}

func (s *SimpleScenario) Then(description string, tf func(t *testing.T)) NextStep {
	return s.step(fmt.Sprintf("Then %s", description), tf)
}

func (s *SimpleScenario) And(description string, tf func(t *testing.T)) NextStep {
	return s.step(fmt.Sprintf("And %s", description), tf)
}
