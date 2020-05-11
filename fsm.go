package fsm

import (
	"errors"
)

// A state.
//
// Typically states are identified by Go-style enums, but may take any form.
type State interface {
}

var ErrIllegalStateMachine = errors.New("illegal state machine - no start state")

// A state machine.
type StateMachine struct {
	transitions map[State]map[State]bool
	start       State
}

// Return a new instance of a state machine.
func New() *StateMachine {
	return &StateMachine{
		transitions: make(map[State]map[State]bool),
	}
}

// Configure the machine to allow the given state transition.
//
// Returns the state machine.
func (this *StateMachine) Allow(from State, to State) *StateMachine {
	m, ok := this.transitions[from]
	if !ok {
		m = make(map[State]bool)
		this.transitions[from] = m
	}

	m[to] = true

	return this
}

// Set the initial state.
//
// All state machines must have a starting state.
//
// Returns the state machine.
func (this *StateMachine) Start(s State) *StateMachine {
	this.start = s

	return this
}

// Determine is the given state transition is allowed.
func (this *StateMachine) IsAllowed(from State, to State) bool {
	if m, ok := this.transitions[from]; ok {
		if _, ok := m[to]; ok {
			return true
		}
	}

	return false
}

// Create a new instance of the state machine.
//
// If the state machine itself is invalid (i.e. does not have a start state), an error is returned along with a nil
// instance.
//
// Returns an instance of the state machine at the configured start state.
func (this *StateMachine) NewInstance() (*Instance, error) {
	if this.start == nil {
		return nil, ErrIllegalStateMachine
	}

	return &Instance{
		StateMachine: this,
		CurrentState: this.start,
	}, nil
}
