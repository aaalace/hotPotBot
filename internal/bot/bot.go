package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hotPotBot/internal/config"
	"hotPotBot/internal/context"
	"hotPotBot/internal/handlers"
	"hotPotBot/internal/logger"
	"log"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	updates tgbotapi.UpdatesChannel
}

func NewBot(cfg *config.Config) *Bot {
	logger.Log.Info("Initializing bot")

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		logger.Log.Fatalf("Failed to create bot | %v", err.Error())
		return nil
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	return &Bot{
		bot:     bot,
		updates: updates,
	}
}

func (b *Bot) Start(ctx *context.AppContext) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Panic occurred | %v", r)
				}
			}()

			for update := range b.updates {
				handlers.HandleUpdate(ctx, b.bot, update)
			}
		}()
	}
}
