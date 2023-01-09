package message

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/context"
	"strconv"
)

const (
	messageSendURL               = "https://qyapi.weixin.qq.com/cgi-bin/message/send"
	messageUpdateTemplateCardURL = "https://api.weixin.qq.com/cgi-bin/message/update_template_card"
	messageDelURL                = "https://api.weixin.qq.com/cgi-bin/message/recall"
)

// Client 应用消息
type Client struct {
	*context.Context
}

// NewClient 实例化
func NewClient(ctx *context.Context) *Client {
	return &Client{
		ctx,
	}
}

// AppMessage 发送的模板消息内容
type AppMessage struct {
	ToUser                 string  `json:"touser"`  // 必须, 成员ID列表（多个接收者用‘|’分隔，最多支持1000个 ,指定为”@all”，则向该企业应用的全部成员发送
	Toparty                string  `json:"toparty"` //部门ID列表,当touser为”@all”时忽略本参数
	Totag                  string  `json:"totag"`   //标签ID列表,当touser为”@all”时忽略本参数
	Msgtype                MsgType `json:"msgtype"`
	Agentid                int     `json:"agentid"`
	Safe                   int     `json:"safe"`
	EnableIdTrans          int     `json:"enable_id_trans"`
	EnableDuplicateCheck   int     `json:"enable_duplicate_check"`
	DuplicateCheckInterval int     `json:"duplicate_check_interval"`
	Text                   *Text   `json:"text"`
	*Image                 `json:"image"`
	*Voice                 `json:"voice"`
	*Video                 `json:"video"`
	File                   *PushFile     `json:"file"`
	TextCard               *PushTextCard `json:"textcard"`
	News                   *News         `json:"news"`
	MpNews                 *MpNews       `json:"mpnews"`
	Markdown               *Text         `json:"markdown"`
	//todo(hb1707) 可能会发生变化的字段直接用interface{}了
	MiniprogramNotice interface{} `json:"miniprogram_notice"`
	TemplateCard      interface{} `json:"template_card"`
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
func (tpl *Client) Send(msg *AppMessage) (msgID string, responseCode string, err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", messageSendURL, accessToken)
	var response []byte
	if msg.Agentid == 0 {
		msg.Agentid, _ = strconv.Atoi(tpl.Context.AgentID)
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
func (tpl *Client) UpdateTemplate(msg *TemplateUpdate) (msgID string, err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
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
func (tpl *Client) Recall(msgID int64) (err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
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
