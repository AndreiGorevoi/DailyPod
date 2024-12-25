package dallas

import (
	"DailyPod/config"
	"fmt"
)

type Dallas struct {
	config *config.Config
}

func NewDallas(cfg *config.Config) *Dallas {
	fmt.Println(cfg.API_NBA_key)
	return &Dallas{
		config: cfg,
	}
}
