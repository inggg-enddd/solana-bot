package domain

// Command 表示bot命令
type Command string

const (
	StartCommand    Command = "start"
	CreateCommand   Command = "create_wallet"
	ImportCommand   Command = "import_wallet"
	BalanceCommand  Command = "balance"
	TransferCommand Command = "transfer"
	HistoryCommand  Command = "history"
	HelpCommand     Command = "help"
)

// BotMessage 表示机器人消息领域模型
type BotMessage struct {
	ChatID    int64
	Text      string
	Command   Command
	Arguments []string
}

// BotResponse 表示机器人响应领域模型
type BotResponse struct {
	ChatID int64
	Text   string
	Markup any // Telegram键盘标记
}
