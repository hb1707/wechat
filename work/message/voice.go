package message

//Voice 语音消息
type Voice struct {
	CommonToken `json:"-"`
	Voice       struct {
		MediaID string `xml:"MediaId" json:"media_id"`
	} `xml:"Voice" json:"voice"`
}

//NewVoice 回复语音消息
func NewVoice(mediaID string) *Voice {
	voice := new(Voice)
	voice.Voice.MediaID = mediaID
	return voice
}
