package config

import (
	"encoding/json"
	"fmt"
	"github.com/starbase-343/globalinv/internal/config/env"
	"github.com/starbase-343/globalinv/internal/config/profile"
	"os"
	"path"
)

const (
	PostgresDriverKey    env.Key = "POSTGRES_DRIVER"
	PostgresHostKey      env.Key = "POSTGRES_HOST"
	PostgresPortKey      env.Key = "POSTGRES_PORT"
	PostgresDBKey        env.Key = "POSTGRES_DB"
	PostgresUserKey      env.Key = "POSTGRES_USER"
	PostgresPassKey      env.Key = "POSTGRES_PASS"
	PostgresMigrationKey env.Key = "POSTGRES_MIGRATION"
)

const (
	ProfileKey env.Key = "PROFILE"
	PathKey    env.Key = "CONFIG_PATH"
)

func Load() (*App, error) {
	var conf *App

	configDir := env.OrDefault(PathKey, "config/")
	p := env.OrDefault(ProfileKey, "dev")

	parsedProfile, err := profile.Parse(p)
	if err != nil {
		return nil, err
	}

	if parsedProfile == profile.Dev {
		conf, err = newDevConfig(configDir)
	} else if parsedProfile == profile.Prod {
		conf, err = newProdConfig(configDir)
	}

	if err != nil {
		return nil, err
	}

	return conf, nil
}

func newDevConfig(dir string) (*App, error) {
	conf := &App{
		Profile: profile.Dev,
	}

	configPath := path.Join(dir, fmt.Sprintf("%s.json", profile.Dev))
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, conf); err != nil {
		return nil, err
	}

	return conf, nil
}

func newProdConfig(dir string) (*App, error) {
	conf := &App{
		Profile: profile.Prod,
		DB: DB{
			Driver:       env.OrDefault(PostgresDriverKey, "postgres"),
			Name:         env.Must(PostgresDBKey),
			Host:         env.Must(PostgresHostKey),
			Port:         env.OrDefault(PostgresPortKey, "5432"),
			User:         env.Must(PostgresUserKey),
			Password:     env.Must(PostgresPassKey),
			MigrationDir: env.OrDefault(PostgresMigrationKey, "./migrations"),
		},
	}

	configPath := path.Join(dir, fmt.Sprintf("%s.json", profile.Prod))
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, conf); err != nil {
		return nil, err
	}

	return conf, nil
}
