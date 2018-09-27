package merakiclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

type MerakiClient struct {
	URL        string
	Key        string
	OrgID      string
	NetworkIDs []string
	Period     time.Duration
}

func NewMerakiClient(url, key, orgID string, networkIDs []string, period time.Duration) MerakiClient {
	return MerakiClient{
		URL:        url,
		Key:        key,
		OrgID:      orgID,
		NetworkIDs: networkIDs,
		Period:     period,
	}
}

func (mc *MerakiClient) getData(netURL string) ([]byte, error) {
	client := &http.Client{}
	var lag time.Duration = 300 * time.Second
	req, err := http.NewRequest("GET", netURL, nil)
	if err != nil {
		logp.Info("Failed to connect Meraki API %s", err.Error())
		return nil, err
	}
	req.Header.Add("X-Cisco-Meraki-API-Key", mc.Key)

	q := req.URL.Query()
	endTime := time.Now().Add(0 - lag).Unix()
	startTime := time.Now().Add(0 - (lag + mc.Period)).Unix()

	//startTime = startTime - 600

	q.Add("t0", strconv.FormatInt(startTime, 10))
	q.Add("t1", strconv.FormatInt(endTime, 10))
	req.URL.RawQuery = q.Encode()

	logp.Info("Calling API URL %+v", req.URL)
	resp, err := client.Do(req)
	if err != nil {
		logp.Info("Failed to connect Meraki API %s", err.Error())
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	logp.Info("%s", string(body[:]))
	return body, err
}

func (mc *MerakiClient) GetNetworksForOrg() (NetworkDetailList, error) {

	netURL := fmt.Sprintf("%s/api/v0/organizations/%s/networks", mc.URL, mc.OrgID)

	body, err := mc.getData(netURL)
	if err != nil {
		logp.Info("Failed to get Network List from Meraki API %s", err.Error())
		return nil, err
	}

	var data NetworkDetailList
	err = json.Unmarshal(body, &data)
	if err != nil {
		logp.Info("Failed to Unmarshal data from API %s", err.Error())
		return nil, err
	}
	var wirelessnws NetworkDetailList
	for _, nw := range data {
		if nw.Type == "wireless" {
			wirelessnws = append(wirelessnws, nw)
		}
	}
	return wirelessnws, err
}

func (mc *MerakiClient) GetNetworkConnectionStat(networkID string) (common.MapStr, error) {
	netURL := fmt.Sprintf("%s/api/v0/networks/%s/connectionStats", mc.URL, networkID)

	body, err := mc.getData(netURL)
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

	body, err := mc.getData(netURL)
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

	body, err := mc.getData(netURL)
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

	for _, value := range data {
		additionalMap := map[string]string{
			"networkid": networkID,
			"deviceid":  value.Serial,
		}
		nwStat, _ := value.ConnectionStats.GetMapStr("DeviceNetworkConnectionStat", additionalMap)
		devicesStat = append(devicesStat, nwStat)
	}
	logp.Info("%+v", devicesStat)
	return devicesStat, err

}

func (mc *MerakiClient) GetDevicesLatencyStat(networkID string) (devicesStat []common.MapStr, err error) {
	netURL := fmt.Sprintf("%s/api/v0/networks/%s/devices/latencyStats", mc.URL, networkID)

	body, err := mc.getData(netURL)
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

	for _, value := range data {
		additionalMap := map[string]string{
			"networkid": networkID,
			"deviceid":  value.Serial,
		}
		latencyStat, _ := value.LatencyStats.GetMapStr("DeviceNetworkLatencyStat", additionalMap)
		devicesStat = append(devicesStat, latencyStat)
	}
	logp.Info("%+v", devicesStat)
	return devicesStat, err

}

func (mc *MerakiClient) GetClientConnectionStat(networkID string) (clientsStat []common.MapStr, err error) {
	netURL := fmt.Sprintf("%s/api/v0/networks/%s/clients/connectionStats", mc.URL, networkID)

	body, err := mc.getData(netURL)
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

	for _, value := range data {
		additionalMap := map[string]string{
			"networkid":  networkID,
			"clientid":   value.MAC,
			"client.Mac": value.MAC,
		}
		nwStat, _ := value.ConnectionStats.GetMapStr("ClientNetworkConnectionStat", additionalMap)
		clientsStat = append(clientsStat, nwStat)
	}
	logp.Info("%+v", clientsStat)
	return clientsStat, err

}

func (mc *MerakiClient) GetClientLatencyStat(networkID string) (clientsStat []common.MapStr, err error) {
	netURL := fmt.Sprintf("%s/api/v0/networks/%s/clients/latencyStats", mc.URL, networkID)

	body, err := mc.getData(netURL)
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

	for _, value := range data {
		additionalMap := map[string]string{
			"networkid":  networkID,
			"clientid":   value.MAC,
			"client.Mac": value.MAC,
		}
		latencyStat, _ := value.LatencyStats.GetMapStr("ClientNetworkLatencyStat", additionalMap)
		clientsStat = append(clientsStat, latencyStat)
	}
	logp.Info("%+v", clientsStat)
	return clientsStat, err

}
