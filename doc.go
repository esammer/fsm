// The fsm package contains a simple programmable state machine. A StateMachine is configured once with a set of valid
// transitions and an initial or start state. Instances of the state machine are then created, and track current state
// as well as validate any attempted transitions. The StateMachine / Instance relationship is analogous to a compiled
// regex and match attempts.
//
// The State interface represents a state and can be any type. Go-style enums are highly recommended (see the State
// example).
package fsm
