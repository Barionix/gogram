package telego

import (
	"fmt"
	"net/url"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

//Handle usual errors
func handle_erro(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//Set the the url to make a forwardMessage request
func (forward *ToForward) makeParam() url.Values {
	param := url.Values{}
	param.Set("text", forward.Text)
	param.Set("chat_id", strconv.Itoa(forward.ChatID))
	param.Set("from_chat_id", strconv.Itoa(forward.FromChatID))
	param.Set("message_id", strconv.Itoa(forward.MessageID))

	return param
}

//Set the the url to make a reply_to_message_id request
func (reply *ToReply) makeParam() url.Values {
	param := url.Values{}
	param.Set("text", reply.Text)
	param.Set("reply_to_message_id", strconv.Itoa(reply.MessageID))
	param.Set("chat_id", strconv.Itoa(reply.ChatID))

	return param
}

//Set the the url to make a SendMessage request
func (send *ToSend) makeParam() url.Values {
	param := url.Values{}
	param.Set("text", send.Text)
	param.Set("chat_id", strconv.Itoa(send.ChatID))
	return param
}



//Make the requests
func (Config *Load) makeAPIrequest(concatenado url.Values) (string, int) {
	API_ADRESS = fmt.Sprintf("https://api.telegram.org/bot%s/%s", Config.Token, Config.Metodo)
	resp, err := http.PostForm(API_ADRESS, concatenado)
	defer resp.Body.Close()
	body, er := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &Bind)
	handle_erro(err)
	handle_erro(er)
	return string(body), resp.StatusCode
}
