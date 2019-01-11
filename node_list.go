package gplib

import (
	"fmt"
	"github.com/levigross/grequests"
	"net/http"
)

type NodeListItem struct {
	ID             string `json:"id"`
	Type           string `json:"type"`
	Node           string `json:"node"`
	Status         string `json:"status"`
	Level          string `json:"level"`
	Uptime         int64  `json:"uptime"`
	SSLFingerprint string `json:"ssl_fingerprint"`

	Mem    int64 `json:"mem"`
	MaxMem int64 `json:"maxmem"`

	CPU    float64 `json:"cpu"`
	MaxCPU float64 `json:"maxcpu"`

	Disk int64 `json:"disk"`
}

type NodeListContainer struct {
	Data []NodeListItem `json:"data"`
}

func (c *proxmoximpl) GetNodeList() ([]NodeListItem, error) {

	tokenCookie := http.Cookie{
		Name:  "PVEAuthCookie",
		Value: c.ticket,
	}
	reqopt := grequests.RequestOptions{
		Cookies:            []*http.Cookie{&tokenCookie},
		InsecureSkipVerify: c.insecureSkipVerify,
	}

	res, err := grequests.Get(fmt.Sprintf("%s/api2/json/nodes", c.serverURL), &reqopt)
	if err != nil {
		return []NodeListItem{}, err
	}

	resp := NodeListContainer{}
	err = res.JSON(&resp)
	if err != nil {
		return []NodeListItem{}, err
	} else {
		return resp.Data, nil
	}
}
