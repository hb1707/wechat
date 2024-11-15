package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	listUrl           = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list"
	getUrl            = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get"
	getByUserBatchUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/batch/get_by_user"
)

type ReqGetByUser struct {
	UseridList []string `json:"userid_list"`
	Cursor     string   `json:"cursor"`
	Limit      int      `json:"limit"`
}
type OneUser struct {
	util.CommonError
	ExternalContact ExternalContact `json:"external_contact"`
	FollowUser      []FollowInfo    `json:"follow_user"` //注意，仅获取单个客户详情的时候这里返回的是跟进人列表
	NextCursor      string          `json:"next_cursor"`
}
type resUserList struct {
	util.CommonError
	ExternalContactList []UserInfo `json:"external_contact_list"`
	NextCursor          string     `json:"next_cursor"`
}
type resUserids struct {
	util.CommonError
	ExternalUserid []string `json:"external_userid"`
}

type UserInfo struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowInfo      FollowInfo      `json:"follow_info"` //企业成员客户跟进人信息，可以参考获取客户详情，但标签信息只会返回企业标签和规则组标签的tag_id，个人标签将不再返回
}

// GetUseridList 获取我的客户列表
func (tpl *Client) GetUseridList(myUserid string) (externalUserid []string, err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&userid=%s", listUrl, accessToken, myUserid)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result resUserids
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("template msg send error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	externalUserid = result.ExternalUserid
	return
}

// GetUseridList 获取我的全部客户列表及详情
func (tpl *Client) GetQyUserInfoList(qyUserid []string) ([]UserInfo, error) {
	var userInfoList []UserInfo
	var req ReqGetByUser
	req.UseridList = qyUserid
	req.Limit = 100
	for {
		userInfoPage, resCursor, err := tpl.GetUserInfoListByUserIds(req)
		if err != nil {
			return userInfoList, err
		}
		userInfoList = append(userInfoList, userInfoPage...)
		if resCursor != "" {
			req.Cursor = resCursor
		} else {
			break
		}
	}
	return userInfoList, nil
}

// GetUserInfoAndAllFollow 获取客户详情以及全部跟进人
func (tpl *Client) GetUserInfoAndAllFollow(userid string) (OneUser, error) {
	var result, res OneUser
	var err error
	var cursor string
	for {
		res, err = tpl.GetUserInfo(userid, cursor)
		if err != nil {
			return result, err
		}
		result.FollowUser = append(result.FollowUser, res.FollowUser...)
		result.ExternalContact = res.ExternalContact
		if res.NextCursor != "" {
			cursor = res.NextCursor
		} else {
			break
		}
	}
	return result, nil
}

// GetUserInfo 获取客户详情
func (tpl *Client) GetUserInfo(externalUserid string, cursor ...string) (result OneUser, err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	var page = ""
	if len(cursor) > 0 {
		page = cursor[0]
	}
	uri := fmt.Sprintf("%s?access_token=%s&external_userid=%s&cursor=%s", getUrl, accessToken, externalUserid, page)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
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

// GetUserInfoListByUserId 批量获取客户详情
func (tpl *Client) GetUserInfoListByUserIds(req ReqGetByUser) (userList []UserInfo, nextCursor string, err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", getByUserBatchUrl, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, req)
	if err != nil {
		return
	}
	var result resUserList
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("template msg send error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	userList = result.ExternalContactList
	nextCursor = result.NextCursor
	return
}
