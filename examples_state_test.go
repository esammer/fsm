package fsm

// Our special state type.
type SpecialState int

// Transition errors will contain helpful text if our State implements fmt.Stringer.
func (s SpecialState) String() string {
	return specialStateStrings[s]
}

// Define our valid states.
const (
	SpecialStateNotRunning SpecialState = iota
	SpecialStateStarting
	SpecialStateRunning
	SpecialStateStopping
	SpecialStateStopped
	SpecialStateEnd
)

// Define string representations for each state.
var specialStateStrings = []string{
	"NOT_RUNNING",
	"STARTING",
	"RUNNING",
	"STOPPING",
	"STOPPED",
	"DONE",
}

func ExampleState() {
	// Create an instance of the state machine with valid state transitions.
	sm := New().
		Allow(SpecialStateNotRunning, SpecialStateStarting).
		Allow(SpecialStateStarting, SpecialStateRunning).
		Allow(SpecialStateRunning, SpecialStateStopping).
		Allow(SpecialStateStopping, SpecialStateStopped).
		Allow(SpecialStateStopped, SpecialStateEnd).
		Start(SpecialStateNotRunning)

	// Create a new instance of the state machine.
	smi, err := sm.NewInstance()
	if err != nil {
		panic(err)
	}

	// Attempt to transition through a series of states.
	for _, s := range []SpecialState{SpecialStateStarting, SpecialStateRunning, SpecialStateStopping, SpecialStateStopped, SpecialStateEnd} {
		smi.MustTransition(s)
	}
}
