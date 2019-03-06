package errlist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	one := fmt.Errorf("One")
	two := fmt.Errorf("Two")

	t.Run("when the input list is nil", func(t *testing.T) {
		assert.NoError(t, Error(nil))
	})

	t.Run("when the input list is empty", func(t *testing.T) {
		assert.NoError(t, Error([]error{}))
	})

	t.Run("when the input list consists of nils only", func(t *testing.T) {
		assert.Nil(t, Append(nil, nil, nil))
	})

	t.Run("should make a list of errors", func(t *testing.T) {
		assert.Equal(t, []error{one, two}, Append(nil, one, two))

		t.Run("should drop nils", func(t *testing.T) {
			assert.Equal(t, []error{one, two}, Append(nil, one, nil, two, nil))
		})
	})

	t.Run("message of a single error", func(t *testing.T) {
		assert.Equal(t, "One", Error([]error{one}).Error())
	})

	t.Run("message of a list of errors", func(t *testing.T) {
		assert.Equal(t, "One; Two", Error([]error{one, two}).Error())
	})
}
