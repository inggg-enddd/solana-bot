package ports

import (
	"context"

	"github.com/your/project/internal/bot/domain"
)

// MessageHandler 定义消息处理器接口
type MessageHandler interface {
	Handle(ctx context.Context, msg *domain.BotMessage) (*domain.BotResponse, error)
}

// CommandRegistry 定义命令注册接口
type CommandRegistry interface {
	Register(command domain.Command, handler MessageHandler)
	Get(command domain.Command) (MessageHandler, bool)
}

// BotService 定义Bot服务接口
type BotService interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Send(ctx context.Context, response *domain.BotResponse) error
}
