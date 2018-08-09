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

	return data.GetMapStr(networkID)

}

func NewMerakiClient(url, key, orgID string, networkIDs []string) MerakiClient {
	return MerakiClient{
		URL:        url,
		Key:        key,
		OrgID:      orgID,
		NetworkIDs: networkIDs,
	}
}
