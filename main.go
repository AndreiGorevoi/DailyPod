package main

import (
	"DailyPod/config"
	"DailyPod/service/bot"
	"DailyPod/service/dallas"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	dls := dallas.NewDallas(cfg, &http.Client{})
	b := bot.NewBot(cfg, dls)
	b.Run()

}
