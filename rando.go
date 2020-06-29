package rando

import (
	"flag"
	"reflect"
	"strings"
	"testing"
)

var seed = flag.Int64("rando.seed", 0, "Specify a seed for the randomization of test execution")

// TestSuite is the interface we use to ensure that all consumers of the package
// are using the actual Suite struct in their testsuite struct.
type TestSuite interface {
	T() *testing.T
}

type Suite struct {
	t *testing.T
}

func (s *Suite) T() *testing.T {
	return s.t
}

func Run(t *testing.T, s TestSuite) {
	// 1. Get test methods
	// 2. Randomize the order
	// 3. Run the tests

	tester := reflect.TypeOf(s)

	var tests []testing.InternalTest

	for i := 0; i < tester.NumMethod(); i++ {
		method := tester.Method(i)

		if strings.HasPrefix(method.Name, "Test") {
			test := testing.InternalTest{
				Name: method.Name,
				F: func(t *testing.T) {
					values := []reflect.Value{
						reflect.ValueOf(s),
						reflect.ValueOf(t),
					}
					method.Func.Call(values)
				},
			}
			tests = append(tests, test)
		}
	}
	for _, te := range tests {
		t.Run(te.Name, te.F)
	}
}
