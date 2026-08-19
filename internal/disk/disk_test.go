package disk_test

import (
	"testing"
	"time"

	"github.com/Damoz1606/mini-storage-backend-golang/internal/disk"
	"github.com/Damoz1606/mini-storage-backend-golang/pkg/domainevent"
	"github.com/Damoz1606/mini-storage-backend-golang/pkg/testkit"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	testSuite := []testkit.TestCase[string, any, *disk.Disk]{
		{
			Name: "Should return a defined Disk when valid params are passed",
			Input: func() string {
				return "Test Disk"
			},
			Setup: func(t *testing.T) any {
				return ""
			},
			Assert: func(t *testing.T, result *disk.Disk, err error) {
				t.Helper()
				assert.NotNil(t, result)
				require.NoError(t, err)
			},
		},
		{
			Name: "Should return nil and error when invalid params are passed",
			Input: func() string {
				return ""
			},
			Setup: func(t *testing.T) any {
				return ""
			},
			Assert: func(t *testing.T, result *disk.Disk, err error) {
				t.Helper()
				assert.Nil(t, result)
				assert.Error(t, err)
			},
		},
	}

	for _, tc := range testSuite {
		t.Run(tc.Name, func(t *testing.T) {
			input := tc.Input()
			value, err := disk.New(input)
			tc.Assert(t, value, err)
		})
	}
}

func TestDiskGetter(t *testing.T) {
	t.Run("Should return expected values from getters when the getters are called", func(t *testing.T) {
		input := "Test disk name"
		beforeCreateAt := time.Now().UnixMilli()
		value, err := disk.New(input)
		afterCreateAt := time.Now().UnixMilli()

		require.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, value.ID())

		assert.Equal(t, input, value.Name())

		assert.NotZero(t, value.CreateAt())
		assert.GreaterOrEqual(t, beforeCreateAt, value.CreateAt())
		assert.LessOrEqual(t, afterCreateAt, value.CreateAt())
	})
}

func TestDiskDelete(t *testing.T) {
	t.Run("Should add a time when delete is called", func(t *testing.T) {
		input := "Test disk name"
		value, err := disk.New(input)
		require.NoError(t, err)

		beforeDeleteAt := time.Now().UnixMilli()
		value.Delete()
		afterDeleteAt := time.Now().UnixMilli()

		require.NotNil(t, value.DeleteAt())
		assert.NotZero(t, value.DeleteAt())
		resultDeleteAt := value.DeleteAt()
		assert.GreaterOrEqual(t, beforeDeleteAt, *resultDeleteAt)
		assert.LessOrEqual(t, afterDeleteAt, *resultDeleteAt)

		expectedEvent := domainevent.Event{
			Type:    disk.EventDiskDeleted,
			Payload: disk.EventPayloadDiskDeleted{Value: value.ID()},
		}
		require.NotNil(t, value.UncommittedEvents())
		assert.Contains(t, value.UncommittedEvents(), expectedEvent)
	})
}
