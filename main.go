package main

import (
	"DailyPod/config"
	"DailyPod/service/dallas"
)

func main() {
	cfg := config.LoadConfig()
	_ = dallas.NewDallas(cfg)

}
