package apperror_test

import (
	"encoding/json"
	"testing"

	"github.com/Damoz1606/mini-storage-backend-golang/pkg/apperror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("Should create a new AppError when New is called", func(t *testing.T) {
		value := apperror.New(404, "not found")

		expected := apperror.AppError{
			Code:    404,
			Message: "not found",
		}

		assert.Equal(t, expected, value)
	})
}

func TestAppError_Error(t *testing.T) {
	t.Run("Should return a formatted messages when Error is called", func(t *testing.T) {
		value := apperror.New(404, "not found")

		expected := "404 - not found"
		assert.Equal(t, expected, value.Error())
	})
}

func TestAppError_Marshal(t *testing.T) {
	t.Run("Should return a formatted messages when Marshal is called", func(t *testing.T) {
		value := apperror.New(404, "not found")

		marshal := value.Marshal()
		expected := `{"code":404,"message":"not found"}`

		assert.Equal(t, expected, string(marshal))
	})

	t.Run("Should return a valid JSON when error is marshal", func(t *testing.T) {
		value := apperror.New(404, "not found")

		marshal := value.Marshal()

		var decoded apperror.AppError
		err := json.Unmarshal(marshal, &decoded)
		require.NoError(t, err)

		assert.Equal(t, decoded, value)
	})
}
