package rando

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestNothing(t *testing.T) {}

type s struct {
	Suite
	run bool
}

func (s *s) TestThing(t *testing.T) {
	s.run = true
}

func TestThatTestsAreRun(t *testing.T) {
	ts := new(s)
	Run(t, ts)

	assert.Assert(t, ts.run)
}
