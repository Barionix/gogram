package telego

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

type PollOption struct {
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
	ID       string       `json:"id"`
	Question string       `json:"question"`
	Option   []PollOption `json:"poll_option"`
	IsClosed bool         `json:"is_closed"`
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
