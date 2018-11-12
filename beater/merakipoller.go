package beater

import (
	"time"

	"github.com/CiscoDevNet/merakibeat/config"
	"github.com/CiscoDevNet/merakibeat/merakiclient"
	_ "github.com/elastic/beats/libbeat/common"
)

type MerakiPoller struct {
	merakibeat *Merakibeat
	config     config.Config
	timeout    time.Duration
	mc         merakiclient.MerakiClient
}

type MerakiPolleriIntf interface {
	Run()
}
