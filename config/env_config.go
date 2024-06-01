package config

import (
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"github.com/yashmanuda/golang-common/logger"
)

func New() *koanf.Koanf {
	var kConfig = koanf.New("__")
	kConfig.Load(env.Provider("", "__", nil), nil)
	logger.Logger.Info("Configs loaded")
	return kConfig
}
