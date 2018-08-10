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

type DevicesNetworkStat map[string]NetworkStat
type ClientsNetworkStat map[string]NetworkStat

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

type DevicesLatencyStat map[string]LatencyStats
type ClientsLatencyStat map[string]LatencyStats

func (ls *LatencyStats) GetMapStr(stattype string, addlnKVP map[string]string) (common.MapStr, error) {
	mapStr := common.MapStr{
		"type":                    stattype,
		"BackgroundTraffic.0":     ls.BackgroundTraffic.Num0,
		"BackgroundTraffic.2":     ls.BackgroundTraffic.Num2,
		"BackgroundTraffic.4":     ls.BackgroundTraffic.Num4,
		"BackgroundTraffic.8":     ls.BackgroundTraffic.Num8,
		"BackgroundTraffic.16":    ls.BackgroundTraffic.Num16,
		"BackgroundTraffic.32":    ls.BackgroundTraffic.Num32,
		"BackgroundTraffic.64":    ls.BackgroundTraffic.Num64,
		"BackgroundTraffic.128":   ls.BackgroundTraffic.Num128,
		"BackgroundTraffic.256":   ls.BackgroundTraffic.Num256,
		"BackgroundTraffic.512":   ls.BackgroundTraffic.Num512,
		"BackgroundTraffic.1024":  ls.BackgroundTraffic.Num1024,
		"BackgroundTraffic.20148": ls.BackgroundTraffic.Num2048,

		"BestEffortTraffic.0":     ls.BestEffortTraffic.Num0,
		"BestEffortTraffic.2":     ls.BestEffortTraffic.Num2,
		"BestEffortTraffic.4":     ls.BestEffortTraffic.Num4,
		"BestEffortTraffic.8":     ls.BestEffortTraffic.Num8,
		"BestEffortTraffic.16":    ls.BestEffortTraffic.Num16,
		"BestEffortTraffic.32":    ls.BestEffortTraffic.Num32,
		"BestEffortTraffic.64":    ls.BestEffortTraffic.Num64,
		"BestEffortTraffic.128":   ls.BestEffortTraffic.Num128,
		"BestEffortTraffic.256":   ls.BestEffortTraffic.Num256,
		"BestEffortTraffic.512":   ls.BestEffortTraffic.Num512,
		"BestEffortTraffic.1024":  ls.BestEffortTraffic.Num1024,
		"BestEffortTraffic.20148": ls.BestEffortTraffic.Num2048,

		"VideoTraffic.0":     ls.VideoTraffic.Num0,
		"VideoTraffic.2":     ls.VideoTraffic.Num2,
		"VideoTraffic.4":     ls.VideoTraffic.Num4,
		"VideoTraffic.8":     ls.VideoTraffic.Num8,
		"VideoTraffic.16":    ls.VideoTraffic.Num16,
		"VideoTraffic.32":    ls.VideoTraffic.Num32,
		"VideoTraffic.64":    ls.VideoTraffic.Num64,
		"VideoTraffic.128":   ls.VideoTraffic.Num128,
		"VideoTraffic.256":   ls.VideoTraffic.Num256,
		"VideoTraffic.512":   ls.VideoTraffic.Num512,
		"VideoTraffic.1024":  ls.VideoTraffic.Num1024,
		"VideoTraffic.20148": ls.VideoTraffic.Num2048,

		"VoiceTraffic.0":     ls.VoiceTraffic.Num0,
		"VoiceTraffic.2":     ls.VoiceTraffic.Num2,
		"VoiceTraffic.4":     ls.VoiceTraffic.Num4,
		"VoiceTraffic.8":     ls.VoiceTraffic.Num8,
		"VoiceTraffic.16":    ls.VoiceTraffic.Num16,
		"VoiceTraffic.32":    ls.VoiceTraffic.Num32,
		"VoiceTraffic.64":    ls.VoiceTraffic.Num64,
		"VoiceTraffic.128":   ls.VoiceTraffic.Num128,
		"VoiceTraffic.256":   ls.VoiceTraffic.Num256,
		"VoiceTraffic.512":   ls.VoiceTraffic.Num512,
		"VoiceTraffic.1024":  ls.VoiceTraffic.Num1024,
		"VoiceTraffic.20148": ls.VoiceTraffic.Num2048,
	}
	for key, value := range addlnKVP {
		mapStr.Put(key, value)
	}
	return mapStr, nil
}
