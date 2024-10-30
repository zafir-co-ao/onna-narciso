package event

const WildcardEventName = "*"

type Bus interface {
	Publish(event Event)
	PublishAll(events ...Event)
	Subscribe(eventName string, subscriber Handler)
	SubscribeFunc(eventName string, f func(Event))
}

type Handler interface {
	Handle(event Event)
}

type HandlerFunc func(Event)

func (f HandlerFunc) Handle(event Event) {
	f(event)
}

func WrapHandlerFunc(f HandlerFunc) Handler {
	return HandlerFunc(f)
}

type Middleware func(Handler) Handler

func Chain(s Handler, m ...Middleware) Handler {
	for i := len(m) - 1; i >= 0; i-- {
		s = m[i](s)
	}
	return s
}
