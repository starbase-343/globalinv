package config

import (
	"github.com/starbase-343/globalinv/internal/config/profile"
	"log/slog"
)

type App struct {
	Profile   profile.Profile `json:"profile"`
	LogLevel  slog.Level      `json:"logLevel"`
	DB        DB              `json:"db"`
	MsgBroker MsgBroker       `json:"msgBroker"`
}
