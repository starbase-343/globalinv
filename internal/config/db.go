package config

type DB struct {
	Driver       string `json:"driver"`
	Name         string `json:"name"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	MigrationDir string `json:"migrationDir"`
}
