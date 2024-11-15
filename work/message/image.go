package message

//Image 图片消息
type Image struct {
	CommonToken `json:"-"`
	Image       struct {
		MediaID string `xml:"MediaId" json:"media_id"`
	} `xml:"Image" json:"image"`
}

//NewImage 回复图片消息
func NewImage(mediaID string) *Image {
	image := new(Image)
	image.Image.MediaID = mediaID
	return image
}
