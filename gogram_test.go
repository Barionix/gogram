package gogram

import (
	"testing"
)

var Bot API = SetNewBot("811530665:AAG5X41LQS5tJbwBaTmbs8tVXgeWYGhqYrM")

func TryConf(t *testing.T) {
	Bot := SetNewBot("811530665:AAG5X41LQS5tJbwBaTmbs8tVXgeWYGhqYrM")
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

func TestSendPhoto(t *testing.T) {
	try := Bot.SendPhoto(476036515, "AgADAQADF6gxG0nGsEbYLmuxSm6YlGwdFDAABI1ha-i8vPBpih4CAAEC")
	if try != 200 {
		t.Fatalf("Something went wrong in SendPhoto.\nStatus Code:[%d]", try)
	}
}

func TestSendVideoNote(t *testing.T) {
	payload := ToSendVideoNote{
		ChatID:    476036515,
		VideoNote: "BAACAgEAAxkBAANAXx-GBPU7S-vvVS32vNz4dmQ1V8IAAtUAA4ukAUXuhv0okZi9phoE",
	}
	try := Bot.SendVideoNote(payload)
	if try != 200 {
		t.Fatalf("Something went wrong in SendVideoNote.\nStatus Code:[%d]", try)
	}
}

func TestKickChatMember(t *testing.T) {
	payload := ToKickChatMember{
		ChatID: -1001288797747,
		UserID: 1122169250,
	}
	try := Bot.KickChatMember(payload)
	if try != 200 {
		t.Fatalf("Something went wrong in KicChatMember.\nStatus Code:[%d]", try)
	}
}

func TestRestrictChatMembers(t *testing.T) {
	payload := ToRestrictChatMembers{
		ChatID: -1001288797747,
		UserID: 1122169250,
		Permissions: ChatPermissions{
			CanSendMessages:       true,
			CanSendMediaMessage:   false,
			CanSendPolls:          false,
			CanSendOtherMessages:  false,
			CanAddWebPagePreviews: false,
			CanChangeInfo:         true,
			CanInviteUsers:        false,
			CanPinMessage:         false,
		},
	}
	try := Bot.RestrictChatMember(payload)
	if try != 200 {
		t.Fatalf("Something went wrong in RestrictChatMember.\nStatus Code:[%d]", try)
	}
}

func TestSetChatTitle(t *testing.T) {
	payload := ToSetChatTitle{
		ChatID: -1001288797747,
		Title:  "GOGRAM DEVELOPMENT PROCESS",
	}
	try := Bot.SetChatTitle(payload)
	if try != 200 {
		t.Fatalf("Something went wrong in SetChatTitle.\nStatus Code:[%d]", try)
	}
}

func TestSetChatDescription(t *testing.T) {
	payload := ToSetChatDescription{
		ChatID:      -1001288797747,
		Description: "GOGRAM DEVELOPMENT PROCESS - DESC",
	}
	try := Bot.SetChatDescription(payload)
	if try != 200 {
		t.Fatalf("Something went wrong in SetChatDescription.\nStatus Code:[%d]", try)
	}
}

func TestPinChatMessage(t *testing.T) {
	payload := ToPinChatMessage{
		ChatID:    -1001288797747,
		MessageID: 4,
	}
	try := Bot.PinChatMessage(payload)
	if try != 200 {
		t.Fatalf("Something went wrong in PinChatMessage.\nStatus Code:[%d]", try)
	}
}
func TestSetChatAdministratorCustomTitle(t *testing.T) {
	payload := ToSetChatAdministratorCustomTitle{
		ChatID:      -1001288797747,
		UserID:      811530665,
		CustomTitle: "Freaking Boss",
	}
	try := Bot.SetChatAdministratorCustomTitle(payload)
	if try != 200 {
		t.Fatalf("Something went wrong in SetChatAdministratorCustomTitle.\nStatus Code:[%d]", try)
	}
}

func TestSetChatPermissions(t *testing.T) {
	payload := ToSetChatPermissions{
		ChatID: -1001288797747,
		Permissions: ChatPermissions{
			CanSendMessages:       false,
			CanSendMediaMessage:   false,
			CanSendPolls:          false,
			CanSendOtherMessages:  false,
			CanAddWebPagePreviews: false,
			CanChangeInfo:         true,
			CanInviteUsers:        false,
			CanPinMessage:         false,
		},
	}
	try := Bot.SetChatPermissions(payload)
	if try != 200 {
		t.Fatalf("Something went wrong in SetChatPermissions.\nStatus Code:[%d]", try)
	}
}

/*
func TestSendMediaGroup(t *testing.T) {
	payload := ToSendMediaGroup{
		ChatID: 476036515,
		Media: []InputMediaPhoto{
			{
				"photo",
				"https://commons.wikimedia.org/wiki/File:Moon.jpg",
				"Lorem fucking ipsum",
			},
			{
				"photo",
				"https://commons.wikimedia.org/wiki/File:Moon.jpg",
				"Lorem fucking ipsum",
			},
		},
	}
	try := Bot.SendMediaGroup(payload)
	if try != 200 {
		t.Fatalf("Something went wrong in SendMediaGroup.\nStatus Code:[%d]", try)
	}
}*/

func TestSendAudio(t *testing.T) {
	try := Bot.SendAudio(476036515, "CQADAgADmIoDAAEcg5gKx_YQOaoQbjoC")
	if try != 200 {
		t.Fatalf("Something went wrong in SendAudio.\nStatus Code:[%d]", try)
	}
}

func TestSendDocument(t *testing.T) {
	try := Bot.SendDocument(476036515, "BQADAQADTQADScawRi4Q_TGQxqu4Ag")
	if try != 200 {
		t.Fatalf("Something went wrong in sendDocument.\nStatus Code:[%d]", try)
	}
}

func TestSendVideo(t *testing.T) {
	try := Bot.SendVideo(476036515, "BAADAQADbAADMIGwRlHrrddJsZ3KAg")
	if try != 200 {
		t.Fatalf("Something went wrong in SendVideo.\nStatus Code:[%d]", try)
	}
}

func TestSendAnimation(t *testing.T) {
	try := Bot.SendAnimation(476036515, "CgADBAADDQQAAr8ZZAeC0k40J-kcJQI")
	if try != 200 {
		t.Fatalf("Something went wrong in SendAnimation.\nStatus Code:[%d]", try)
	}
}

func TestSendVoice(t *testing.T) {
	try := Bot.SendVoice(476036515, "AwADAQADbgADMIGwRpptt4iJrzjsAg")
	if try != 200 {
		t.Fatalf("Something went wrong in SendVoice.\nStatus Code:[%d]", try)
	}
}

func TestSendLocation(t *testing.T) {
	try := Bot.SendLocation(476036515, -37.824075, -37.824075)
	if try != 200 {
		t.Fatalf("Something went wrong in SendLocation.\nStatus Code:[%d]", try)
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
	if !me.Ok {
		t.Fatalf("Eror in getMe!")
	}
}
