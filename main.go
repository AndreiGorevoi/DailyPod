package main

import (
	"DailyPod/config"
	"DailyPod/service/bot"
	"DailyPod/service/dallas"
)

func main() {
	cfg := config.LoadConfig()
	dls := dallas.NewDallas(cfg)
	b := bot.NewBot(cfg, dls)
	b.Run()

}
