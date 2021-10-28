package tools

import (
	"github.com/silenceper/wechat/v2/work/context"
)

const (
	calendarURL = "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/get?access_token=%s"
)

//Calendar 日历管理
type Calendar struct {
	*context.Context
}

//NewCalendar 实例化
func NewCalendar(context *context.Context) *Calendar {
	calendar := new(Calendar)
	calendar.Context = context
	return calendar
}
