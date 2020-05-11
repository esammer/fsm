package fsm

func ExampleStateMachine() {
	sm := New().
		Allow(MyStateA, MyStateB).
		Allow(MyStateB, MyStateB).
		Allow(MyStateB, MyStateC).
		Start(MyStateA)
	smi, err := sm.NewInstance()
	if err != nil {
		panic(err)
	}

	for _, s := range []TestState{MyStateB, MyStateB, MyStateB, MyStateC} {
		if err := smi.Transition(s); err != nil {
			panic(err)
		}
	}
}
