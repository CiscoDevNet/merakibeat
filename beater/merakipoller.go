package beater

import (
	"time"

	"github.com/elastic/beats/libbeat/beat"
	_ "github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/npateriya/merakibeat/config"
	"github.com/npateriya/merakibeat/merakiclient"
)

type MerakiPoller struct {
	merakibeat *Merakibeat
	config     config.Config
	timeout    time.Duration
	mc         merakiclient.MerakiClient
}

func NewMerakiPoller(merakibeat *Merakibeat, config config.Config) *MerakiPoller {
	mc := merakiclient.NewMerakiClient(config.MerakiHost, config.MerakiKey,
		config.MerakiOrgID, config.MerakiNetworkIDs)

	poller := &MerakiPoller{
		merakibeat: merakibeat,
		config:     config,
		mc:         mc,
	}
	return poller
}

// This is function that will call MerakiClient to fetch & publish data based on
// config item.  MerakiClient should have no understanding of beats framework except
// function that returns mapstr type.
func (p *MerakiPoller) Run() {

	// Publish Network Event
	for _, netID := range p.config.MerakiNetworkIDs {
		mapStr, err := p.mc.GetNetworkConnectionStat(netID)
		if err == nil {
			event := beat.Event{
				Timestamp: time.Now(),
				Fields:    mapStr,
			}
			p.merakibeat.client.Publish(event)
			logp.Info("Network Connection Stat event sent")
		}
	}
}
