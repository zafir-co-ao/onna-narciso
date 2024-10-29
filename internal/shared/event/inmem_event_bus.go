package event

import (
	"sync"
)

type inmemEventBus struct {
	subscribers map[string][]Handler

	mutex sync.RWMutex
	queue []Event
}

func NewInmemEventBus() Bus {
	return &inmemEventBus{
		subscribers: make(map[string][]Handler),
		queue:       make([]Event, 0),
	}
}

func (p *inmemEventBus) PublishAll(events ...Event) {
	for _, e := range events {
		p.Publish(e)
	}
}

func (p *inmemEventBus) Publish(e Event) {

	p.mutex.Lock()
	p.queue = append(p.queue, e)
	p.mutex.Unlock()

	if len(p.queue) == 1 {
		p.start()
	}
}
func (p *inmemEventBus) start() {
	for len(p.queue) > 0 {
		p.mutex.Lock()
		e := p.queue[0]
		p.queue = p.queue[1:]
		p.mutex.Unlock()

		p.handle(e)
	}
}

func (p *inmemEventBus) handle(e Event) {
	wg := sync.WaitGroup{}

	f := func(s Handler, e Event) {
		defer wg.Done()
		s.Handle(e)
	}

	for _, s := range p.subscribers[e.Name()] {
		wg.Add(1)
		go f(s, e)
	}

	if e.Name() != WildcardEventName {
		for _, s := range p.subscribers[WildcardEventName] {
			wg.Add(1)
			go f(s, e)
		}
	}

	wg.Wait()
}

func (p *inmemEventBus) Subscribe(theType string, subscriber Handler) {
	subcribers, exist := p.subscribers[theType]

	if !exist {
		p.subscribers[theType] = []Handler{subscriber}
		return
	}

	p.subscribers[theType] = append(subcribers, subscriber)
}

func (p *inmemEventBus) SubscribeFunc(ename string, f func(Event)) {
	p.Subscribe(ename, WrapHandlerFunc(f))
}
