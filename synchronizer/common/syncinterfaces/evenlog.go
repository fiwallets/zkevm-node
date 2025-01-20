package syncinterfaces

import (
	"context"

	"github.com/fiwallets/zkevm-node/event"
)

// EventLogInterface write an event to the event log database
type EventLogInterface interface {
	LogEvent(ctx context.Context, event *event.Event) error
}
