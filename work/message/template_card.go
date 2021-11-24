package message

//UpdateButton 模板卡片按钮
type UpdateButton struct {
	CommonToken `json:"-"`
	Button      struct {
		ReplaceName string `xml:"ReplaceName" json:"replace_name"`
	} `xml:"Button" json:"button"`
}

//NewUpdateButton 更新点击用户的按钮文案
func NewUpdateButton(replaceName string) *UpdateButton {
	btn := new(UpdateButton)
	btn.Button.ReplaceName = replaceName
	return btn
}

//TemplateCard 被动回复模板卡片
//https://open.work.weixin.qq.com/api/doc/90000/90135/90241
type TemplateCard struct {
	CommonToken  `json:"-"`
	TemplateCard interface{} `xml:"TemplateCard" json:"template_card"`
}

// NewTemplateCard 更新点击用户的整张卡片
func NewTemplateCard(cardXml interface{}) *TemplateCard {
	card := new(TemplateCard)
	card.TemplateCard = cardXml
	return card
}
