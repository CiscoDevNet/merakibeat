package merakiclient

import (
	"github.com/elastic/beats/libbeat/common"
)

type NetworkStat struct {
	Assoc   int `json:"assoc"`
	Auth    int `json:"auth"`
	Dhcp    int `json:"dhcp"`
	DNS     int `json:"dns"`
	Success int `json:"success"`
}
type DeviceNetworkStat struct {
	Serial          string      `json:"serial"`
	ConnectionStats NetworkStat `json:"connectionStats"`
}

type ClientNetworkStat struct {
	MAC             string      `json:"mac"`
	ConnectionStats NetworkStat `json:"connectionStats"`
}

type DevicesNetworkStat []DeviceNetworkStat
type ClientsNetworkStat []ClientNetworkStat

func (ns *NetworkStat) GetMapStr(stattype string, addlnKVP map[string]string) (common.MapStr, error) {
	mapStr := common.MapStr{
		"type":    stattype,
		"assoc":   ns.Assoc,
		"auth":    ns.Auth,
		"dhcp":    ns.Dhcp,
		"dns":     ns.DNS,
		"success": ns.Success,
	}
	for key, value := range addlnKVP {
		mapStr.Put(key, value)
	}
	return mapStr, nil
}

type LatencyRange struct {
	Num0    int `json:"0"`
	Num2    int `json:"2"`
	Num4    int `json:"4"`
	Num8    int `json:"8"`
	Num16   int `json:"16"`
	Num32   int `json:"32"`
	Num64   int `json:"64"`
	Num128  int `json:"128"`
	Num256  int `json:"256"`
	Num512  int `json:"512"`
	Num1024 int `json:"1024"`
	Num2048 int `json:"2048"`
}

type LatencyStats struct {
	BackgroundTraffic LatencyRange `json:"backgroundTraffic"`
	BestEffortTraffic LatencyRange `json:"bestEffortTraffic"`
	VideoTraffic      LatencyRange `json:"videoTraffic"`
	VoiceTraffic      LatencyRange `json:"voiceTraffic"`
}

type DeviceLatencyStat struct {
	Serial       string       `json:"serial"`
	LatencyStats LatencyStats `json:"latencyStats"`
}

type ClientLatencyStat struct {
	MAC          string       `json:"mac"`
	LatencyStats LatencyStats `json:"latencyStats"`
}

type DevicesLatencyStat []DeviceLatencyStat
type ClientsLatencyStat []ClientLatencyStat

func (lr *LatencyRange) GetAvgLat() float32 {
	den := (lr.Num0 + lr.Num2 + lr.Num4 + lr.Num8 + lr.Num16 + lr.Num32 + lr.Num64 +
		lr.Num128 + lr.Num256 + lr.Num512 + lr.Num1024 + lr.Num2048)

	if den == 0 {
		return 0.0
	}

	latency := float32(lr.Num0*(1/2)+lr.Num2*1+lr.Num4*2+lr.Num8*4+lr.Num16*8+lr.Num32*16+
		lr.Num64*32+lr.Num128*64+lr.Num256*128+lr.Num512*256+lr.Num1024*512+lr.Num2048*1024) /
		float32(den)

	return latency
}
func (ls *LatencyStats) GetMapStr(stattype string, addlnKVP map[string]string) (common.MapStr, error) {
	bglat := ls.BackgroundTraffic.GetAvgLat()
	belat := ls.BestEffortTraffic.GetAvgLat()
	vidlat := ls.VideoTraffic.GetAvgLat()
	vclat := ls.VoiceTraffic.GetAvgLat()
	latency := (bglat + belat + vidlat + vclat) / 4

	mapStr := common.MapStr{
		"type":                      stattype,
		"BackgroundTraffic.0":       ls.BackgroundTraffic.Num0,
		"BackgroundTraffic.2":       ls.BackgroundTraffic.Num2,
		"BackgroundTraffic.4":       ls.BackgroundTraffic.Num4,
		"BackgroundTraffic.8":       ls.BackgroundTraffic.Num8,
		"BackgroundTraffic.16":      ls.BackgroundTraffic.Num16,
		"BackgroundTraffic.32":      ls.BackgroundTraffic.Num32,
		"BackgroundTraffic.64":      ls.BackgroundTraffic.Num64,
		"BackgroundTraffic.128":     ls.BackgroundTraffic.Num128,
		"BackgroundTraffic.256":     ls.BackgroundTraffic.Num256,
		"BackgroundTraffic.512":     ls.BackgroundTraffic.Num512,
		"BackgroundTraffic.1024":    ls.BackgroundTraffic.Num1024,
		"BackgroundTraffic.2048":    ls.BackgroundTraffic.Num2048,
		"BackgroundTraffic.latency": bglat,

		"BestEffortTraffic.0":       ls.BestEffortTraffic.Num0,
		"BestEffortTraffic.2":       ls.BestEffortTraffic.Num2,
		"BestEffortTraffic.4":       ls.BestEffortTraffic.Num4,
		"BestEffortTraffic.8":       ls.BestEffortTraffic.Num8,
		"BestEffortTraffic.16":      ls.BestEffortTraffic.Num16,
		"BestEffortTraffic.32":      ls.BestEffortTraffic.Num32,
		"BestEffortTraffic.64":      ls.BestEffortTraffic.Num64,
		"BestEffortTraffic.128":     ls.BestEffortTraffic.Num128,
		"BestEffortTraffic.256":     ls.BestEffortTraffic.Num256,
		"BestEffortTraffic.512":     ls.BestEffortTraffic.Num512,
		"BestEffortTraffic.1024":    ls.BestEffortTraffic.Num1024,
		"BestEffortTraffic.2048":    ls.BestEffortTraffic.Num2048,
		"BestEffortTraffic.latency": belat,

		"VideoTraffic.0":       ls.VideoTraffic.Num0,
		"VideoTraffic.2":       ls.VideoTraffic.Num2,
		"VideoTraffic.4":       ls.VideoTraffic.Num4,
		"VideoTraffic.8":       ls.VideoTraffic.Num8,
		"VideoTraffic.16":      ls.VideoTraffic.Num16,
		"VideoTraffic.32":      ls.VideoTraffic.Num32,
		"VideoTraffic.64":      ls.VideoTraffic.Num64,
		"VideoTraffic.128":     ls.VideoTraffic.Num128,
		"VideoTraffic.256":     ls.VideoTraffic.Num256,
		"VideoTraffic.512":     ls.VideoTraffic.Num512,
		"VideoTraffic.1024":    ls.VideoTraffic.Num1024,
		"VideoTraffic.2048":    ls.VideoTraffic.Num2048,
		"VideoTraffic.latency": vidlat,

		"VoiceTraffic.0":       ls.VoiceTraffic.Num0,
		"VoiceTraffic.2":       ls.VoiceTraffic.Num2,
		"VoiceTraffic.4":       ls.VoiceTraffic.Num4,
		"VoiceTraffic.8":       ls.VoiceTraffic.Num8,
		"VoiceTraffic.16":      ls.VoiceTraffic.Num16,
		"VoiceTraffic.32":      ls.VoiceTraffic.Num32,
		"VoiceTraffic.64":      ls.VoiceTraffic.Num64,
		"VoiceTraffic.128":     ls.VoiceTraffic.Num128,
		"VoiceTraffic.256":     ls.VoiceTraffic.Num256,
		"VoiceTraffic.512":     ls.VoiceTraffic.Num512,
		"VoiceTraffic.1024":    ls.VoiceTraffic.Num1024,
		"VoiceTraffic.2048":    ls.VoiceTraffic.Num2048,
		"VoiceTraffic.latency": vclat,
		"Latency":              latency,
	}
	for key, value := range addlnKVP {
		mapStr.Put(key, value)
	}
	return mapStr, nil
}

