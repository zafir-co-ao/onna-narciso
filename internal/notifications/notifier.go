package notifications

type Notifier interface {
	Notify(n, msg string) error
}

type NotifierFunc func(n, msg string) error

func (f NotifierFunc) Notify(n, msg string) error {
	return f(n, msg)
}
