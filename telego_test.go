package telego

import (
	"testing"
	"fmt"
)

var Bot Load = Conf("811530665:AAG5X41LQS5tJbwBaTmbs8tVXgeWYGhqYrM")

func TryConf(t *testing.T) {
	Bot := Conf("811530665:AAG5X41LQS5tJbwBaTmbs8tVXgeWYGhqYrM")
	if Bot.Token == "" {
		t.Fatalf("TOKEN NULO")
	}
}

func TestSendMessage(t *testing.T) {
	try := Bot.Send_Message(476036515, "SendMessage funcionando!")
	if try != 200 {
		t.Fatalf("Something went wrong in SendMessage.\nStatus Code:[%d]", try)
	}
}
func TestReply_TO(t *testing.T) {
	Config.Updated.MessageID = 1
	try := Bot.Reply_To(Config.Updated, 476036515, "ReplyTo funcionando")
	if try != 200 {
		t.Fatalf("Something went wrong in SendMessage.\nStatus Code:[%d]", try)
	}
}

func TestFoward(t *testing.T) {
	Config.Updated.Chat.ID = 476036515
	Config.Updated.MessageID = 2
	try := Bot.ForwardMessage(Config.Updated, 476036515)
	if try != 200 {
		t.Fatalf("Something went wrong in FowardMessage.\nStatus Code:[%d]", try)
	}
}

func TestGetMe(t *testing.T) {
	me := Bot.GetMe()
	fmt.Println(me.Ok)
	if !me.Ok {
		t.Fatalf("Eror in getMe!")
	}
}
