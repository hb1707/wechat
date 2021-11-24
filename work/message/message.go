package message

import (
	"encoding/xml"

	"github.com/silenceper/wechat/v2/officialaccount/device"
)

// MsgType 企业微信普通消息类型
type MsgType string

// EventType 企业微信事件消息类型
type EventType string

// InfoType 第三方平台授权事件类型
type InfoType string

const (
	//MsgTypeEvent 表示事件推送消息 [限接收]
	MsgTypeEvent = "event"

	//MsgTypeText 表示文本消息
	MsgTypeText MsgType = "text"
	//MsgTypeImage 表示图片消息
	MsgTypeImage = "image"
	//MsgTypeVoice 表示语音消息
	MsgTypeVoice = "voice"
	//MsgTypeVideo 表示视频消息
	MsgTypeVideo = "video"
	//MsgTypeNews 表示图文消息[限回复与发送应用消息]
	MsgTypeNews = "news"

	//MsgTypeLink 表示链接消息[限接收]
	MsgTypeLink = "link"
	//MsgTypeLocation 表示坐标消息[限接收]
	MsgTypeLocation = "location"

	//MsgTypeUpdateButton 更新点击用户的按钮文案[限回复应用消息]
	MsgTypeUpdateButton = "update_button"
	//MsgTypeUpdateTemplateCard 更新点击用户的整张卡片[限回复应用消息]
	MsgTypeUpdateTemplateCard = "update_template_card"

	//MsgTypeFile 文件消息[限发送应用消息]
	MsgTypeFile = "file"
	//MsgTypeTextCard 文本卡片消息[限发送应用消息]
	MsgTypeTextCard = "textcard"
	//MsgTypeMpNews 图文消息[限发送应用消息] 跟普通的图文消息一致，唯一的差异是图文内容存储在企业微信
	MsgTypeMpNews = "mpnews"
	//MsgTypeMarkdown markdown消息[限发送应用消息]
	MsgTypeMarkdown = "markdown"
	//MsgTypeMiniprogramNotice 小程序通知消息[限发送应用消息]
	MsgTypeMiniprogramNotice = "miniprogram_notice"
	//MsgTypeTemplateCard 模板卡片消息[限发送应用消息]
	MsgTypeTemplateCard = "template_card"
)

const (
	//EventSubscribe 成员关注，成员已经加入企业，管理员添加成员到应用可见范围(或移除可见范围)时
	EventSubscribe EventType = "subscribe"
	//EventUnsubscribe 成员取消关注，成员已经在应用可见范围，成员加入(或退出)企业时
	EventUnsubscribe = "unsubscribe"
	//EventEnterAgent 本事件在成员进入企业微信的应用时触发
	EventEnterAgent = "enter_agent"
	//EventLocation 上报地理位置事件
	EventLocation = "LOCATION"
	//EventBatchJobResult 异步任务完成事件推送
	EventBatchJobResult = "batch_job_result"
	//EventClick 点击菜单拉取消息时的事件推送
	EventClick = "click"
	//EventView 点击菜单跳转链接时的事件推送
	EventView = "view"
	//EventScancodePush 扫码推事件的事件推送
	EventScancodePush = "scancode_push"
	//EventScancodeWaitmsg 扫码推事件且弹出“消息接收中”提示框的事件推送
	EventScancodeWaitmsg = "scancode_waitmsg"
	//EventPicSysphoto 弹出系统拍照发图的事件推送
	EventPicSysphoto = "pic_sysphoto"
	//EventPicPhotoOrAlbum 弹出拍照或者相册发图的事件推送
	EventPicPhotoOrAlbum = "pic_photo_or_album"
	//EventPicWeixin 弹出微信相册发图器的事件推送
	EventPicWeixin = "pic_weixin"
	//EventLocationSelect 弹出地理位置选择器的事件推送
	EventLocationSelect = "location_select"

	//EventOpenApprovalChange 审批状态通知事件推送
	EventOpenApprovalChange = "open_approval_change"

	//EventShareAgentChange 共享应用事件回调
	EventShareAgentChange = "share_agent_change"

	//EventTemplateCard 模板卡片事件推送
	EventTemplateCard = "template_card_event"

	//EventTemplateCardMenu 通用模板卡片右上角菜单事件推送
	EventTemplateCardMenu = "template_card_menu_event"

	//EventChangeExternalContact 企业客户事件推送
	//add_external_contact 添加
	//edit_external_contact 编辑
	//add_half_external_contact 免验证添加
	//del_external_contact 员工删除客户
	//del_follow_user 客户删除跟进员工
	//transfer_fail 企业将客户分配给新的成员接替后，客户添加失败
	//change_external_chat 客户群创建事件
	EventChangeExternalContact = "change_external_contact"

	//EventChangeExternalChat 企业客户群变更事件推送
	//create 客户群创建
	//update 客户群变更
	//dismiss 客户群解散
	EventChangeExternalChat = "change_external_chat"

	//EventChangeExternalTag 企业客户标签创建事件推送
	//create 创建标签
	//update 变更标签
	//delete 删除标签
	//shuffle 重新排序
	EventChangeExternalTag = "change_external_tag"

	//EventKfMsg 企业微信客服回调事件
	EventKfMsg = "kf_msg_or_event"
	//EventLivingStatusChange 直播回调事件
	EventLivingStatusChange = "living_status_change"

	//EventMsgauditNotify 会话内容存档开启后，产生会话回调事件
	EventMsgauditNotify = "msgaudit_notify"
)

