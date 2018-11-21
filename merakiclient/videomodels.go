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
	StartTs      string  `json:"startTs"`
	EndTs        string  `json:"endTs"`
	Entrances    int     `json:"entrances"`
	AverageCount float32 `json:"averageCount"`
}

type ZoneRecentInfoList []ZoneRecentInfo

type ZoneRecentInfo struct {
	CameraCommon
	ZoneID       int     `json:"zone_id"`
	StartTs      string  `json:"startTs"`
	EndTs        string  `json:"endTs"`
	Entrances    int     `json:"entrances"`
	AverageCount float32 `json:"averageCount"`
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
		"timestamp":     zi.StartTs,
	}
	for key, value := range addlnKVP {
		mapStr.Put(key, value)
	}
	return mapStr, nil
}

func (zi *ZoneRecentInfo) GetMapStr(stattype string, addlnKVP map[string]string) (common.MapStr, error) {
	mapStr := common.MapStr{
		"type":          stattype,
		"cameraserial":  zi.CameraSerial,
		"entrances":     zi.Entrances,
		"average_count": zi.AverageCount,
		"timestamp":     zi.StartTs,
		"zone_id":       zi.ZoneID,
	}
	for key, value := range addlnKVP {
		mapStr.Put(key, value)
	}
	return mapStr, nil
}
