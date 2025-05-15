package config

type MsgBroker struct {
	Region  string   `json:"region"`
	Brokers []string `json:"brokers"`
	SSL     bool     `json:"ssl"`
}
