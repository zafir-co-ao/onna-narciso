package notifications

type Contact struct {
	Name   string
	Email  string
	Mobile string
}

type Message struct {
	Subject string
	Body    string
}

type Notifier interface {
	Notify(c Contact, msg Message) error
}

type NotifierFunc func(c Contact, msg Message) error

func (f NotifierFunc) Notify(c Contact, msg Message) error {
	return f(c, msg)
}
