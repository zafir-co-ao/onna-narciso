package stubs

import "github.com/zafir-co-ao/onna-narciso/internal/notifications"

func NewNotificationsACL() *notificationsStub {
	return &notificationsStub{}
}

type notificationsStub struct {
	Contact notifications.Contact
	Message notifications.Message
}

func (n *notificationsStub) Notify(c notifications.Contact, msg notifications.Message) error {
	n.Contact = c
	n.Message = msg
	return nil
}
