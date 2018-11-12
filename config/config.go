// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period            time.Duration `config:"period"`
	MerakiHost        string        `config:"merakihost"`
	MerakiKey         string        `config:"merakikey"`
	MerakiNetworkIDs  []string      `config:"merakinewtorkids"`
	MerakiOrgID       string        `config:"merakiorgid"`
	NwConnStat        int           `config:"nwconnstat"`
	NwLatencyStat     int           `config:"nwlatencystat"`
	DeviceConnStat    int           `config:"devconnstat"`
	DeviceLatencyStat int           `config:"devlatencystat"`
	ClientConnStat    int           `config:"clconnstat"`
	ClientLatencyStat int           `config:"cllatencystat"`
	ScanSecret        string        `config:"scanSecret"`
	ScanValidator     string        `config:"scanValidator"`
	ScanEnable        int           `config:"scanEnable"`
	VideoPeriod       time.Duration `config:"videoPeriod"`
	CameraZoneList    []string      `config:"cameraZoneList"`

	MerakiNetworkIDsAll map[string]string
}

var DefaultConfig = Config{
	Period:     1 * time.Minute,
	MerakiHost: "http://locahost:5050",
}
