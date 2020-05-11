package fsm

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInstance(t *testing.T) {
	sm := New().
		Allow(MyStateA, MyStateA).
		Allow(MyStateA, MyStateB).
		Allow(MyStateB, MyStateB).
		Allow(MyStateB, MyStateC).
		Start(MyStateA)

	tests := []struct {
		name        string
		transitions []State
		expectedErr error
		errIdx      int
	}{
		{
			name:        "a, b, c",
			transitions: []State{MyStateA, MyStateB, MyStateC},
		},
		{
			name:        "a, b, b, c",
			transitions: []State{MyStateA, MyStateB, MyStateB, MyStateC},
		},
		{
			name:        "a, b, c, a(err)",
			transitions: []State{MyStateA, MyStateB, MyStateC, MyStateA},
			expectedErr: &IllegalTransitionError{
				From: MyStateC,
				To:   MyStateA,
			},
			errIdx: 3,
		},
		{
			name:        "c(err)",
			transitions: []State{MyStateC},
			expectedErr: &IllegalTransitionError{
				From: MyStateA,
				To:   MyStateC,
			},
			errIdx: 0,
		},
	}

	// We loop twice just to keep test runs together in output.

	for _, test := range tests {
		t.Run("Transition/"+test.name, func(t *testing.T) {
			smi, err := sm.NewInstance()
			require.NoError(t, err)
			require.NotNil(t, smi)

			for i, s := range test.transitions {
				err := smi.Transition(s)
				if test.expectedErr != nil && i == test.errIdx {
					require.Equal(t, test.expectedErr, err)
				} else {
					require.NoError(t, err)
				}
			}
		})
	}

	for _, test := range tests {
		t.Run("MustTransition/"+test.name, func(t *testing.T) {
			smi, err := sm.NewInstance()
			require.NoError(t, err)
			require.NotNil(t, smi)

			defer func() {
				err := recover()
				require.Equal(t, test.expectedErr, err)
			}()

			for _, s := range test.transitions {
				smi.MustTransition(s)
			}
		})
	}
}

func BenchmarkInstance(b *testing.B) {
	sm := New().
		Allow(MyStateA, MyStateA).
		Allow(MyStateA, MyStateB).
		Allow(MyStateB, MyStateB).
		Allow(MyStateB, MyStateC).
		Start(MyStateA)

	b.Run("NewInstance", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := sm.NewInstance()
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Transition", func(b *testing.B) {
		smi, err := sm.NewInstance()
		if err != nil {
			b.Fatal(err)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Ignore errors.
			_ = smi.Transition(MyStateB)
		}
	})

	b.Run("MustTransition", func(b *testing.B) {
		smi, err := sm.NewInstance()
		if err != nil {
			b.Fatal(err)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			smi.MustTransition(MyStateB)
		}
	})
}
