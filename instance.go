package fsm

// An instance of a state machine.
type Instance struct {
	StateMachine *StateMachine
	CurrentState State
}

// Transition to the given state from the current state.
//
// If the transition is not allowed by the state machine, IllegalTransitionError is produced.
func (this *Instance) Transition(to State) error {
	if !this.StateMachine.IsAllowed(this.CurrentState, to) {
		return &IllegalTransitionError{
			From: this.CurrentState,
			To:   to,
		}
	}

	this.CurrentState = to

	return nil
}

// Like Transition(State) but panic on error.
//
// This method is a convenience for cases where a failure to transition means there is a panic-worthy bug in the
// program.
func (this *Instance) MustTransition(to State) {
	if err := this.Transition(to); err != nil {
		panic(err)
	}
}
