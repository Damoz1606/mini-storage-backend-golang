package domainevent_test

import (
	"testing"

	"github.com/Damoz1606/mini-storage-backend-golang/pkg/domainevent"
	"github.com/stretchr/testify/assert"
)

const (
	testEventType domainevent.EventType = "Test"
)

var (
	testEvent = domainevent.Event{
		Type:    testEventType,
		Payload: "Hello",
	}
)

func TestRecorderRecord(t *testing.T) {
	t.Run("Should add events to the recorder when a new event is record", func(t *testing.T) {
		expected := []domainevent.Event{testEvent}
		recorder := domainevent.Recorder{}

		assert.Nil(t, recorder.UncommittedEvents())
		recorder.Record(testEvent)
		assert.NotNil(t, recorder.UncommittedEvents())
		assert.ElementsMatch(t, expected, recorder.UncommittedEvents())
	})
}

func TestRecorderUncommitedEvents(t *testing.T) {
	t.Run("Should return all the recordes when UncommittedEvents is called", func(t *testing.T) {
		expected := []domainevent.Event{testEvent, testEvent, testEvent}
		recorder := domainevent.Recorder{}

		assert.Nil(t, recorder.UncommittedEvents())
		recorder.Record(testEvent)
		recorder.Record(testEvent)
		recorder.Record(testEvent)
		assert.NotNil(t, recorder.UncommittedEvents())
		assert.ElementsMatch(t, expected, recorder.UncommittedEvents())
	})
}

func TestRecorderClearEvents(t *testing.T) {
	t.Run("Should clear all the events from the recorder when ClearEvents is called", func(t *testing.T) {
		recorder := domainevent.Recorder{}

		assert.Nil(t, recorder.UncommittedEvents())
		recorder.Record(testEvent)
		assert.NotNil(t, recorder.UncommittedEvents())
		recorder.ClearEvents()
		assert.Nil(t, recorder.UncommittedEvents())
	})
}
