package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"hotPotBot/internal/config"
	"hotPotBot/internal/context"
	"hotPotBot/internal/handlers"
	"hotPotBot/internal/logger"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	updates tgbotapi.UpdatesChannel
}

func NewBot(cfg *config.Config) *Bot {
	logger.Log.Info("Initializing bot")

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		logger.Log.Fatalf("Failed to create bot: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		logger.Log.Fatalf("Failed to get updates channel: %v", err)
	}

	return &Bot{
		bot:     bot,
		updates: updates,
	}
}

func (b *Bot) Start(ctx *context.AppContext) {
	for update := range b.updates {
		handlers.HandleUpdate(ctx, b.bot, update)
	}
}
