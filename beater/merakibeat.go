package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/CiscoDevNet/merakibeat/config"
)

type Merakibeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
	b      *beat.Beat
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Merakibeat{
		done:   make(chan struct{}),
		config: c,
		b:      b,
	}
	return bt, nil
}

func (bt *Merakibeat) Run(b *beat.Beat) error {
	logp.Info("merakibeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	poller := NewMerakiPoller(bt, bt.config)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}
		poller.Run()
	}
}

func (bt *Merakibeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
