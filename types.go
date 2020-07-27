package gogram

import (
	"net/http"
	//	"net/url"
)

var (
	API_ADRESS string // Telegram API URL
	Response   string
	NewMsg     bool
	Bind       MsgUpdater
	BotInfo    MeBot
	Config     API
	Nova       SetMessage
)

type SetUser struct {
	ID            int
	Is_Bot        bool
	FirstName     string
	Last_name     string
	Username      string
	language_code string
}

type API struct {
	Token    string
	Method   string
	Current  int
	Client   *http.Client
	Updated  SetMessage
	Response string
}

type SetEntity struct {
	Offset int     `json:"offset"`
	Length int     `json:"length"`
	Type   string  `json:"type"`
	User   SetUser `json:"user"`
	Url    string  `json:"url"`
}

type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"widht"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"`
}

type SetAudio struct {
	FileID    string    `json:"file_id"`
	Duration  int       `json:"duration"`
	Perfomer  string    `json:"performer"`
	Title     string    `json:"title"`
	Mime_Type string    `json:"mime_type"`
	FileSize  int       `json:"file_size"`
	Thumb     PhotoSize `json:"thumb"`
}

type SetMaskPosition struct {
	Point  string  `json:"point"`
	Xshift float64 `json:"x_shift"`
	Yshift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}
type SetSticker struct {
	FileID       string          `json:"file_id"`
	Width        int             `json:"widht"`
	Height       int             `json:"height"`
	Thumb        PhotoSize       `json:"thumb"`
	emoji        string          `json:"emoji"`
	SetName      string          `json:"set_name"`
	MaskPosition SetMaskPosition `json:"mask_position"`
	FileSize     int             `json:"file_size"`
}

type SetDocument struct {
	FileID    string    `json:"file_id"`
	Thumb     PhotoSize `json:"thumb"`
	FileName  string    `json:"file_name"`
	Mime_Type string    `json:"mime_type"`
	FileSize  int       `json:"file_size"`
}

type SetVideo struct {
	FileID    string    `json:"file_id"`
	Width     int       `json:"widht"`
	Height    int       `json:"height"`
	Duration  int       `json:"duration"`
	Thumb     PhotoSize `json:"thumb"`
	Mime_Type string    `json:"mime_type"`
	FileSize  int       `json:"file_size"`
}

type SetVideoNote struct {
	FileID   string    `json:"file_id"`
	Length   int       `json:"length"`
	Duration int       `json:"duration"`
	Thumb    PhotoSize `json:"thumb"`
	FileSize int       `json:"file_size"`
}

type SetAnimation struct {
	FileID    string    `json:"file_id"`
	Width     int       `json:"widht"`
	Height    int       `json:"height"`
	Duration  int       `json:"duration"`
	Thumb     PhotoSize `json:"thumb"`
	FileName  string    `json:"file_name"`
	Mime_Type string    `json:"mime_type"`
	FileSize  int       `json:"file_size"`
}

type SetVoice struct {
	FileID    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Mime_Type string `json:"mime_type"`
	FileSize  int    `json:"file_size"`
}

type SetContact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserID      int    `json:"user_id"`
	Vcard       string `json:"vcard"`
}

type SetLocation struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type SetVenue struct {
	Location       SetLocation `json:"location"`
	Title          string      `json:"title"`
	Adress         string      `json:"adress"`
	FoursquareID   string      `json:"foursquare_id"`
	FoursquareType string      `json:"foursquare_type"`
}

type SetPollOption struct {
	PollID     string  `json:"poll_id"`
	User       SetUser `json:"user"`
	OptionsIDS []int   `json:"option_ids"`
}

type SetPollAnswer struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voteCount"`
}

