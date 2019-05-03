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

//Object set FowardMessage method
type ToForward struct {
	ChatID     int
	FromChatID int
	MessageID  int
	Text       string
}

//Object set to reply_to_message_id in SendMessage method
type ToReply struct {
	ChatID    int
	MessageID int
	Text      string
}

//Object set to SendMessage
type ToSend struct {
	ChatID int
	Text   string
}
