package credential

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/util"
)

// 获取ticket的url  https://developer.work.weixin.qq.com/document/path/90506
const getQyWxTicketURL = "https://qyapi.weixin.qq.com/cgi-bin/get_jsapi_ticket?access_token=%s"
const getQyAppTicketURL = "https://qyapi.weixin.qq.com/cgi-bin/ticket/get?access_token=%s&type=agent_config"

// WorkJsTicket 默认获取js ticket方法
type WorkJsTicket struct {
	appID          string
	agentID        string
	cacheKeyPrefix string
	cache          cache.Cache
	//jsAPITicket 读写锁 同一个AppID一个
	jsAPITicketLock *sync.Mutex
}

// NewWorkJsTicket new
func NewWorkJsTicket(appID string, agentID string, cacheKeyPrefix string, cache cache.Cache) JsTicketHandle {
	return &WorkJsTicket{
		appID:           appID,
		agentID:         agentID,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		jsAPITicketLock: new(sync.Mutex),
	}
}

// GetTicket 获取企业微信jsapi_ticket
func (js *WorkJsTicket) GetTicket(accessToken string) (ticketStr string, err error) {
	//先从cache中取
	jsAPITicketCacheKey := fmt.Sprintf("%s_jsapi_ticket_%s", js.cacheKeyPrefix, js.appID)
	if val := js.cache.Get(jsAPITicketCacheKey); val != nil {
		return val.(string), nil
	}

	js.jsAPITicketLock.Lock()
	defer js.jsAPITicketLock.Unlock()

	// 双检，防止重复从微信服务器获取
	if val := js.cache.Get(jsAPITicketCacheKey); val != nil {
		return val.(string), nil
	}

	var ticket ResTicket
	ticket, err = GetQyWxTicketFromServer(accessToken, js.agentID != "")
	if err != nil {
		return
	}
	expires := ticket.ExpiresIn - 1500
	err = js.cache.Set(jsAPITicketCacheKey, ticket.Ticket, time.Duration(expires)*time.Second)
	ticketStr = ticket.Ticket
	return
}

// GetQyWxTicketFromServer 从企业微信服务器中获取ticket
func GetQyWxTicketFromServer(accessToken string, isApp bool) (ticket ResTicket, err error) {
	var response []byte
	url := fmt.Sprintf(getQyWxTicketURL, accessToken)
	if isApp {
		url = fmt.Sprintf(getQyAppTicketURL, accessToken)
	}
	response, err = util.HTTPGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &ticket)
	if err != nil {
		return
	}
	if ticket.ErrCode != 0 {
		err = fmt.Errorf("getTicket Error : errcode=%d , errmsg=%s", ticket.ErrCode, ticket.ErrMsg)
		return
	}
	return
}
