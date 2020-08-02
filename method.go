package gogram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func handle_erro(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//Make a resquest to the API
func (Config *API) makeAPIrequest(payload url.Values) (string, int) {
	API_ADDRESS = fmt.Sprintf("https://api.telegram.org/bot%s/%s", Config.Token, Config.Method)
	resp, err := http.PostForm(API_ADDRESS, payload)
	defer resp.Body.Close()
	body, er := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &Bind)
	handle_erro(err)
	handle_erro(er)
	return string(body), resp.StatusCode
}

//Make an API request to the forwardMessage method
func (forward *ToForward) makeParam() url.Values {
	param := url.Values{}
	param.Set("text", forward.Text)
	param.Set("chat_id", strconv.Itoa(forward.ChatID))
	param.Set("from_chat_id", strconv.Itoa(forward.FromChatID))
	param.Set("message_id", strconv.Itoa(forward.MessageID))

	return param
}

//Make an API request to the reply_to_message_id method
func (reply *ToReply) makeParam() url.Values {
	param := url.Values{}
	param.Set("text", reply.Text)
	param.Set("reply_to_message_id", strconv.Itoa(reply.MessageID))
	param.Set("chat_id", strconv.Itoa(reply.ChatID))

	return param
}

//Make an API request to the SendMessage method
func (send *ToSend) makeParam() url.Values {
	param := url.Values{}
	param.Set("text", send.Text)
	param.Set("chat_id", strconv.Itoa(send.ChatID))
	return param
}

//Make an API request to the sendPhoto method
func (photo *ToSendPhoto) makeParam() url.Values {
	param := url.Values{}
	param.Set("chat_id", strconv.Itoa(photo.ChatID))
	param.Set("photo", photo.PhotoID)
	return param
}

//Make an API request to the sendAudio method
func (audio *ToSendAudio) makeParam() url.Values {
	param := url.Values{}
	param.Set("chat_id", strconv.Itoa(audio.ChatID))
	param.Set("audio", audio.AudioID)
	return param
}

//Make an API request to the sendDocument method
func (document *ToSendDocument) makeParam() url.Values {
	param := url.Values{}
	param.Set("chat_id", strconv.Itoa(document.ChatID))
	param.Set("document", document.DocumentID)
	return param
}

func (video *ToSendVideoNote) makeParam() url.Values {
	param := url.Values{}
	param.Set("chat_id", strconv.Itoa(video.ChatID))
	param.Set("video_note", video.VideoNote)
	param.Set("duration", strconv.Itoa(video.Duration))
	param.Set("length", strconv.Itoa(video.Length))
	param.Set("thumb", video.Thumb)
	param.Set("disable_notification", strconv.FormatBool(video.DisableNotification))
	param.Set("reply_to_message_id", strconv.Itoa(video.ReplyToMessageID))

	return param
}

//Make an API request to the sendVideoNote method
func (video *ToSendVideo) makeParam() url.Values {
	param := url.Values{}
	param.Set("chat_id", strconv.Itoa(video.ChatID))
	param.Set("video", video.VideoID)
	return param
}

//Make an API request to the sendAnimation method
func (animation *ToSendAnimation) makeParam() url.Values {
	param := url.Values{}
	param.Set("chat_id", strconv.Itoa(animation.ChatID))
	param.Set("animation", animation.AnimationID)
	return param
}

//Make an API request to the sendVoice method
func (voice *ToSendVoice) makeParam() url.Values {
	param := url.Values{}
	param.Set("chat_id", strconv.Itoa(voice.ChatID))
	param.Set("voice", voice.VoiceID)
	return param
}

//Make an API request to the sendLocation method
func (location *ToSendLocation) makeParam() url.Values {
	param := url.Values{}
	param.Set("chat_id", strconv.Itoa(location.ChatID))
	param.Set("latitude", fmt.Sprintf("%f", location.Latitude))
	param.Set("longitude", fmt.Sprintf("%f", location.Longitude))
	return param
}

//Make an API request to the kickChatMember method
func (chat *ToKickChatMember) makeParam() url.Values {
	param := url.Values{}
	param.Set("chat_id", strconv.Itoa(chat.ChatID))
	param.Set("user_id", strconv.Itoa(chat.UserID))
	param.Set("until_date", strconv.Itoa(chat.UntilDate))
	return param
}

//Make an API request to the restrictChatMembet method
func (chat *ToRestrictChatMembers) makeParam() url.Values {
	param := url.Values{}
	Permissions, err := json.Marshal(chat.Permissions)
	handle_erro(err)
	param.Set("chat_id", strconv.Itoa(chat.ChatID))
	param.Set("user_id", strconv.Itoa(chat.UserID))
	param.Set("permissions", BytesToString(Permissions))
	param.Set("until_date", strconv.Itoa(chat.UntilDate))
	return param
}

//Make an API request to the setChatTitle method
func (chat *ToSetChatTitle) makeParam() url.Values {
	param := url.Values{}
	param.Set("chat_id", strconv.Itoa(chat.ChatID))
	param.Set("title", chat.Title)
	return param
}

//Make an API request to the setChatDescriptions method
func (chat *ToSetChatDescription) makeParam() url.Values {
	param := url.Values{}
	param.Set("chat_id", strconv.Itoa(chat.ChatID))
	param.Set("description", chat.Description)
	return param
}

//Make an API request to the setChatPermissions method
func (chat *ToSetChatPermissions) makeParam() url.Values {
	param := url.Values{}
	Permissions, err := json.Marshal(chat.Permissions)
	handle_erro(err)
	param.Set("chat_id", strconv.Itoa(chat.ChatID))
	param.Set("permissions", BytesToString(Permissions))
	return param
}

//WIP
func (webhook *ToSetWebhookWithCert) makeParam() url.Values {
	param := url.Values{}
	param.Set("url", webhook.Url)
	param.Set("certificate", webhook.Certificate)
	return param
}

func (webhook *ToSetWebhook) makeParam() url.Values {
	param := url.Values{}
	param.Set("url", webhook.Url)
	return param
}
