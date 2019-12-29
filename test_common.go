package arguments

import (
	"testing"
)

/*
 * Variables
 */

/*
 * Functions
 */

func match(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func noError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
}

func withError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Should got error but nil")
	}
}

/*
 * Tests
 */
