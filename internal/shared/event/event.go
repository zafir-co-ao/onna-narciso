package event

import (
	"errors"
	"time"

	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

const (
	HeaderEventID       = "EventID"
	HeaderEventName     = "EventName"
	HeaderIssuedAt      = "IssuedAt"
	HeaderAggregateType = "AggregateType"
	HeaderAggregateID   = "AggregateID"
	HeaderCommandID     = "CommandID"
)

var ErrEventNotFound = errors.New("event not found")

type EventOpt func(Event) Event

func WithPayload(b interface{}) EventOpt {
	return func(e Event) Event {
		e.payload = b
		return e
	}
}

type HeaderEntry [2]string

func WithHeader(k string, v string) EventOpt {
	return func(e Event) Event {
		e.headers[k] = v
		return e
	}
}

func WithHeaders(h ...HeaderEntry) EventOpt {
	return func(evt Event) Event {
		for _, e := range h {
			evt.headers[e[0]] = e[1]
		}
		return evt
	}
}

func WithOptions(opts ...EventOpt) EventOpt {
	return func(evt Event) Event {
		for _, opt := range opts {
			evt = opt(evt)
		}
		return evt
	}
}

type headers map[string]string

type Event struct {
	headers headers
	payload interface{}
}

func New(name string, opts ...EventOpt) Event {

	id := id.MustRandom()

	myHeaders := make(map[string]string)

	myHeaders[HeaderEventName] = name

	myHeaders[HeaderEventID] = id.String()
	myHeaders[HeaderIssuedAt] = time.Now().Format("2006-01-02T15:04:05Z")

	e := Event{headers: myHeaders}

	for _, opt := range opts {
		e = opt(e)
	}

	return e
}

func (e Event) Decorate(opts ...EventOpt) Event {

	headers := make([]HeaderEntry, 0)

	return New(
		e.Name(),
		WithHeaders(headers...),
		WithPayload(e.Payload()),
		WithOptions(opts...),
	)
}

func (e Event) ID() string {
	return e.headers[HeaderEventID]
}

func (e Event) Name() string {
	return e.headers[HeaderEventName]
}

func (e Event) IssuedAt() string {
	t := e.headers[HeaderIssuedAt]
	return t
}

func (e Event) Header(key string) string {
	return e.headers[key]
}

func (e Event) Payload() interface{} {
	return e.payload
}
