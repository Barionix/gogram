package telego

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//Set the bot token
func Conf(token string) Load {
	Config.Token = token
	return Config
}

//Make the requests
func (Config *Load) Req(metodo string, concatenado url.Values) (string, int) {
	API_ADRESS = fmt.Sprintf("https://api.telegram.org/bot%s/%s", Config.Token, metodo)
	resp, err := http.PostForm(API_ADRESS, concatenado)
	defer resp.Body.Close()
	body, er := ioutil.ReadAll(resp.Body)
	handle_erro(err)
	handle_erro(er)
	return string(body), resp.StatusCode
}

//Send a message
func (Config *Load) Send_Message(chat_id int, text string) int {
	MessageBind.sending = ToSend{
		chat_id,
		text,
	}
	_, stat := Config.Req("SendMessage", MessageBind.sending.makeParam())
	return stat
}

//Reply to a message
func (Config *Load) Reply_To(message SetMessage, chat_id int, text string) int {
	MessageBind.replying = ToReply{
		chat_id,
		message.MessageID,
		text,
	}
	_, stat := Config.Req("SendMessage", MessageBind.replying.makeParam())
	return stat
}

//Get the bot info
func (Config *Load) GetMe() MeBot {
	Config.Metodo = "getMe"
	Response, _ = Config.Req("getMe", Concatenado)
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
	_, stat := Config.Req("forwardMessage", MessageBind.forwarding.makeParam())
	return stat
}

//Make a "getUpdates" requests and handle the json
func (Config *Load) GetAllUpdates() (bool, SetMessage) {
	Config.Metodo = "getUpdates"
	Response, _ = Config.Req("getUpdates", Concatenado)
	json.Unmarshal([]byte(Response), &Bind)
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
