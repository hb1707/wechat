package work

import (
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/work/config"
	"github.com/silenceper/wechat/v2/work/context"
	"github.com/silenceper/wechat/v2/work/kf"
	"github.com/silenceper/wechat/v2/work/msgaudit"
	"github.com/silenceper/wechat/v2/work/oauth"
	"github.com/silenceper/wechat/v2/work/server"
	"github.com/silenceper/wechat/v2/work/tools"
	"github.com/silenceper/wechat/v2/work/user"
	"net/http"
)

// Work 企业微信
type Work struct {
	ctx *context.Context
}

//NewWork init work
func NewWork(cfg *config.Config) *Work {
	defaultAkHandle := credential.NewWorkAccessToken(cfg.CorpID, cfg.CorpSecret, credential.CacheKeyWorkPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &Work{ctx: ctx}
}

//GetContext get Context
func (wk *Work) GetContext() *context.Context {
	return wk.ctx
}

// GetServer 消息管理：接收事件，被动回复消息管理
func (wk *Work) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	srv := server.NewServer(wk.ctx)
	srv.Request = req
	srv.Writer = writer
	return srv
}

//GetOauth get oauth
func (wk *Work) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(wk.ctx)
}

// GetMsgAudit get msgAudit
func (wk *Work) GetMsgAudit() (*msgaudit.Client, error) {
	return msgaudit.NewClient(wk.ctx.Config)
}

// GetKF get kf
func (wk *Work) GetKF() (*kf.Client, error) {
	return kf.NewClient(wk.ctx.Config)
}

//GetUser get user
func (wk *Work) GetUser() *user.User {
	return user.NewUser(wk.ctx)
}

//GetCalendar get calendar
func (wk *Work) GetCalendar() *tools.Calendar {
	return tools.NewCalendar(wk.ctx)
}
