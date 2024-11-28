package notifications

type Notifier interface {
	 Notify(n, msg string) error
}
