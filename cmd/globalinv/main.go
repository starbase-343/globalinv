package main

import (
	"context"
	"github.com/starbase-343/globalinv/internal/config"
	"github.com/starbase-343/globalinv/internal/infra/kafka"
	"os"
	"os/signal"
)

var gConf *config.App

func init() {
	appConf, err := config.Load()
	if err != nil {
		panic(err)
	}
	gConf = appConf
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	cli, err := kafka.Connect(gConf.MsgBroker)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	p, err := kafka.NewSyncProducer("test-topic", cli)
	if err != nil {
		panic(err)
	}

	c, err := kafka.NewConsumer("test-topic", cli)
	if err != nil {
		panic(err)
	}

	if err := p.Produce([]byte("Hello")); err != nil {
		panic(err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-c.Consume(ctx):
			if msg == nil {
				return
			}
			println(string(msg))
		}
	}
}