type NetworkDetailList []NetworkDetails

type NetworkDetails struct {
	ID                 string `json:"id"`
	OrganizationID     string `json:"organizationId"`
	Name               string `json:"name"`
	TimeZone           string `json:"timeZone"`
	Tags               string `json:"tags"`
	Type               string `json:"type"`
	DisableMyMerakiCom bool   `json:"disableMyMerakiCom,omitempty"`
	ConfigTemplateID   string `json:"configTemplateId,omitempty"`
}

// Location Scan Data

type ScanData struct {
	Type    string     `json:"type"`
	Secret  string     `json:"secret"`
	Version string     `json:"version"`
	Data    ClientData `json:"data"`
}

type ClientData struct {
	ApMac        string        `json:"apMac"`
	ApFloors     []string      `json:"apFloors"`
	ApTags       []string      `json:"apTags"`
	Observations []Observation `json:"observations"`
}

type Observation struct {
	Ssid         string       `json:"ssid"`
	Ipv4         string       `json:"ipv4"`
	Ipv6         string       `json:"ipv6"`
	SeenEpoch    float64      `json:"seenEpoch"`
	SeenTime     string       `json:"seenTime"`
	Rssi         int          `json:"rssi"`
	Manufacturer string       `json:"manufacturer"`
	Os           string       `json:"os"`
	Location     LocationData `json:"location"`
	ClientMac    string       `json:"clientMac"`
}

type LocationData struct {
	Lat float64   `json:"lat"`
	X   []float64 `json:"x"`
	Lng float64   `json:"lng"`
	Unc float64   `json:"unc"`
	Y   []float64 `json:"y"`
}

type ElasticLoc struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (sd *ScanData) GetMapStr(stattype string, addlnKVP map[string]string) ([]common.MapStr, error) {

	var mapStrArr []common.MapStr
	for _, observation := range sd.Data.Observations {
		elLoc := ElasticLoc{
			Lat: observation.Location.Lat,
			Lon: observation.Location.Lng,
		}
		mapStr := common.MapStr{
			"type":                stattype,
			"datatype":            sd.Type,
			"apMac":               sd.Data.ApMac,
			"apFloors":            sd.Data.ApFloors,
			"apTags":              sd.Data.ApTags,
			"client.ssid":         observation.Ssid,
			"client.rssi":         observation.Rssi,
			"cliet.ipv4":          observation.Ipv4,
			"client.ipv6":         observation.Ipv6,
			"client.manufacturer": observation.Manufacturer,
			"client.seenTime":     observation.SeenTime,
			"client.seenEpoch":    observation.SeenEpoch,
			"client.os":           observation.Os,
			"client.Mac":          observation.ClientMac,
			"client.lat":          observation.Location.Lat,
			"client.lng":          observation.Location.Lng,
			"client.unc":          observation.Location.Unc,
			"client.x":            observation.Location.X,
			"client.y":            observation.Location.Y,
			"location":            elLoc,
		}
		for key, value := range addlnKVP {
			mapStr.Put(key, value)
		}
		mapStrArr = append(mapStrArr, mapStr)
	}
	return mapStrArr, nil
}