type SetMessage struct {
	MessageID             int          `json:"message_id"`
	From                  SetFrom      `json:"from"`
	Chat                  SetChat      `json:"chat"`
	Entity                []SetEntity  `json:"entities,omitempty"`
	ForwardFrom           SetUser      `json:"forward_from"`
	ForwardFromChat       SetUser      `json:"forward_from_chat"`
	ForwardSignature      string       `json:"forward_signature"`
	ForwardSenderName     string       `json:"forward_sender_name"`
	ForwardDate           string       `json:"forward_date"`
	ReplyTo               *SetMessage  `json:"reply_to_message"`
	EditDate              int          `json:"edit_date"`
	MediaGroupID          string       `json:"media_group_id"`
	AuthorSignature       int          `json:"author_signature"`
	CaptionEntities       []SetEntity  `json:"caption_entities"`
	Audio                 SetAudio     `json:"audio"`
	Document              SetDocument  `json:"document"`
	Animation             SetAnimation `json:"animation"`
	Game                  SetGame      `json:"game"`
	Photo                 []PhotoSize  `json:"photo"`
	Sticker               SetSticker   `json:"sticker"`
	Date                  int          `json:"date"`
	Text                  string       `json:"text"`
	Voice                 SetVoice     `json:"voice"`
	VideoNote             SetVideoNote `json:"video_note"`
	Contact               SetContact   `json:"contact"`
	Location              SetLocation  `json:"location"`
	Venue                 SetVenue     `json:"venue"`
	Poll                  SetPoll      `json:"poll"`
	NewChatMembers        []SetUser    `json:"new_chat_participant"`
	LeftChatMember        SetUser      `json:"left_chat_participant"`
	NewChatTitle          string       `json:"new_chat_title"`
	NewChatPhoto          string       `json:"new_chat_photo"`
	DeleteChatPhoto       bool         `json:"delete_chat_photo"`
	GroupChatCreated      bool         `json:"group_chat_created"`
	SuperGroupChatCreated bool         `json:"supergroup_chat_created"`
	ChannelChatCreated    bool         `json:"channel_chat_created"`
	PinnedMessage         *SetMessage  `json:"pinned_message"`
}

type SetPoll struct {
	ID                  string               `json:"id"`
	Question            string               `json:"question"`
	Option              []SetPollOption      `json:"poll_option"`
	IsClosed            bool                 `json:"is_closed"`
	TotalVoterCounter   int                  `json:"total_voter_count"`
	IsAnonymous         bool                 `json:"is_anonymous"`
	Type                string               `json:"type"`
	AllowMutipleAnswers bool                 `json:"allows_multiple_answers"`
	CorrectOptionID     int                  `json:"correct_option_id"`
	Explanation         string               `json:"explanation"`
	ExplanationEntities []SetMessageEntities `json:"explanation_entities"`
	OpenPeriod          int                  `json:"open_period"`
	CloseDate           int                  `json:"close_date"`
}

type SetGame struct {
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Photo        []PhotoSize  `json:"photo"`
	Text         string       `json:"text"`
	TextEntities []SetEntity  `json:"text_entities"`
	Animation    SetAnimation `json:"animation"`
}
type SetFrom struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}
type UserProfilePhotos struct {
	TotalCount int         `json:"total_coutn"`
	Photos     []PhotoSize `json:"Photos"`
}

type SetChat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type SetFile struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size"`
	FilePath string `json:"file_path"`
}

type SetDice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

type SetInputMediaPhoto struct {
	Type       string `json:"type"`
	Media      string `json:"media"`
	Caption    string `json:"caption"`
	Parse_mode string `json:"parse_mode"`
}

type SetInputMediaVideo struct {
	Type               string `json:"type"`
	Media              string `json:"media"`
	Thumb              string `json:"thumb"`
	Caption            string `json:"caption"`
	Parse_mode         string `json:"parse_mode"`
	Width              int    `json:"width"`
	Height             int    `json:"height"`
	Duration           int    `json:"duration"`
	Supports_streaming bool   `json:"supports_streaming"`
}

type SetInputMediaAnimation struct {
	Type string `json:"type"`
	Media string `json:"media"`
	Thumb string `json:"thumb"`
	Caption string `json:"caption"`
	Parse_mode string `json:"parse_mode"`
	Duration int `json:"duration"`
	Perfomer string `json:"performer"`
	Title string `json:"title"`
}

type SetInputMediaDocument struct {
	Type string `json:"type"`
	Media string `json:"media"`
	Thumb string `json:"thumb"`
	Caption string `json:"caption"`
	Parse_mode string `json:"parse_mode"`
}
type SetMediaGroup struct {
	ChatID               int                `json:"chat_id"`
	Media                SetInputMediaPhoto `json:"media"`
	Disable_Notification bool               `json:"disable_notification"`
	ReplyToMessageID     int                `json:"reply_to_message_id"`
}

type SetMessageEntities struct {
	Type     string  `json:"type"`
	Offset   int     `json:"offset"`
	Length   int     `json:"length"`
	Url      string  `json:"url"`
	User     SetUser `json:"user"`
	Language string  `json:"language"`
}

type SetReplyKeyboardMarkup struct {
	keyboard        []SetKeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool                `json:"resize_keyboard"`
	OneTimeKeyboard bool                `json:"one_time_keyboard"`
	SelectiveButton bool                `json:"selective"`
}

