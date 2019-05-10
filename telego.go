package telego

import (
	"encoding/json"
	//"fmt"
//	"io/ioutil"
//	"net/http"
	//"net/url"
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
	Response, _ = Config.makeAPIrequest(Concatenado)
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

//Make a "getUpdates" requests and handle the json
func (Config *Load) GetAllUpdates() (bool, SetMessage) {
	Config.Metodo = "getUpdates"
	_, _ = Config.makeAPIrequest(Concatenado)
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
