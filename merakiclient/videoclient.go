package merakiclient

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

func (mc *MerakiClient) getVideoHistoryStartTime() int64 {
	var lag time.Duration = 150 * time.Second
	startTime := time.Now().Add(0 - (lag + mc.VideoPeriod)).Unix()
	return startTime
}

func (mc *MerakiClient) GetZoneHistory(cameraSerial string, zoneID string) (mapStrArr []common.MapStr, err error) {

	netURL := fmt.Sprintf("%s/api/v0/devices/%s/camera/analytics/zones/%s/history", mc.URL, cameraSerial, zoneID)
	params := map[string]string{"startingAfter": fmt.Sprintf("%d", mc.getVideoHistoryStartTime())}

	body, err := mc.getDataQueryParam(netURL, params)
	if err != nil {
		logp.Info("Failed to get Network List from Meraki API %s", err.Error())
		return nil, err
	}

	var zoneList ZoneHistoryInfoList
	err = json.Unmarshal(body, &zoneList)
	if err != nil {
		logp.Info("Failed to Unmarshal data from API %s", err.Error())
		return nil, err
	}

	additionalMap := map[string]string{
		"zone_id": zoneID,
	}
	for _, zone := range zoneList {
		zone.CameraSerial = cameraSerial
		mapStr, err := zone.GetMapStr("CameraHistoryZoneInfo", additionalMap)
		if err == nil {
			mapStrArr = append(mapStrArr, mapStr)
		} else {
			logp.Info("Failed to get mapStr from Zone info %s", err.Error())
		}
	}
	return mapStrArr, err
}

func (mc *MerakiClient) GetZoneRecentInfo(cameraSerial string, zoneID string) (mapStrArr []common.MapStr, err error) {

	netURL := fmt.Sprintf("%s/api/v0/devices/%s/camera/analytics/zones/%s/recent", mc.URL, cameraSerial, zoneID)
	params := map[string]string{}

	body, err := mc.getDataQueryParam(netURL, params)
	if err != nil {
		logp.Info("Failed to get Network List from Meraki API %s", err.Error())
		return nil, err
	}

	var zoneRecentList ZoneRecentInfoList
	err = json.Unmarshal(body, &zoneRecentList)
	if err != nil {
		logp.Info("Failed to Unmarshal data from API %s", err.Error())
		return nil, err
	}

	additionalMap := map[string]string{}
	for _, zone := range zoneRecentList {
		zone.CameraSerial = cameraSerial
		mapStr, err := zone.GetMapStr("CameraHistoryZoneInfo", additionalMap)
		if err == nil {
			mapStrArr = append(mapStrArr, mapStr)
		} else {
			logp.Info("Failed to get mapStr from Zone info %s", err.Error())
		}
	}
	return mapStrArr, err
}
