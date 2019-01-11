package gplib

import (
	"fmt"
	"github.com/levigross/grequests"
	"net/http"
)

type NodeVMListItem struct {
	VMID     string `json:"string"`
	Name     string `json:"name"`
	PID      *int64 `json:"pid"`
	Status   string `json:"status"`
	Uptime   int64  `json:"uptime"`
	Template string `json:"template"`

	CPU  float64 `json:"cpu"`
	CPUs int     `json:"cpus"`

	Mem int64 `json:"mem"`

	NetIn  int64 `json:"netin"`
	NetOut int64 `json:"netout"`

	Disk      int64 `json:"disk"`
	DiskRead  int64 `json:"diskread"`
	DiskWrite int64 `json:"diskwrite"`
	MaxDisk   int64 `json:"maxdisk"`
}

type NodeVMListContainer struct {
	Data []NodeVMListItem `json:"data"`
}

func (c *proxmoximpl) GetNodeVMs(node string) ([]NodeVMListItem, error) {

	tokenCookie := http.Cookie{
		Name:  "PVEAuthCookie",
		Value: c.ticket,
	}
	reqopt := grequests.RequestOptions{
		Cookies:            []*http.Cookie{&tokenCookie},
		InsecureSkipVerify: c.insecureSkipVerify,
	}

	res, err := grequests.Get(fmt.Sprintf("%s/api2/json/nodes/%s/qemu", c.serverURL, node), &reqopt)
	if err != nil {
		return []NodeVMListItem{}, err
	}

	resp := NodeVMListContainer{}
	err = res.JSON(&resp)
	if err != nil {
		return []NodeVMListItem{}, err
	} else {
		return resp.Data, nil
	}
}
