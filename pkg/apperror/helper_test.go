package apperror_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Damoz1606/mini-storage-backend-golang/pkg/apperror"
	"github.com/stretchr/testify/assert"
)

const mockedMessage = "mocked message"

func TestBaseRequest(t *testing.T) {
	t.Run("Should return an apperror with code 400 and a message when BadRequest is called", func(t *testing.T) {
		err := apperror.BadRequest(mockedMessage)

		assert.Equal(t, http.StatusBadRequest, err.Code)
		assert.Equal(t, mockedMessage, err.Message)

	})
}

func TestIsValid(t *testing.T) {
	t.Run("Should return an appError and true when IsValid is called with appError", func(t *testing.T) {
		err := apperror.New(400, mockedMessage)

		value, ok := apperror.IsValid(err)

		assert.True(t, ok)
		assert.Equal(t, err, value)
	})

	t.Run("Should return an invalid error and false when IsValid is called with invalid apperror", func(t *testing.T) {
		err := errors.New(mockedMessage)

		value, ok := apperror.IsValid(err)

		assert.False(t, ok)
		assert.NotEqual(t, err, value)
	})
}
