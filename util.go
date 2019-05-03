package telego

import (
	"fmt"
	//"strings"
	"net/url"
	"strconv"
)

func handle_erro(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

/*
func is_command(message SetMessage) bool {
	return message.Entity.Type == "bot_command"
}*/

func (forward *ToForward) makeParam() url.Values {
	param := url.Values{}
	param.Set("text", forward.Text)
	param.Set("chat_id", strconv.Itoa(forward.ChatID))
	param.Set("from_chat_id", strconv.Itoa(forward.FromChatID))
	param.Set("message_id", strconv.Itoa(forward.MessageID))

	return param
}

func (reply *ToReply) makeParam() url.Values {
	param := url.Values{}
	param.Set("text", reply.Text)
	param.Set("reply_to_message_id", strconv.Itoa(reply.MessageID))
	param.Set("chat_id", strconv.Itoa(reply.ChatID))

	return param
}

func (send *ToSend) makeParam() url.Values {
	param := url.Values{}
	param.Set("text", send.Text)
	param.Set("chat_id", strconv.Itoa(send.ChatID))
	fmt.Println(send.Text)
	return param
}
