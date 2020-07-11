package rando

import (
	"reflect"
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
	assert.Assert(t, !ts.run)
	Run(t, ts)

	assert.Assert(t, ts.run)
}

func DeferredPanicTest(t *testing.T) {
	r := recover()
	assert.Assert(t, r == nil)
}

func TestThePanic(t *testing.T) {
	defer DeferredPanicTest(t)
	panic(nil)
}

func testMock() []testing.InternalTest {
	return []testing.InternalTest{
		{
			Name: "do",
		},
		{
			Name: "re",
		},
		{
			Name: "mi",
		},
		{
			Name: "fa",
		},
		{
			Name: "so",
		},
		{
			Name: "la",
		},
		{
			Name: "ti",
		},
		{
			Name: "do",
		},
	}
}

func TestRandomization(t *testing.T) {
	target := testMock()
	comparison := testMock()

	if !reflect.DeepEqual(target, comparison) {
		t.Fail()
	}

	target = randomize(target, t)

	if reflect.DeepEqual(target, comparison) {
		t.Fail()
	}
}
