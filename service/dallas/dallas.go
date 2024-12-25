package dallas

import (
	"DailyPod/config"
)

type Dallas struct {
	config *config.Config
}

func NewDallas(cfg *config.Config) *Dallas {
	return &Dallas{
		config: cfg,
	}
}
