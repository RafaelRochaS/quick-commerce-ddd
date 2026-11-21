package order

type Messenger interface {
	SendMessage(eventName, msg string) error
}
