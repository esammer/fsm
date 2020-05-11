package fsm

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIllegalTransitionError_Error(t *testing.T) {
	err := &IllegalTransitionError{
		From: MyStateA,
		To:   MyStateB,
	}

	require.Equal(t, "illegal state transition A -> B", err.Error())
}
