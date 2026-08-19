package domainevent

type EventType string

type Event struct {
	Type    EventType
	Payload any
}

type Recorder struct {
	events []Event
}

func (r *Recorder) Record(event Event) {
	r.events = append(r.events, event)
}

func (r Recorder) UncommittedEvents() []Event {
	return r.events
}

func (r *Recorder) ClearEvents() {
	r.events = nil
}
