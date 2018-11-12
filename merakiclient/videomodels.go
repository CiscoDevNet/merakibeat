package merakiclient

import (
	"github.com/elastic/beats/libbeat/common"
)

type CameraCommon struct {
	CameraSerial string `json:"camera_serial"`
}

type ZoneHistoryInfoList []ZoneHistoryInfo

type ZoneHistoryInfo struct {
	CameraCommon
	Ts           string  `json:"ts"`
	Entrances    int     `json:"entrances"`
	AverageCount float32 `json:"average_count"`
}

type ZoneRecentInfoList []ZoneRecentInfo

type ZoneRecentInfo struct {
	CameraCommon
	ZoneID       int     `json:"zone_id"`
	SecondsAgo   int     `json:"seconds_ago"`
	Ts           float64 `json:"ts"`
	Entrances    int     `json:"entrances"`
	AverageCount float64 `json:"average_count"`
}

type CameraOverviewInfoList []CameraOverviewInfo

type CameraOverviewInfo struct {
	CameraCommon
	T0           float64 `json:"t0"`
	T1           float64 `json:"t1"`
	ZoneID       int     `json:"zone_id"`
	Entrances    int     `json:"entrances"`
	AverageCount float64 `json:"average_count"`
}

func (zi *ZoneHistoryInfo) GetMapStr(stattype string, addlnKVP map[string]string) (common.MapStr, error) {
	mapStr := common.MapStr{
		"type":          stattype,
		"cameraserial":  zi.CameraSerial,
		"entrances":     zi.Entrances,
		"average_count": zi.AverageCount,
		"timestamp":     zi.Ts,
	}
	for key, value := range addlnKVP {
		mapStr.Put(key, value)
	}
	return mapStr, nil
}
