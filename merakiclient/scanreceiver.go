package merakiclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/elastic/beats/libbeat/beat"
)

//api/v1/scanning/receiver/

type ScanReceiver struct {
	Secret     string
	Validator  string
	Version    string
	Mux        *http.ServeMux
	HostPort   string
	BeatClient beat.Client
}

func NewScanReceiver(secret, validator string, bc beat.Client) *ScanReceiver {

	sr := ScanReceiver{
		Secret:     secret,
		Validator:  validator,
		Version:    "2.0",
		Mux:        http.NewServeMux(),
		HostPort:   ":5001",
		BeatClient: bc,
	}
	sr.Mux.HandleFunc("/api/v1/scanning/receiver/", sr.handleReceive)

	return &sr
}

func (sr *ScanReceiver) handleReceive(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Entering handleReceive")
	switch r.Method {
	case http.MethodGet:
		sr.handleReceiveValidation(w, r)
		return
	case http.MethodPost:
		sr.handleReceiveData(w, r)
		return
	default:
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	}
}
func (sr *ScanReceiver) handleReceiveValidation(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, sr.Validator)
	return
}

func (sr *ScanReceiver) handleReceiveData(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Entering scanreciever\n")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading body %s\n", err.Error())
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
		return
	}
	fmt.Printf("Body %s", string(body[:]))
	var scanData ScanData
	err = json.Unmarshal(body, &scanData)
	if err != nil {
		fmt.Printf("Error unmarhaling json body %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if scanData.Secret != sr.Secret {
		fmt.Printf("Error Ivalid secret \n")
		http.Error(w, "Invalid Secrect", http.StatusMethodNotAllowed)
		return
	}
	fmt.Printf("Publishing scan data %+v\n", scanData)
	mapstrArr, err := scanData.GetMapStr("MerakiScanEvent", map[string]string{})
	for _, mapStr := range mapstrArr {
		seenTime, _ := mapStr.GetValue("client.seenTime")
		seenTimeStr, _ := seenTime.(string)
		ts, err := time.Parse("2006-01-02T15:04:05.999999999", seenTimeStr)
		fmt.Printf("Timestamp %s %+v", seenTimeStr, ts)
		if err != nil {
			ts = time.Now()
		}
		sr.BeatClient.Publish(beat.Event{
			Timestamp: ts,
			Fields:    mapStr,
		})
		fmt.Printf("Published event %+v\n", mapStr)
	}
	return
}

func (sr *ScanReceiver) Run() {
	log.Fatal(http.ListenAndServe(sr.HostPort, sr.Mux))
}
