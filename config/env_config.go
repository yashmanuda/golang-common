package config

import (
	"sync"

	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

// singleton instance
var instance *koanf.Koanf
var once sync.Once

func GetConfigInstance() *koanf.Koanf {
	once.Do(
		func() {
			instance = koanf.New("__")
			instance.Load(env.Provider("", "__", nil), nil)
		})
	return instance
}
