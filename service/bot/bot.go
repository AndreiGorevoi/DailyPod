package bot

import (
	"DailyPod/config"
	"DailyPod/service/dallas"
	"DailyPod/service/formatter"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	cfg           *config.Config
	dallasService *dallas.Dallas
}

func NewBot(cfg *config.Config, dls *dallas.Dallas) *TelegramBot {
	return &TelegramBot{
		cfg:           cfg,
		dallasService: dls,
	}
}

func (b *TelegramBot) Run() {
	bot, err := tgbotapi.NewBotAPI(b.cfg.TelegramToken)

	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		switch update.Message.Text {
		case "/dls_next":
			var msg tgbotapi.Chattable
			games, err := b.dallasService.Next3Games()
			if err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Sth went wrong")
			}
			txt := formatter.FormatGamesToString(games)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, txt)
			bot.Send(msg)
		}
	}
}
