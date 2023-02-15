package message

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
	"strconv"
)

const (
	messageSendURL               = "https://qyapi.weixin.qq.com/cgi-bin/message/send"
	messageUpdateTemplateCardURL = "https://api.weixin.qq.com/cgi-bin/message/update_template_card"
	messageDelURL                = "https://api.weixin.qq.com/cgi-bin/message/recall"
)

// UpdateButton 模板卡片按钮
type UpdateButton struct {
	CommonToken `json:"-"`
	Button      struct {
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
	CommonToken  `json:"-"`
	TemplateCard interface{} `xml:"TemplateCard" json:"template_card"`
}

// NewTemplateCard 更新点击用户的整张卡片
func NewTemplateCard(cardXml interface{}) *TemplateCard {
	card := new(TemplateCard)
	card.TemplateCard = cardXml
	return card
}

// AppMessage 发送的模板消息内容
type AppMessage struct {
	ToUser                 string  `json:"touser"`            // 必须, 成员ID列表（多个接收者用‘|’分隔，最多支持1000个 ,指定为”@all”，则向该企业应用的全部成员发送
	Toparty                string  `json:"toparty,omitempty"` //部门ID列表,当touser为”@all”时忽略本参数
	Totag                  string  `json:"totag,omitempty"`   //标签ID列表,当touser为”@all”时忽略本参数
	Msgtype                MsgType `json:"msgtype"`
	Agentid                int     `json:"agentid"`
	Safe                   int     `json:"safe,omitempty"`
	EnableIdTrans          int     `json:"enable_id_trans,omitempty"`
	EnableDuplicateCheck   int     `json:"enable_duplicate_check,omitempty"`
	DuplicateCheckInterval int     `json:"duplicate_check_interval,omitempty"`
	Text                   *Text   `json:"text,omitempty"`
	*Image                 `json:"image,omitempty"`
	*Voice                 `json:"voice,omitempty"`
	*Video                 `json:"video,omitempty"`
	File                   *PushFile     `json:"file,omitempty"`
	TextCard               *PushTextCard `json:"textcard,omitempty"`
	News                   *News         `json:"news,omitempty"`
	MpNews                 *MpNews       `json:"mpnews,omitempty"`
	Markdown               *Text         `json:"markdown,omitempty"`
	//todo(wind) 可能会发生变化的字段直接用interface{}了
	MiniprogramNotice interface{}         `json:"miniprogram_notice,omitempty"`
	TemplateCard      *TemplateCardButton `json:"template_card,omitempty"`
}
type Action struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}
type HorizontalContent struct {
	Keyname string `json:"keyname"`
	Value   string `json:"value,omitempty"`
	Type    int    `json:"type,omitempty"`
	Url     string `json:"url,omitempty"`
	MediaId string `json:"media_id,omitempty"`
	Userid  string `json:"userid,omitempty"`
}
type VerticalContent struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type Jump struct {
	Type     int    `json:"type,omitempty"`
	Title    string `json:"title"`
	Url      string `json:"url,omitempty"`
	Appid    string `json:"appid,omitempty"`
	Pagepath string `json:"pagepath,omitempty"`
}
type OptionCheckBox struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	IsChecked bool   `json:"is_checked"`
}
type Option struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}
type Select struct {
	QuestionKey string   `json:"question_key"`
	Title       string   `json:"title"`
	SelectedId  string   `json:"selected_id"`
	OptionList  []Option `json:"option_list"`
}
type Button struct {
	Text  string `json:"text"`
	Style int    `json:"style,omitempty"`
	Key   string `json:"key,omitempty"`
}

type TemplateCardButton struct {
	CardType string `json:"card_type"`
	Source   struct {
		IconUrl   string `json:"icon_url,omitempty"`
		Desc      string `json:"desc,omitempty"`
		DescColor int    `json:"desc_color,omitempty"`
	} `json:"source,omitempty"`

	ActionMenu struct {
		Desc       string   `json:"desc,omitempty"`
		ActionList []Action `json:"action_list"`
	} `json:"action_menu,omitempty"`
	MainTitle struct {
		Title string `json:"title"`
		Desc  string `json:"desc,omitempty"`
	} `json:"main_title,omitempty"`
	QuoteArea struct {
		Type      int    `json:"type,omitempty"`
		Url       string `json:"url,omitempty"`
		Title     string `json:"title,omitempty"`
		QuoteText string `json:"quote_text,omitempty"`
	} `json:"quote_area"`
	SubTitleText          string              `json:"sub_title_text,omitempty"`
	HorizontalContentList []HorizontalContent `json:"horizontal_content_list"`
	VerticalContentList   []VerticalContent   `json:"vertical_content_list,omitempty"`
	CardAction            struct {
		Type     int    `json:"type,omitempty"`
		Url      string `json:"url,omitempty"`
		Appid    string `json:"appid,omitempty"`
		Pagepath string `json:"pagepath,omitempty"`
	} `json:"card_action"`
	JumpList        []Jump `json:"jump_list,omitempty"`
	EmphasisContent struct {
		Title string `json:"title,omitempty"`
		Desc  string `json:"desc,omitempty"`
	} `json:"emphasis_content,omitempty"`
	ImageTextArea struct {
		Type     int    `json:"type"`
		Url      string `json:"url"`
		Title    string `json:"title"`
		Desc     string `json:"desc"`
		ImageUrl string `json:"image_url"`
	} `json:"image_text_area,omitempty"`
	CardImage struct {
		Url         string  `json:"url"`
		AspectRatio float64 `json:"aspect_ratio"`
	} `json:"card_image,omitempty"`
	Checkbox struct {
		QuestionKey string           `json:"question_key"`
		OptionList  []OptionCheckBox `json:"option_list"`
		Mode        int              `json:"mode"`
	} `json:"checkbox,omitempty"`
	SelectList      []Select `json:"select_list"`
	TaskId          string   `json:"task_id"`
	ButtonSelection struct {
		QuestionKey string   `json:"question_key"`
		Title       string   `json:"title,omitempty"`
		OptionList  []Option `json:"option_list,omitempty"`
		SelectedId  string   `json:"selected_id,omitempty"`
	} `json:"button_selection,omitempty"`
	ButtonList []Button `json:"button_list"`
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

// Send 发送应用消息
func (r *Client) Send(msg *AppMessage) (msgID string, responseCode string, err error) {
	var accessToken string
	accessToken, err = r.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", messageSendURL, accessToken)
	var response []byte
	if msg.Agentid == 0 {
		msg.Agentid, _ = strconv.Atoi(r.Context.AgentID)
	}
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
	responseCode = result.ResponseCode
	return
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
