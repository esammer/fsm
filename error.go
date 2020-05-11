package fsm

import "fmt"

// An illegal state transition error.
//
// This error is returned when an illegal state transition is attempted. It contains the current state (From) and the
// state to which a transition was attempted (To).
type IllegalTransitionError struct {
	From State
	To   State
}

func (this *IllegalTransitionError) Error() string {
	return fmt.Sprintf("illegal state transition %v -> %v", this.From, this.To)
}
