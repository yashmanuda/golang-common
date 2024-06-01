package config

import (
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

func InitializeConfig() *koanf.Koanf {
	var kConfig = koanf.New("__")
	kConfig.Load(env.Provider("", "__", nil), nil)
	return kConfig
}
