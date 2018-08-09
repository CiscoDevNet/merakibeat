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

func (ns *NetworkStat) GetMapStr(networkID string) (common.MapStr, error) {
	return common.MapStr{
		"type":    "NetworkConnectionStat",
		"assoc":   ns.Assoc,
		"auth":    ns.Auth,
		"dhcp":    ns.Dhcp,
		"dns":     ns.DNS,
		"success": ns.Success,
		"object":  networkID,
	}, nil
}
