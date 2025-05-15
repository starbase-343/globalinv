package config_test

import (
	"encoding/json"
	"github.com/starbase-343/globalinv/internal/config"
	"github.com/starbase-343/globalinv/internal/config/profile"
	"os"
	"path/filepath"
	"testing"
)

func TestNewDevConfig(t *testing.T) {
	d := t.TempDir()

	wantDB := config.DB{
		Driver:       "foo",
		Name:         "devdb",
		Host:         "devhost",
		Port:         "1111",
		User:         "devuser",
		Password:     "devpass",
		MigrationDir: "devmigrations",
	}
	wantApp := config.App{
		Profile: profile.Dev,
		DB:      wantDB,
	}
	b, err := json.Marshal(wantApp)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	if err := os.WriteFile(filepath.Join(d, "dev.json"), b, 0644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	os.Unsetenv(string(config.ProfileKey))
	os.Setenv(string(config.PathKey), d)

	// Call Load()
	conf, err := config.Load()
	if err != nil {
		t.Fatalf("config.Load() error: %v", err)
	}

	// Verify profile
	if conf.Profile != profile.Dev {
		t.Errorf("Profile = %v; want %v", conf.Profile, profile.Dev)
	}

	// Verify DB
	if conf.DB != wantDB {
		t.Errorf("DB = %+v; want %+v", conf.DB, wantDB)
	}
}

func TestNewProdConfig_DefaultsAndEnv(t *testing.T) {
	d := t.TempDir()

	if err := os.WriteFile(filepath.Join(d, "prod.json"), []byte("{}"), 0644); err != nil {
		t.Fatalf("WriteFile prod.json: %v", err)
	}

	os.Setenv(string(config.ProfileKey), string(profile.Prod))
	os.Setenv(string(config.PathKey), d)

	env := map[string]string{
		string(config.PostgresDBKey):   "proddb",
		string(config.PostgresHostKey): "prodhost",
		string(config.PostgresUserKey): "produser",
		string(config.PostgresPassKey): "prodpass",
	}
	for k, v := range env {
		os.Setenv(k, v)
	}

	os.Unsetenv(string(config.PostgresDriverKey))
	os.Unsetenv(string(config.PostgresPortKey))
	os.Unsetenv(string(config.PostgresMigrationKey))

	// Call Load()
	conf, err := config.Load()
	if err != nil {
		t.Fatalf("config.Load() error: %v", err)
	}

	// Verify profile
	if conf.Profile != profile.Prod {
		t.Errorf("Profile = %v; want %v", conf.Profile, profile.Prod)
	}

	// Verify DB defaults and env
	db := conf.DB
	if db.Driver != "postgres" {
		t.Errorf("Driver = %v; want postgres", db.Driver)
	}
	if db.Name != env[string(config.PostgresDBKey)] {
		t.Errorf("Name = %v; want %v", db.Name, env[string(config.PostgresDBKey)])
	}
	if db.Host != env[string(config.PostgresHostKey)] {
		t.Errorf("Host = %v; want %v", db.Host, env[string(config.PostgresHostKey)])
	}
	if db.Port != "5432" {
		t.Errorf("Port = %v; want 5432", db.Port)
	}
	if db.User != env[string(config.PostgresUserKey)] {
		t.Errorf("User = %v; want %v", db.User, env[string(config.PostgresUserKey)])
	}
	if db.Password != env[string(config.PostgresPassKey)] {
		t.Errorf("Password = %v; want %v", db.Password, env[string(config.PostgresPassKey)])
	}
	if db.MigrationDir != "./migrations" {
		t.Errorf("MigrationDir = %v; want ./migrations", db.MigrationDir)
	}
}
