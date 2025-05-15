package kafka

import (
	"context"
	"github.com/IBM/sarama"
)

const msgChanSize = 1000

type Consumer struct {
	t       string
	pc      sarama.PartitionConsumer
	msgChan chan []byte
}

func NewConsumer(topic string, cli sarama.Client) (*Consumer, error) {
	c, err := sarama.NewConsumerFromClient(cli)
	if err != nil {
		return nil, err
	}

	pc, err := c.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		t:       topic,
		pc:      pc,
		msgChan: make(chan []byte, msgChanSize),
	}, nil
}

func (c *Consumer) Consume(ctx context.Context) <-chan []byte {
	go c.listenWithContext(ctx)
	return c.msgChan
}

func (c *Consumer) listenWithContext(ctx context.Context) {
	defer close(c.msgChan)

	for {
		select {
		case msg, ok := <-c.pc.Messages():
			if !ok {
				return
			}
			c.msgChan <- msg.Value
		case <-ctx.Done():
			return
		}
	}
}
