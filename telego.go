package telego

import (
	"encoding/json"
	"net/url"
)

//Set the bot token
func SetNewBot(token string) API {
	Config.Token = token
	return Config
}

//Send a message
func (Config *API) Send_Message(chat_id int, text string) int {
	Config.Method = "SendMessage"
	MessageBind.sending = ToSend{
		chat_id,
		text,
	}
	_, stat := Config.makeAPIrequest(MessageBind.sending.makeParam())
	return stat
}

//Reply to a message
func (Config *API) Reply_To(message SetMessage, chat_id int, text string) int {
	Config.Method = "SendMessage"
	MessageBind.replying = ToReply{
		chat_id,
		message.MessageID,
		text,
	}
	_, stat := Config.makeAPIrequest(MessageBind.replying.makeParam())
	return stat
}

//Get the bot info
func (Config *API) GetMe() MeBot {
	Config.Method = "getMe"
	Response, _ = Config.makeAPIrequest(url.Values{})
	json.Unmarshal([]byte(Response), &BotInfo)
	return BotInfo
}

//Foward a message
func (Config *API) ForwardMessage(message SetMessage, chat_id int) int {
	Config.Method = "forwardMessage"
	MessageBind.forwarding = ToForward{
		chat_id,
		message.Chat.ID,
		message.MessageID,
		message.Text}
	_, stat := Config.makeAPIrequest(MessageBind.forwarding.makeParam())
	return stat
}

func (Config *API) SendPhoto(chat_id int, photo string) int {
	Config.Method = "sendPhoto"
	MessageBind.sendingphoto = ToSendPhoto{
		chat_id,
		photo,
	}
	_, stat := Config.makeAPIrequest(MessageBind.sendingphoto.makeParam())
	return stat
}

func (Config *API) SendAudio(chat_id int, audio string) int {
	Config.Method = "sendAudio"
	MessageBind.sendingaudio = ToSendAudio{
		chat_id,
		audio,
	}
	_, stat := Config.makeAPIrequest(MessageBind.sendingaudio.makeParam())
	return stat
}

func (Config *API) SendDocument(chat_id int, document string) int {
	Config.Method = "sendDocument"
	MessageBind.sendingdocument = ToSendDocument{
		chat_id,
		document,
	}
	_, stat := Config.makeAPIrequest(MessageBind.sendingdocument.makeParam())
	return stat
}

func (Config *API) SendVideo(chat_id int, video string) int {
	Config.Method = "sendVideo"
	MessageBind.sendingvideo = ToSendVideo{
		chat_id,
		video,
	}
	_, stat := Config.makeAPIrequest(MessageBind.sendingvideo.makeParam())
	return stat
}

func (Config *API) SendAnimation(chat_id int, animation string) int {
	Config.Method = "sendAnimation"
	MessageBind.sendinganimation = ToSendAnimation{
		chat_id,
		animation,
	}
	_, stat := Config.makeAPIrequest(MessageBind.sendinganimation.makeParam())
	return stat
}

func (Config *API) SendVoice(chat_id int, voice string) int {
	Config.Method = "sendVoice"
	MessageBind.sendingvoice = ToSendVoice{
		chat_id,
		voice,
	}
	_, stat := Config.makeAPIrequest(MessageBind.sendingvoice.makeParam())
	return stat
}

func (Config *API) SendLocation(chat_id int, latitude float64, longitude float64) int {
	Config.Method = "sendLocation"
	MessageBind.sendinglocation = ToSendLocation{
		chat_id,
		latitude,
		longitude,
	}
	_, stat := Config.makeAPIrequest(MessageBind.sendinglocation.makeParam())
	return stat
}

//Make a "getUpdates" requests and handle the json
func (Config *API) GetAllUpdates() (bool, SetMessage) {
	Config.Method = "getUpdates"
	_, _ = Config.makeAPIrequest(url.Values{})
	for _, msg := range Bind.Result {
		Nova = msg.Message
	}
	if len(Nova.Text) > 0 && Nova.MessageID != Config.Current {
		Config.Updated = Nova
		Config.Current = Nova.MessageID
		NewMsg = true
	} else {
		Config.Updated.Text = ""
		NewMsg = false

	}
	return NewMsg, Nova

}

func (Config *API) GetMsgUpdates() MsgUpdater {
	Config.Method = "getUpdates"
	_, _ = Config.makeAPIrequest(url.Values{})
	return Bind

}

//Webhook

func (Config *API) SetWebhookWithCert(url string, certificate string) int {
	Config.Method = "setWebhook"
	MessageBind.settingwebhookwithcert = ToSetWebhookWithCert{
		url,
		certificate,
	}
	_, stat := Config.makeAPIrequest(MessageBind.settingwebhookwithcert.makeParam())
	return stat
}

func (Config *API) SetWebhook(url string) int {
	Config.Method = "setWebhook"
	MessageBind.settingwebhook = ToSetWebhook{
		url,
	}
	_, stat := Config.makeAPIrequest(MessageBind.settingwebhook.makeParam())
	return stat
}

func (Config *API) DeleteWebhook() int {
	Config.Method = "DeleteWebhook"
	_, stat := Config.makeAPIrequest(url.Values{})
	return stat
}

func (Config *API) GetWebhookinfo() (string, int) {
	Config.Method = "getWebhookinfo"
	res, stat := Config.makeAPIrequest(url.Values{})
	return res, stat
}
