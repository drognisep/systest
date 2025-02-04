package runtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestState_RunDir(t *testing.T) {
	s := New()
	assert.NoError(t, s.RunDir("testing"))
}
