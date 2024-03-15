package config

import (
	"back-end/internal/db"
	"back-end/internal/server"
	"encoding/json"
	"os"
)

type Config struct {
	Db  db.Database   `json:"database"`
	Srv server.Server `json:"server"`
}

func GetConfigs() (*Config, error) {
	bytes, err := os.ReadFile("./internal/configs.json")
	if err != nil {
		return nil, err
	}
	var conf Config
	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, err
}
