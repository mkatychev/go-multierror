package multierror

import (
	"errors"
	"testing"
)

func TestListFormatFuncSingle(t *testing.T) {
	t.Run("Flat", func(t *testing.T) {
		expected := `foo`

		errors := []error{
			errors.New("foo"),
		}

		actual := ListFormatFunc(errors)
		if actual != expected {
			t.Fatalf("bad: %#v", actual)
		}
	})

	t.Run("Nested", func(t *testing.T) {
		expected := `foo`

		nestedErrors := &Error{
			Errors: []error{
				&Error{Errors: []error{errors.New("foo")}},
			},
		}

		actual := ListFormatFunc(nestedErrors.Errors)
		if actual != expected {
			t.Fatalf("bad: %#v", actual)
		}
	})
}

func TestListFormatFuncMultiple(t *testing.T) {
	expected := `2 errors occurred:
	* foo
	* bar

`

	errors := []error{
		errors.New("foo"),
		errors.New("bar"),
	}

	actual := ListFormatFunc(errors)
	if actual != expected {
		t.Fatalf("bad: %#v", actual)
	}
}
