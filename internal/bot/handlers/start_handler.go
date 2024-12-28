package handlers

import (
	"context"

	"github.com/your/project/internal/bot/domain"
)

type StartHandler struct {
	// 依赖注入
}

func NewStartHandler() *StartHandler {
	return &StartHandler{}
}

func (h *StartHandler) Handle(ctx context.Context, msg *domain.BotMessage) (*domain.BotResponse, error) {
	return &domain.BotResponse{
		ChatID: msg.ChatID,
		Text: "欢迎使用 Solana 交易执行系统!\n" +
			"请使用以下命令：\n" +
			"/create_wallet - 创建新钱包\n" +
			"/import_wallet - 导入已有钱包\n" +
			"/balance - 查询余额\n" +
			"/transfer - 发起转账\n" +
			"/history - 查询历史\n" +
			"/help - 帮助信息",
	}, nil
}
