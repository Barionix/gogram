package telego

type SetValues struct {
	forwarding   ToForward
	replying     ToReply
	sending      ToSend
	sendingphoto ToSendPhoto
	sendingaudio ToSendAudio
	sendingdocument ToSendDocument
	sendingvideo ToSendVideo
	sendinganimation ToSendAnimation
	sendingvoice ToSendVoice
	sendinglocation ToSendLocation
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

type ToSendPhoto struct {
	ChatID int
	PhotoID  string
}

type ToSendAudio struct {
	ChatID int
	AudioID  string
}

type ToSendDocument struct {
	ChatID int
	DocumentID  string
}

type ToSendVideo struct {
	ChatID int
	VideoID  string
}

type ToSendAnimation struct {
	ChatID int
	AnimationID  string
}

type ToSendVoice struct {
	ChatID int
	VoiceID  string
}

type ToSendLocation struct {
	ChatID int
	Latitude  float64
	Longitude float64
}
