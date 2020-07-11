package rando

import (
	"flag"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
)

var seed = flag.Int64("rando.seed", 0, "Specify a seed for the randomization of test execution")

// TestSuite is the interface we use to ensure that all consumers of the package
// are using the actual Suite struct in their testsuite struct.
type TestSuite interface {
	T() *testing.T
}

// Suite is the reified struct to embed in your TestSuite
type Suite struct {
	t *testing.T
}

// T returns the *testing.T in the Suite
func (s *Suite) T() *testing.T {
	return s.t
}

// Run takes a *testing.T and variadic number of TestSuites, derives all the tests
// that start with Test (like the go test tool), randomizes all the tests, and
// then runs them.
func Run(t *testing.T, ts ...TestSuite) {
	var tests []testing.InternalTest

	for _, s := range ts {
		tests = append(tests, getTests(s)...)
	}

	tests = randomize(tests, t)
	for _, test := range tests {
		t.Run(test.Name, test.F)
	}
}

func getTests(s TestSuite) []testing.InternalTest {
	tester := reflect.TypeOf(s)

	var tests []testing.InternalTest

	for i := 0; i < tester.NumMethod(); i++ {
		method := tester.Method(i)

		if strings.HasPrefix(method.Name, "Test") {
			test := testing.InternalTest{
				Name: method.Name,
				F: func(t *testing.T) {
					values := []reflect.Value{
						reflect.ValueOf(s), reflect.ValueOf(t),
					}
					method.Func.Call(values)
				},
			}
			tests = append(tests, test)
		}
	}
	return tests
}

func randomize(tests []testing.InternalTest, t *testing.T) []testing.InternalTest {
	var rando int64
	if *seed == 0 {
		rando = time.Now().UnixNano()
	} else {
		rando = *seed
	}
	t.Logf("Randomizing tests using seed %d", seed)
	rand.Seed(rando)

	rand.Shuffle(len(tests), func(i, j int) {
		tests[i], tests[j] = tests[j], tests[i]
	})
	return tests
}
