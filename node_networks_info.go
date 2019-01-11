package gplib

import (
	"fmt"
	"github.com/levigross/grequests"
	"net/http"
)

type NodeNetworkInfoContainer struct {
	Data NodeNetworkListItem `json:"data"`
}

func (c *proxmoximpl) GetNodeNetworkInfo(node string, iface string) (NodeNetworkListItem, error) {

	tokenCookie := http.Cookie{
		Name:  "PVEAuthCookie",
		Value: c.ticket,
	}
	reqopt := grequests.RequestOptions{
		Cookies:            []*http.Cookie{&tokenCookie},
		InsecureSkipVerify: c.insecureSkipVerify,
	}

	res, err := grequests.Get(fmt.Sprintf("%s/api2/json/nodes/%s/network/%s", c.serverURL, node, iface), &reqopt)
	if err != nil {
		return NodeNetworkListItem{}, err
	}

	resp := NodeNetworkInfoContainer{}
	err = res.JSON(&resp)
	if err != nil {
		return NodeNetworkListItem{}, err
	} else {
		return resp.Data, nil
	}
}