//todo 第三方应用开发
/*const (
    //微信开放平台需要用到

    // InfoTypeVerifyTicket 返回ticket
    InfoTypeVerifyTicket InfoType = "component_verify_ticket"
    // InfoTypeAuthorized 授权
    InfoTypeAuthorized = "authorized"
    // InfoTypeUnauthorized 取消授权
    InfoTypeUnauthorized = "unauthorized"
    // InfoTypeUpdateAuthorized 更新授权
    InfoTypeUpdateAuthorized = "updateauthorized"
)*/

//MixMessage 存放所有企业微信官方发送过来的消息和事件
type MixMessage struct {
	CommonToken

	//接收普通消息
	MsgID   int64 `xml:"MsgId"`   //其他消息推送过来是MsgId
	AgentID int   `xml:"AgentID"` //企业应用的id，整型。可在应用的设置页面查看

	Content      string `xml:"Content"`      //文本消息内容
	Format       string `xml:"Format"`       //语音消息格式，如amr，speex等
	ThumbMediaID string `xml:"ThumbMediaId"` //视频消息缩略图的媒体id，可以调用获取媒体文件接口拉取数据，仅三天内有效

	Title       string `xml:"Title"`       //链接消息，标题
	Description string `xml:"Description"` //链接消息，描述
	URL         string `xml:"Url"`         //链接消息，链接跳转的url

	PicURL  string `xml:"PicUrl"`  ////图片消息或者链接消息，封面缩略图的url
	MediaID string `xml:"MediaId"` //图片媒体文件id//语音媒体文件id//视频消息缩略图的媒体id，可以调用获取媒体文件接口拉取，仅三天内有效

	LocationX float64 `xml:"Location_X"` //位置消息，地理位置纬度
	LocationY float64 `xml:"Location_Y"` //位置消息，地理位置经度
	Scale     float64 `xml:"Scale"`      //位置消息，地图缩放大小
	Label     string  `xml:"Label"`      //位置消息，地理位置信息

	AppType string `xml:"AppType"` //接收地理位置时存在，app类型，在企业微信固定返回wxwork，在微信不返回该字段

	//TemplateMsgID int64   `xml:"MsgID"` //模板消息推送成功的消息是MsgID
	///Recognition   string  `xml:"Recognition"`

	//事件相关
	Event    EventType `xml:"Event"`
	EventKey string    `xml:"EventKey"`

	//仅上报地理位置事件
	Latitude  string `xml:"Latitude"`  //地理位置纬度
	Longitude string `xml:"Longitude"` //地理位置经度
	Precision string `xml:"Precision"` //地理位置精度

	//仅异步任务完成事件
	JobId   string `xml:"JobId"`   //异步任务id，最大长度为64字符
	JobType string `xml:"JobType"` //异步任务，操作类型，字符串，目前分别有：sync_user(增量更新成员)、 replace_user(全量覆盖成员）、invite_user(邀请成员关注）、replace_party(全量覆盖部门)
	ErrCode int    `xml:"ErrCode"` //异步任务，返回码
	ErrMsg  string `xml:"ErrMsg"`  //异步任务，对返回码的文本描述内容

	//开启通讯录回调通知 https://open.work.weixin.qq.com/api/doc/90000/90135/90967

	// todo 第三方平台相关 字段名可能不准确
	/*InfoType                     InfoType `xml:"InfoType"`
	  AppID                        string   `xml:"AppId"`
	  ComponentVerifyTicket        string   `xml:"ComponentVerifyTicket"`
	  AuthorizerAppid              string   `xml:"AuthorizerAppid"`
	  AuthorizationCode            string   `xml:"AuthorizationCode"```````````````````````````````````````
	  AuthorizationCodeExpiredTime int64    `xml:"AuthorizationCodeExpiredTime"`
	  PreAuthCode                  string   `xml:"PreAuthCode"`*/

	//设备相关
	device.MsgDevice
}

//EventPic 发图事件推送
type EventPic struct {
	PicMd5Sum string `xml:"PicMd5Sum"`
}

//EncryptedXMLMsg 安全模式下的消息体
type EncryptedXMLMsg struct {
	XMLName      struct{} `xml:"xml" json:"-"`
	ToUserName   string   `xml:"ToUserName" json:"ToUserName"`
	AgentID      string   `xml:"AgentID" json:"AgentID"`
	EncryptedMsg string   `xml:"Encrypt"    json:"Encrypt"`
}

//ResponseEncryptedXMLMsg 需要返回的消息体
type ResponseEncryptedXMLMsg struct {
	XMLName      struct{} `xml:"xml" json:"-"`
	EncryptedMsg string   `xml:"Encrypt"      json:"Encrypt"`
	MsgSignature string   `xml:"MsgSignature" json:"MsgSignature"`
	Timestamp    int64    `xml:"TimeStamp"    json:"TimeStamp"`
	Nonce        string   `xml:"Nonce"        json:"Nonce"`
}

// CDATA  使用该类型,在序列化为 xml 文本时文本会被解析器忽略
type CDATA string

// MarshalXML 实现自己的序列化方法
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

// CommonToken 消息中通用的结构
type CommonToken struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      MsgType  `xml:"MsgType"`
}

//SetToUserName set ToUserName
func (msg *CommonToken) SetToUserName(toUserName CDATA) {
	msg.ToUserName = toUserName
}

//SetFromUserName set FromUserName
func (msg *CommonToken) SetFromUserName(fromUserName CDATA) {
	msg.FromUserName = fromUserName
}

//SetCreateTime set createTime
func (msg *CommonToken) SetCreateTime(createTime int64) {
	msg.CreateTime = createTime
}

//SetMsgType set MsgType
func (msg *CommonToken) SetMsgType(msgType MsgType) {
	msg.MsgType = msgType
}

//GetOpenID get the FromUserName value
func (msg *CommonToken) GetOpenID() string {
	return string(msg.FromUserName)
}
