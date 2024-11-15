package message

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	messageUpdateTemplateCardURL = "https://api.weixin.qq.com/cgi-bin/message/update_template_card"
	messageDelURL                = "https://api.weixin.qq.com/cgi-bin/message/recall"
)

// UpdateButton 模板卡片按钮
type UpdateButton struct {
	//CommonToken `json:"-"`
	Button struct {
		ReplaceName string `xml:"ReplaceName" json:"replace_name"`
	} `xml:"Button" json:"button"`
}

// NewUpdateButton 更新点击用户的按钮文案
func NewUpdateButton(replaceName string) *UpdateButton {
	btn := new(UpdateButton)
	btn.Button.ReplaceName = replaceName
	return btn
}

// TemplateCard 被动回复模板卡片
// https://open.work.weixin.qq.com/api/doc/90000/90135/90241
type TemplateCard struct {
	//CommonToken  `json:"-"`
	TemplateCard interface{} `xml:"TemplateCard" json:"template_card"`
}

// NewTemplateCard 更新点击用户的整张卡片
func NewTemplateCard(cardXml interface{}) *TemplateCard {
	card := new(TemplateCard)
	card.TemplateCard = cardXml
	return card
}

type PushFile struct {
	MediaID string `json:"media_id"`
}
type PushTextCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Btntxt      string `json:"btntxt"`
}

type resTemplateSend struct {
	util.CommonError
	Invaliduser  string `json:"invaliduser"`   //不合法的userid，不区分大小写，统一转为小写
	Invalidparty string `json:"invalidparty"`  //不合法的partyid
	Invalidtag   string `json:"invalidtag"`    //不合法的标签id
	MsgID        string `json:"msgid"`         //消息id，用于撤回应用消息
	ResponseCode string `json:"response_code"` //仅消息类型为“按钮交互型”，“投票选择型”和“多项选择型”的模板卡片消息返回，应用可使用response_code调用更新模版卡片消息接口，24小时内有效，且只能使用一次
}

// TemplateUpdate  更新模版卡片消息内容
type TemplateUpdate struct {
	Userids      []string `json:"userids"`
	Partyids     []int    `json:"partyids"`
	Tagids       []int    `json:"tagids"`
	Atall        int      `json:"atall"`
	Agentid      int      `json:"agentid"`
	ResponseCode string   `json:"response_code"`
	*UpdateButton
	*TemplateCard
}

// UpdateTemplate 更新模版卡片消息
func (r *Client) UpdateTemplate(msg *TemplateUpdate) (msgID string, err error) {
	var accessToken string
	accessToken, err = r.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", messageUpdateTemplateCardURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, msg)
	if err != nil {
		return
	}
	var result resTemplateSend
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("template msg send error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	msgID = result.MsgID
	return
}

type ReqRecall struct {
	MsgID int64 `json:"msgid"`
}

// Recall 撤回应用消息
func (r *Client) Recall(msgID int64) (err error) {
	var accessToken string
	accessToken, err = r.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", messageDelURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, &ReqRecall{
		MsgID: msgID,
	})
	if err != nil {
		return
	}
	var result util.CommonError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("template msg send error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}