type SetKeyboardButton struct {
	Text            string                    `json:"text"`
	RequestContract bool                      `json:"request_contact"`
	RequestLocation bool                      `json:"request_location"`
	RequestPoll     SetKeyboardButtonPollType `json:"request_poll"`
}

type SetKeyboardButtonPollType struct {
	Type string `json:"type"`
}

type SetReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

type SetInlineKeyboardMarkup struct {
	InlineKeyboard [][]SetInlineKeyboardButton `json:"inline_keyboard"`
}

type SetInlineKeyboardButton struct {
	Text                         string          `json:"text"`
	Url                          string          `json:"url"`
	LoginUrl                     SetLoginUrl     `json:"login_url"`
	CallBackData                 string          `json:"callback_data"`
	SwitchInlineQuery            string          `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat string          `json:"switch_inline_query_current_chat"`
	CallBackGame                 SetCallBackGame `json:"callback_game"`
	Pay                          bool            `json:"pay"`
}

type SetCallBackGame struct {
	UserID             int    `json:"user_id"`
	Score              int    `json:"score"`
	Force              bool   `json:"force"`
	DisableEditMessage bool   `json:"disable_edit_message"`
	ChatID             int    `json:"chat_id"`
	MessageID          int    `json:"message_id"`
	InlineMessageID    string `json:"inline_message_id"`
}

type SetLoginUrl struct {
	Url                string `json:"url"`
	FowardText         string `jsong:"forward_text"`
	BotUsername        string `json:"bot_username"`
	RequestWriteAccess bool   `json:"request_write_access"`
}

type SetCallbackQuery struct {
	ID string `json:"id"`
	From SetUser `json:"from"`
	Message SetMessage `json:"message"`
	InlineMessageID string `json:"inline_message_id"`
	ChatInsance string `json:"chat_instance"`
	Data string `json:"data"`
	GameShortName string `json:"game_short_name"`
}


type SetForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective bool `json:"selective"`
}

type SetChatPhoto struct {
	SmallFileID string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID string `json:"big_file_id"`
	BigFileUniqueID string `json:"big_file_unique_id"`
}

type SetChatMember struct {
	User SetUser `json:"user"`
	Status string `json:"status"`
	CustomTitle string `json:"custom_title"`
	UntilDate int `json:"until_date"`
	CanBeEdited bool `json:"can_be_edited"`
	CanPostMessage bool `json:"can_post_messages"`
	CanEditMessage bool `json:"can_edit_messages"`
	CanDeleteMessage bool `json:"can_delete_messages"`
	CanRestrictMembers bool `json:"can_restrict_members"`
	CanPromoteMembers bool `json:"can_promote_members"`
	CanChangeInfo bool `json:"can_change_info"`
	CanInviteUsers bool `json:"can_invite_users"`
	CanPinMessage bool `json:"can_pin_messages"`
	IsMember bool `json:"is_member"`
	CanSendMessages bool `json:"can_send_messages"`
	CanSendMediaMessage bool `json:"can_send_media_messages"`
	CanSendPolls bool `json:"can_send_polls"`
	CanSendOtherMessages bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
}

type SetChatPermissions struct {
	CanSendMessages bool `json:"can_send_messages"`
	CanSendMediaMessage bool `json:"can_send_media_messages"`
	CanSendPolls bool `json:"can_send_polls"`
	CanSendOtherMessages bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo bool `json:"can_change_info"`
	CanInviteUsers bool `json:"can_invite_users"`
	CanPinMessage bool `json:"can_pin_messages"`
}

type SetBotCommand struct {
	Command string `json:"command"`
	Description string `json:"description"`
}

type SetResponseParameters struct {
	MigrateToChatID string `json:"migrate_to_chat_id"`
	RetryAfter int `json:"retry_after"`
}

type MeBot struct {
	Ok     bool `json:"ok"`
	Result struct {
		ID        int    `json:"id"`
		IsBot     bool   `json:"is_bot"`
		FirstName string `json:"first_name"`
		Username  string `json:"username"`
	} `json:"result"`
}

type MsgUpdater struct {
	Ok     bool `json:"ok"`
	Result []struct {
		UpdateID      int        `json:"update_id"`
		Message       SetMessage `json:"message"`
		EditedMessage SetMessage `json:"edited_message"`
		ChannelPost   SetMessage `json:"channel_post"`
		Poll          SetPoll    `json:"poll"`
	} `json:"result"`
}
