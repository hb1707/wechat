package message

import (
	"github.com/silenceper/wechat/v2/work/context"
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
