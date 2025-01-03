package main

import (
	"context"
	"log"

	"github.com/inggg-enddd/solana-bot/internal/bot/domain"

	"github.com/inggg-enddd/solana-bot/internal/bot/handlers"
	"github.com/inggg-enddd/solana-bot/internal/bot/registry"
	"github.com/inggg-enddd/solana-bot/internal/bot/service"
)

func main() {
	// 创建命令注册器
	cmdRegistry := registry.NewCommandRegistry()

	// 注册命令处理器
	cmdRegistry.Register(domain.StartCommand, handlers.NewStartHandler())
	// 注册其他处理器...

	// 创建 Bot 服务
	botService := service.NewTelegramBotService(cmdRegistry)

	// 启动服务
	ctx := context.Background()
	if err := botService.Start(ctx); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	// 等待信号处理...
}
