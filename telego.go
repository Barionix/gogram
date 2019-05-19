package telego

import (
	"encoding/json"
	"net/url"
)

//Set the bot token
func Conf(token string) Load {
	Config.Token = token
	return Config
}

//Send a message
func (Config *Load) Send_Message(chat_id int, text string) int {
	Config.Metodo = "SendMessage"
	MessageBind.sending = ToSend{
		chat_id,
		text,
	}
	_, stat := Config.makeAPIrequest(MessageBind.sending.makeParam())
	return stat
}

//Reply to a message
func (Config *Load) Reply_To(message SetMessage, chat_id int, text string) int {
	Config.Metodo = "SendMessage"
	MessageBind.replying = ToReply{
		chat_id,
		message.MessageID,
		text,
	}
	_, stat := Config.makeAPIrequest(MessageBind.replying.makeParam())
	return stat
}

//Get the bot info
func (Config *Load) GetMe() MeBot {
	Config.Metodo = "getMe"
	Response, _ = Config.makeAPIrequest(url.Values{})
	json.Unmarshal([]byte(Response), &BindG)
	return BindG
}

//Foward a message
func (Config *Load) ForwardMessage(message SetMessage, chat_id int) int {
	Config.Metodo = "forwardMessage"
	MessageBind.forwarding = ToForward{
		chat_id,
		message.Chat.ID,
		message.MessageID,
		message.Text}
	_, stat := Config.makeAPIrequest(MessageBind.forwarding.makeParam())
	return stat
}

func (Config *Load) SendPhoto(chat_id int, photo string) int {
	Config.Metodo = "sendPhoto"
	MessageBind.sendingphoto = ToSendPhoto{
		chat_id,
		photo,
	}
	_, stat := Config.makeAPIrequest(MessageBind.sendingphoto.makeParam())
	return stat
}

func (Config *Load) SendAudio(chat_id int, audio string) int {
	Config.Metodo = "sendAudio"
	MessageBind.sendingaudio = ToSendAudio{
		chat_id,
		audio,
	}
	_, stat := Config.makeAPIrequest(MessageBind.sendingaudio.makeParam())
	return stat
}

func (Config *Load) SendDocument(chat_id int, document string) int {
	Config.Metodo = "sendDocument"
	MessageBind.sendingdocument = ToSendDocument{
		chat_id,
		document,
	}
	_, stat := Config.makeAPIrequest(MessageBind.sendingdocument.makeParam())
	return stat
}

func (Config *Load) SendVideo(chat_id int, video string) int {
	Config.Metodo = "sendVideo"
	MessageBind.sendingvideo = ToSendVideo{
		chat_id,
		video,
	}
	_, stat := Config.makeAPIrequest(MessageBind.sendingvideo.makeParam())
	return stat
}

func (Config *Load) SendAnimation(chat_id int, animation string) int {
	Config.Metodo = "sendAnimation"
	MessageBind.sendinganimation = ToSendAnimation{
		chat_id,
		animation,
	}
	_, stat := Config.makeAPIrequest(MessageBind.sendinganimation.makeParam())
	return stat
}

func (Config *Load) SendVoice(chat_id int, voice string) int {
	Config.Metodo = "sendVoice"
	MessageBind.sendingvoice = ToSendVoice{
		chat_id,
		voice,
	}
	_, stat := Config.makeAPIrequest(MessageBind.sendingvoice.makeParam())
	return stat
}

func (Config *Load) SendLocation(chat_id int, latitude float64, longitude float64) int {
	Config.Metodo = "sendLocation"
	MessageBind.sendinglocation = ToSendLocation{
		chat_id,
		latitude,
		longitude,
	}
	_, stat := Config.makeAPIrequest(MessageBind.sendinglocation.makeParam())
	return stat
}
//Make a "getUpdates" requests and handle the json
func (Config *Load) GetAllUpdates() (bool, SetMessage) {
	Config.Metodo = "getUpdates"
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

//Webhook

func (Config *Load) SetWebhookWithCert(url string, certificate string) int{
	Config.Metodo = "setWebhook"
	MessageBind.settingwebhookwithcert = ToSetWebhookWithCert{
		url,
		certificate,
	}
	_, stat := Config.makeAPIrequest(MessageBind.settingwebhookwithcert.makeParam())
	return stat
}


func (Config *Load) SetWebhook(url string) int{
	Config.Metodo = "setWebhook"
	MessageBind.settingwebhook = ToSetWebhook{
		url,
	}
	_, stat := Config.makeAPIrequest(MessageBind.settingwebhook.makeParam())
	return stat
}

func (Config *Load) DeleteWebhook() int{
	Config.Metodo = "DeleteWebhook"
	_, stat := Config.makeAPIrequest(url.Values{})
	return stat
}
	
 
func (Config *Load) GetWebhookinfo() (string, int){
	Config.Metodo = "getWebhookinfo"
	res, stat := Config.makeAPIrequest(url.Values{})
	return res, stat
}
	