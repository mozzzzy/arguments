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

/*
 * Tests
 */

func TestValidateIntMin(t *testing.T) {
	t.Run("Validate IntMin without error (-5 >= -5)", func(t *testing.T) {
		var value *int = new(int)
		*value = -5
		err := ValidateIntMin("somekey", value, []int{-5})
		noError(t, err)
	})

	t.Run("Validate IntMin with error (-6 >= -5)", func(t *testing.T) {
		var value *int = new(int)
		*value = -6
		err := ValidateIntMin("somekey", value, []int{-5})
		withError(t, err)
	})
}

func TestValidateIntMax(t *testing.T) {
	t.Run("Validate IntMax without error (5 >= 5)", func(t *testing.T) {
		var value *int = new(int)
		*value = 5
		err := ValidateIntMin("somekey", value, []int{5})
		noError(t, err)
	})

	t.Run("Validate IntMax with error (5 >= 6)", func(t *testing.T) {
		var value *int = new(int)
		*value = 6
		err := ValidateIntMin("somekey", value, []int{5})
		noError(t, err)
	})
}

func TestValidateIntMinMax(t *testing.T) {
	t.Run("Validate IntMinMax without error (5 >= 5 >= -5)", func(t *testing.T) {
		var value *int = new(int)
		*value = 5
		err := ValidateIntMinMax("somekey", value, []int{-5, 5})
		noError(t, err)
	})

	t.Run("Validate IntMinMax without error (5 >= -5 >= -5)", func(t *testing.T) {
		var value *int = new(int)
		*value = -5
		err := ValidateIntMinMax("somekey", value, []int{-5, 5})
		noError(t, err)
	})

	t.Run("Validate IntMinMax with error (5 >= 6 >= -5)", func(t *testing.T) {
		var value *int = new(int)
		*value = 6
		err := ValidateIntMinMax("somekey", value, []int{-5, 5})
		withError(t, err)
	})

	t.Run("Validate IntMinMax with error (5 >= -6 >= -5)", func(t *testing.T) {
		var value *int = new(int)
		*value = -6
		err := ValidateIntMinMax("somekey", value, []int{-5, 5})
		withError(t, err)
	})
}
