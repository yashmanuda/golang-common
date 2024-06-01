package config

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

type DBConfig struct {
	Host string `koanf:"HOST"`
	Port string `koanf:"PORT"`
}

func TestEnvConfig(t *testing.T) {
	godotenv.Load("./test.env")
	assert := assert.New(t)
	var kConfig = New()

	assert.Equal("Golang Common", kConfig.MustString("APPLICATION_NAME"), "Application name doesn't match")

	var dbConfig DBConfig
	kConfig.Unmarshal("DATABASE__PRIMARY", &dbConfig)
	assert.Equal(dbConfig.Host, "localhost", "Database host doesn't match")
	assert.Equal(dbConfig.Port, "5362", "Database port doesn't match")
}
