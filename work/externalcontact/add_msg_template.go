package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	addMsgTemplateUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_msg_template"
)

type ChatType string

const (
	ChatTypeSingle ChatType = "single"
	ChatTypeGroup  ChatType = "group"
)

// ReqMessage 企业群发参数
type ReqMessage struct {
	ChatType       ChatType `json:"chat_type"`       //群发任务的类型，默认为single，表示发送给客户，group表示发送给客户群
	ExternalUserid []string `json:"external_userid"` // 客户的外部联系人id列表，仅在chat_type为single时有效，不可与sender同时为空，最多可传入1万个客户
	Sender         string   `json:"sender"`          //发送企业群发消息的成员userid，当类型为发送给客户群时必填
	Text           struct {
		Content string `json:"content"`
	} `json:"text"`
	Attachments []struct {
		Msgtype     string         `json:"msgtype"`
		Image       MsgImage       `json:"image"`
		Link        MsgLink        `json:"link"`
		Miniprogram MsgMiniprogram `json:"miniprogram"`
		Video       MsgVideo       `json:"video"`
		File        MsgFile        `json:"file"`
	} `json:"attachments"`
}
type MsgImage struct {
	MediaId string `json:"media_id"`
	PicUrl  string `json:"pic_url"`
}
type MsgLink struct {
	Title  string `json:"title"`
	Picurl string `json:"picurl"`
	Desc   string `json:"desc"`
	Url    string `json:"url"`
}
type MsgMiniprogram struct {
	Title      string `json:"title"`
	PicMediaId string `json:"pic_media_id"`
	Appid      string `json:"appid"`
	Page       string `json:"page"`
}
type MsgVideo struct {
	MediaId string `json:"media_id"`
}
type MsgFile struct {
	MediaId string `json:"media_id"`
}

type resTemplateSend struct {
	util.CommonError
	FailList string `json:"fail_list"`
	MsgID    int64  `json:"msgid"`
}

// Send 发送应用消息
func (r *Client) Send(msg *ReqMessage) (msgID int64, err error) {
	var accessToken string
	accessToken, err = r.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", addMsgTemplateUrl, accessToken)
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
