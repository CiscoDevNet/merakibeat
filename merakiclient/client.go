package merakiclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

type MerakiClient struct {
	URL        string
	Key        string
	OrgID      string
	NetworkIDs []string
}

func NewMerakiClient(url, key, orgID string, networkIDs []string) MerakiClient {
	return MerakiClient{
		URL:        url,
		Key:        key,
		OrgID:      orgID,
		NetworkIDs: networkIDs,
	}
}

func (mc *MerakiClient) GetNetworkConnectionStat(networkID string) (common.MapStr, error) {
	netURL := fmt.Sprintf("%s/api/v0/networks/%s/connectionStats", mc.URL, networkID)

	resp, err := http.Get(netURL)
	if err != nil {
		logp.Info("Failed to connect Meraki API %s", err.Error())
		return common.MapStr{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logp.Info("Failed to get data from Meraki API %s", err.Error())
		return common.MapStr{}, err
	}
	var data NetworkStat
	err = json.Unmarshal(body, &data)
	if err != nil {
		logp.Info("Failed to Unmarshal data from API %s", err.Error())
		return common.MapStr{}, err
	}
	additionalMap := map[string]string{
		"networkid": networkID,
	}
	return data.GetMapStr("NetworkConnectionStat", additionalMap)

}

func (mc *MerakiClient) GetNetworkLatencyStat(networkID string) (common.MapStr, error) {
	netURL := fmt.Sprintf("%s/api/v0/networks/%s/latencyStats", mc.URL, networkID)

	resp, err := http.Get(netURL)
	if err != nil {
		logp.Info("Failed to connect Meraki API %s", err.Error())
		return common.MapStr{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logp.Info("Failed to get data from Meraki API %s", err.Error())
		return common.MapStr{}, err
	}
	var data LatencyStats
	err = json.Unmarshal(body, &data)
	if err != nil {
		logp.Info("Failed to Unmarshal data from API %s", err.Error())
		return common.MapStr{}, err
	}
	additionalMap := map[string]string{
		"networkid": networkID,
	}
	return data.GetMapStr("NetworkLatencyStat", additionalMap)
}

func (mc *MerakiClient) GetDevicesConnectionStat(networkID string) (devicesStat []common.MapStr, err error) {
	netURL := fmt.Sprintf("%s/api/v0/networks/%s/devices/connectionStats", mc.URL, networkID)
	logp.Info("URL %s", netURL)
	resp, err := http.Get(netURL)
	if err != nil {
		logp.Info("Failed to connect Meraki API %s", err.Error())
		return []common.MapStr{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logp.Info("Failed to get data from Meraki API %s", err.Error())
		return []common.MapStr{}, err
	}
	var data DevicesNetworkStat
	err = json.Unmarshal(body, &data)
	if err != nil {
		logp.Info("Failed to Unmarshal data from API %s", err.Error())
		return []common.MapStr{}, err
	}
	logp.Info("data from API %+v", data)

	for key, value := range data {
		additionalMap := map[string]string{
			"networkid": networkID,
			"deviceid":  key,
		}
		nwStat, _ := value.GetMapStr("DeviceNetworkConnectionStat", additionalMap)
		devicesStat = append(devicesStat, nwStat)
	}
	logp.Info("%+v", devicesStat)
	return devicesStat, err

}

func (mc *MerakiClient) GetDevicesLatencyStat(networkID string) (devicesStat []common.MapStr, err error) {
	netURL := fmt.Sprintf("%s/api/v0/networks/%s/devices/latencyStats", mc.URL, networkID)
	logp.Info("URL %s", netURL)
	resp, err := http.Get(netURL)
	if err != nil {
		logp.Info("Failed to connect Meraki API %s", err.Error())
		return []common.MapStr{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logp.Info("Failed to get data from Meraki API %s", err.Error())
		return []common.MapStr{}, err
	}
	var data DevicesLatencyStat
	err = json.Unmarshal(body, &data)
	if err != nil {
		logp.Info("Failed to Unmarshal data from API %s", err.Error())
		return []common.MapStr{}, err
	}
	logp.Info("data from API %+v", data)

	for key, value := range data {
		additionalMap := map[string]string{
			"networkid": networkID,
			"deviceid":  key,
		}
		latencyStat, _ := value.GetMapStr("DeviceNetworkLatencyStat", additionalMap)
		devicesStat = append(devicesStat, latencyStat)
	}
	logp.Info("%+v", devicesStat)
	return devicesStat, err

}

func (mc *MerakiClient) GetClientConnectionStat(networkID string) (clientsStat []common.MapStr, err error) {
	netURL := fmt.Sprintf("%s/api/v0/networks/%s/clients/connectionStats", mc.URL, networkID)
	logp.Info("URL %s", netURL)
	resp, err := http.Get(netURL)
	if err != nil {
		logp.Info("Failed to connect Meraki API %s", err.Error())
		return []common.MapStr{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logp.Info("Failed to get data from Meraki API %s", err.Error())
		return []common.MapStr{}, err
	}
	var data ClientsNetworkStat
	err = json.Unmarshal(body, &data)
	if err != nil {
		logp.Info("Failed to Unmarshal data from API %s", err.Error())
		return []common.MapStr{}, err
	}
	logp.Info("data from API %+v", data)

	for key, value := range data {
		additionalMap := map[string]string{
			"networkid": networkID,
			"clientid":  key,
		}
		nwStat, _ := value.GetMapStr("ClientNetworkConnectionStat", additionalMap)
		clientsStat = append(clientsStat, nwStat)
	}
	logp.Info("%+v", clientsStat)
	return clientsStat, err

}

func (mc *MerakiClient) GetClientLatencyStat(networkID string) (clientsStat []common.MapStr, err error) {
	netURL := fmt.Sprintf("%s/api/v0/networks/%s/devices/latencyStats", mc.URL, networkID)
	logp.Info("URL %s", netURL)
	resp, err := http.Get(netURL)
	if err != nil {
		logp.Info("Failed to connect Meraki API %s", err.Error())
		return []common.MapStr{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logp.Info("Failed to get data from Meraki API %s", err.Error())
		return []common.MapStr{}, err
	}
	var data ClientsLatencyStat
	err = json.Unmarshal(body, &data)
	if err != nil {
		logp.Info("Failed to Unmarshal data from API %s", err.Error())
		return []common.MapStr{}, err
	}
	logp.Info("data from API %+v", data)

	for key, value := range data {
		additionalMap := map[string]string{
			"networkid": networkID,
			"deviceid":  key,
		}
		latencyStat, _ := value.GetMapStr("ClientNetworkLatencyStat", additionalMap)
		clientsStat = append(clientsStat, latencyStat)
	}
	logp.Info("%+v", clientsStat)
	return clientsStat, err

}
