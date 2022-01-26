package user

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/context"
)

const (
	userInfoURL = "https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=%s&userid=%s"
	updateURL   = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=%s&department_id=%s&fetch_child=1"
	userListURL = "https://qyapi.weixin.qq.com/cgi-bin/user/get"
	launchCode  = "https://qyapi.weixin.qq.com/cgi-bin/get_launch_code?access_token=%s"
)

//User 用户管理
type User struct {
	*context.Context
}

//NewUser 实例化
func NewUser(context *context.Context) *User {
	user := new(User)
	user.Context = context
	return user
}

//Info 用户基本信息
type Info struct {
	util.CommonError
	Userid         string `json:"userid"`
	Name           string `json:"name"`
	Department     []int  `json:"department"`
	Order          []int  `json:"order"`
	Position       string `json:"position"`
	Mobile         string `json:"mobile"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	IsLeaderInDept []int  `json:"is_leader_in_dept"`
	Avatar         string `json:"avatar"`
	ThumbAvatar    string `json:"thumb_avatar"`
	Telephone      string `json:"telephone"`
	Alias          string `json:"alias"`
	Address        string `json:"address"`
	OpenUserid     string `json:"open_userid"`
	MainDepartment int    `json:"main_department"`
	Extattr        struct {
		Attrs []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
		} `json:"attrs"`
	} `json:"extattr"`
	Status           int    `json:"status"`
	QrCode           string `json:"qr_code"`
	ExternalPosition string `json:"external_position"`
	ExternalProfile  struct {
		ExternalCorpName string `json:"external_corp_name"`
		WechatChannels   struct {
			Nickname string `json:"nickname"`
			Status   int    `json:"status"`
		} `json:"wechat_channels"`
		ExternalAttr []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			Miniprogram struct {
				Appid    string `json:"appid"`
				Pagepath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram,omitempty"`
		} `json:"external_attr"`
	} `json:"external_profile"`
}

// OpenidList 用户列表
type OpenidList struct {
	util.CommonError

	Total int `json:"total"`
	Count int `json:"count"`
	Data  struct {
		OpenIDs []string `json:"openid"`
	} `json:"data"`
	NextOpenID string `json:"next_openid"`
}

//GetUserInfo 获取用户基本信息
func (user *User) GetUserInfo(userID string) (userInfo *Info, err error) {
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(userInfoURL, accessToken, userID)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	userInfo = new(Info)
	err = json.Unmarshal(response, userInfo)
	if err != nil {
		return
	}
	if userInfo.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfo Error , errcode=%d , errmsg=%s", userInfo.ErrCode, userInfo.ErrMsg)
		return
	}
	return
}

// Update 更新员工资料
func (user *User) Update(userID, external_position string) (err error) {
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(updateURL, accessToken, userID)
	var response []byte
	response, err = util.PostJSON(uri, map[string]string{"userid": userID, "external_position": external_position})
	if err != nil {
		return
	}

	return util.DecodeWithCommonError(response, "updateURL")
}

// ListUserOpenIDs 返回用户列表
func (user *User) ListUserOpenIDs(nextOpenid ...string) (*OpenidList, error) {
	accessToken, err := user.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri, _ := url.Parse(userListURL)
	q := uri.Query()
	q.Set("access_token", accessToken)
	if len(nextOpenid) > 0 && nextOpenid[0] != "" {
		q.Set("next_openid", nextOpenid[0])
	}
	uri.RawQuery = q.Encode()

	response, err := util.HTTPGet(uri.String())
	if err != nil {
		return nil, err
	}

	userlist := OpenidList{}

	err = util.DecodeWithError(response, &userlist, "ListUserOpenIDs")
	if err != nil {
		return nil, err
	}

	return &userlist, nil
}

// ListAllUserOpenIDs 返回所有用户OpenID列表
func (user *User) ListAllUserOpenIDs() ([]string, error) {
	nextOpenid := ""
	openids := make([]string, 0)
	count := 0
	for {
		ul, err := user.ListUserOpenIDs(nextOpenid)
		if err != nil {
			return nil, err
		}
		openids = append(openids, ul.Data.OpenIDs...)
		count += ul.Count
		if ul.Total > count {
			nextOpenid = ul.NextOpenID
		} else {
			return openids, nil
		}
	}
}

type RespLaunchCode struct {
	util.CommonError
	LaunchCode string `json:"launch_code"`
}

//GetLaunchCode 用于打开个人聊天窗口schema
func (user *User) GetLaunchCode(userID, other string) (userInfo *RespLaunchCode, err error) {
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(launchCode, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, map[string]interface{}{"operator_userid": userID, "single_chat": map[string]string{"userid": other}})
	if err != nil {
		return
	}
	userInfo = new(RespLaunchCode)
	err = json.Unmarshal(response, userInfo)
	if err != nil {
		return
	}
	if userInfo.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfo Error , errcode=%d , errmsg=%s", userInfo.ErrCode, userInfo.ErrMsg)
		return
	}
	return
}
