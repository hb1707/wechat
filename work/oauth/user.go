package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	code2SessionURL = "https://qyapi.weixin.qq.com/cgi-bin/miniprogram/jscode2session?access_token=%s&js_code=%s&grant_type=authorization_code"
	launchCode      = "https://qyapi.weixin.qq.com/cgi-bin/get_launch_code?access_token=%s"
)

func (ctr *Oauth) Code2Session(code string) (result ResUserInfo, err error) {
	var accessToken string
	accessToken, err = ctr.GetAccessToken()
	if err != nil {
		return
	}
	var response []byte
	response, err = util.HTTPGet(
		fmt.Sprintf(code2SessionURL, accessToken, code),
	)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

type RespLaunchCode struct {
	util.CommonError
	LaunchCode string `json:"launch_code"`
}

// GetLaunchCode 用于打开个人聊天窗口schema
func (ctr *Oauth) GetLaunchCode(userID, other string) (userInfo *RespLaunchCode, err error) {
	var accessToken string
	accessToken, err = ctr.GetAccessToken()
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
