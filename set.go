package telego

type SetValues struct {
	forwarding ToForward
	replying   ToReply
	sending    ToSend
}

var MessageBind SetValues

type Action interface {
	Req()
}

type Parse interface {
	makeParam()
}
type ToForward struct {
	ChatID     int
	FromChatID int
	MessageID  int
	Text       string
}

type ToReply struct {
	ChatID    int
	MessageID int
	Text      string
}

type ToSend struct {
	ChatID int
	Text   string
}
