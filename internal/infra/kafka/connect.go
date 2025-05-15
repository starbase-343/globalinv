package kafka

import (
	"crypto/tls"
	"github.com/IBM/sarama"
	"github.com/starbase-343/globalinv/internal/config"
	"log/slog"
)

func Connect(conf config.MsgBroker) (sarama.Client, error) {
	saramaConf := sarama.NewConfig()
	saramaConf.Net.SASL.Enable = conf.SSL
	saramaConf.Net.SASL.Mechanism = sarama.SASLTypeOAuth

	tlsConfig := tls.Config{
		InsecureSkipVerify: true,
	}
	saramaConf.Net.TLS.Enable = conf.SSL
	saramaConf.Net.TLS.Config = &tlsConfig
	saramaConf.Version = sarama.V3_6_0_0
	saramaConf.Producer.Return.Successes = true
	saramaConf.Producer.Return.Errors = true

	cli, err := sarama.NewClient(conf.Brokers, saramaConf)
	if err != nil {
		return nil, err
	}
	
	slog.Info("Successfully connected to kafka", "hosts", conf.Brokers, "region", conf.Region)
	return cli, nil
}
