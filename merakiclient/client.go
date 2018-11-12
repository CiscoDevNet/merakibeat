package merakiclient

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/elastic/beats/libbeat/logp"
)

type MerakiClient struct {
	URL         string
	Key         string
	OrgID       string
	NetworkIDs  []string
	Period      time.Duration
	VideoPeriod time.Duration
}

func NewMerakiClient(url, key, orgID string, networkIDs []string, period, videoPeriod time.Duration) MerakiClient {
	return MerakiClient{
		URL:         url,
		Key:         key,
		OrgID:       orgID,
		NetworkIDs:  networkIDs,
		Period:      period,
		VideoPeriod: videoPeriod,
	}
}

// THis is health API specific getData, not all API set are same to some
// custom t0-t1 handling
func (mc *MerakiClient) getDataQueryParam(netURL string, params map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", netURL, nil)
	if err != nil {
		logp.Info("Failed to connect Meraki API %s", err.Error())
		return nil, err
	}
	req.Header.Add("X-Cisco-Meraki-API-Key", mc.Key)

	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	logp.Info("Calling API URL %+v, %v", req.URL, req.URL.RawQuery)
	resp, err := client.Do(req)
	if err != nil {
		logp.Info("Failed to connect Meraki API %s", err.Error())
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	logp.Info("%s", string(body[:]))
	return body, err
}
