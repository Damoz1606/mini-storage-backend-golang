package disk

import (
	"github.com/Damoz1606/mini-storage-backend-golang/pkg/domainevent"
	"github.com/google/uuid"
)

const (
	EventDiskDeleted domainevent.EventType = "disk.deleted"
)

type EventPayloadDiskDeleted struct {
	Value uuid.UUID
}
