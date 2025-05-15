package kafka

import (
	"time"

	"github.com/IBM/sarama"
)

type SyncProducer struct {
	t string
	p sarama.SyncProducer
}

func NewSyncProducer(topic string, cli sarama.Client) (*SyncProducer, error) {
	p, err := sarama.NewSyncProducerFromClient(cli)
	if err != nil {
		return nil, err
	}

	return &SyncProducer{
		t: topic,
		p: p,
	}, nil
}

func (sp *SyncProducer) Produce(value []byte) error {
	message := &sarama.ProducerMessage{
		Topic:     sp.t,
		Value:     sarama.ByteEncoder(value),
		Timestamp: time.Now(),
	}
	_, _, err := sp.p.SendMessage(message)

	return err
}
