package service

import (
	"context"

	"github.com/inggg-enddd/solana-bot/internal/bot/domain"
	"github.com/inggg-enddd/solana-bot/internal/bot/ports"
)

type TelegramBotService struct {
	registry ports.CommandRegistry
	// 其他依赖
}

func NewTelegramBotService(registry ports.CommandRegistry) *TelegramBotService {
	return &TelegramBotService{
		registry: registry,
	}
}

func (s *TelegramBotService) Start(ctx context.Context) error {
	// 初始化 Telegram Bot
	return nil
}

func (s *TelegramBotService) Stop(ctx context.Context) error {
	// 停止 Bot 服务
	return nil
}

func (s *TelegramBotService) Send(ctx context.Context, response *domain.BotResponse) error {
	// 发送消息到 Telegram
	return nil
}
