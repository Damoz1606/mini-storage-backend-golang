package disk

import (
	"strings"
	"time"

	"github.com/Damoz1606/mini-storage-backend-golang/pkg/apperror"
	"github.com/Damoz1606/mini-storage-backend-golang/pkg/domainevent"
	"github.com/google/uuid"
)

var (
	errEmptyName = apperror.BadRequest("Name is required")
)

type Disk struct {
	domainevent.Recorder
	id       uuid.UUID
	name     string
	createAt int64
	deleteAt *int64
}

func New(name string) (*Disk, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errEmptyName
	}

	return &Disk{
		id:       uuid.New(),
		name:     name,
		createAt: time.Now().UnixMilli(),
	}, nil
}

func (d Disk) ID() uuid.UUID {
	return d.id
}

func (d Disk) Name() string {
	return d.name
}

func (d Disk) CreateAt() int64 {
	return d.createAt
}

func (d Disk) DeleteAt() *int64 {
	return d.deleteAt
}

func (d *Disk) Delete() {
	now := time.Now().UnixMilli()
	d.deleteAt = &now

	eventPayload := EventPayloadDiskDeleted{Value: d.id}
	d.Record(domainevent.Event{
		Type:    EventDiskDeleted,
		Payload: eventPayload,
	})
}
