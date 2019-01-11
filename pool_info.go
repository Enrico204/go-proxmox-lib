package gplib

import (
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"net/http"
)

type PoolMemberInfo struct {
	ID        string `json:"id"`
	VMID      int    `json:"vmid"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Node      string `json:"node"`
	Uptime    int64  `json:"uptime"`
	Template  int    `json:"template"`
	Disk      int64  `json:"disk"`
	DiskRead  int64  `json:"diskread"`
	DiskWrite int64  `json:"diskwrite"`
	MaxDisk   int64  `json:"maxdisk"`
	CPU       int    `json:"cpu"`
	MaxCPU    int    `json:"maxcpu"`
	Mem       int64  `json:"mem"`
	MaxMem    int64  `json:"maxmem"`
	NetIn     int64  `json:"netin"`
	NetOut    int64  `json:"netout"`
}

type PoolInfo struct {
	Comment string           `json:"comment"`
	Members []PoolMemberInfo `json:"members"`
}

type PoolInfoResponse struct {
	Data PoolInfo `json:"data"`
}

func (c *proxmoximpl) GetPoolInfo(poolid string) (PoolInfo, error) {

	tokenCookie := http.Cookie{
		Name:  "PVEAuthCookie",
		Value: c.ticket,
	}
	reqopt := grequests.RequestOptions{
		Cookies:            []*http.Cookie{&tokenCookie},
		InsecureSkipVerify: c.insecureSkipVerify,
		Headers: map[string]string{
			"CSRFPreventionToken": c.csrfPreventionToken,
		},
		Data: map[string]string{
			"poolid": poolid,
		},
	}

	res, err := grequests.Post(fmt.Sprintf("%s/api2/json/pool/%s", c.serverURL, poolid), &reqopt)
	if err != nil {
		return PoolInfo{}, err
	}

	resp := PoolInfoResponse{}
	err = res.JSON(&resp)
	if err != nil {
		return PoolInfo{}, err
	} else if res.StatusCode >= 400 {
		return PoolInfo{}, errors.New(res.RawResponse.Status)
	} else {
		return resp.Data, nil
	}
}
