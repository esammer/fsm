package fsm

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type TestState int

const (
	MyStateA TestState = iota
	MyStateB
	MyStateC
)

var myStateStrs = []string{
	"A",
	"B",
	"C",
}

func (this TestState) String() string {
	return myStateStrs[this]
}

func TestStateMachine_Allow(t *testing.T) {
	tests := []struct {
		name string
		sm   *StateMachine
		// A series of test cases. The first map key is the from state, the second is the to state. The value indicates
		// if the transition is expected to be allowed.
		expected map[State]map[State]bool
	}{
		{
			name: "empty",
			sm:   New(),
			expected: map[State]map[State]bool{
				MyStateA: {
					MyStateB: false,
					MyStateC: false,
				},
				MyStateB: {
					MyStateB: false,
					MyStateC: false,
				},
			},
		},
		{
			name: "one source/one dest",
			sm: New().
				Allow(MyStateA, MyStateB),
			expected: map[State]map[State]bool{
				MyStateA: {
					MyStateB: true,
					MyStateC: false,
				},
				MyStateB: {
					MyStateB: false,
					MyStateC: false,
				},
			},
		},
		{
			name: "two source/one dest",
			sm: New().
				Allow(MyStateA, MyStateB).Allow(MyStateB, MyStateC),
			expected: map[State]map[State]bool{
				MyStateA: {
					MyStateB: true,
					MyStateC: false,
				},
				MyStateB: {
					MyStateB: false,
					MyStateC: true,
				},
			},
		},
		{
			name: "one source/two dest",
			sm: New().
				Allow(MyStateA, MyStateB).
				Allow(MyStateA, MyStateC),
			expected: map[State]map[State]bool{
				MyStateA: {
					MyStateB: true,
					MyStateC: true,
				},
				MyStateB: {
					MyStateB: false,
					MyStateC: false,
				},
			},
		},
		{
			name: "two source/two dest",
			sm: New().
				Allow(MyStateA, MyStateB).
				Allow(MyStateA, MyStateC).
				Allow(MyStateB, MyStateB).
				Allow(MyStateB, MyStateC),
			expected: map[State]map[State]bool{
				MyStateA: {
					MyStateB: true,
					MyStateC: true,
				},
				MyStateB: {
					MyStateB: true,
					MyStateC: true,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for from, tos := range test.expected {
				for to, expected := range tos {
					require.Equal(t, expected, test.sm.IsAllowed(from, to))
				}
			}
		})
	}
}

func TestStateMachine_Start(t *testing.T) {
	sm := New()
	require.Equal(t, nil, sm.start)

	sm.Start(MyStateA)
	require.Equal(t, MyStateA, sm.start)
}

func TestStateMachine_NewInstance(t *testing.T) {
	sm := New().
		Allow(MyStateA, MyStateB).
		Start(MyStateA)

	smi, err := sm.NewInstance()
	require.NoError(t, err)
	require.NotNil(t, smi)
	require.Equal(t, smi.StateMachine, sm)
	require.Equal(t, smi.CurrentState, MyStateA)
}

func TestStateMachine_NewInstance_InvalidStart(t *testing.T) {
	sm := New()

	smi, err := sm.NewInstance()
	require.Equal(t, ErrIllegalStateMachine, err)
	require.Nil(t, smi)
}
