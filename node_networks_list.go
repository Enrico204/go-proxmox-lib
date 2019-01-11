package gplib

import (
	"fmt"
	"github.com/levigross/grequests"
	"net/http"
)

type NodeNetworkListItem struct {
	IFace     string    `json:"iface"`
	Type      string    `json:"type"`
	Active    *int      `json:"active"`
	Exists    *int      `json:"exists"`
	Autostart *int      `json:"autostart"`
	Priority  *int      `json:"priority"`
	Families  *[]string `json:"families"`
	Slaves    *string   `json:"slaves"`

	BridgeVLANAware *int    `json:"bridge_vlan_aware"`
	BridgeFD        *string `json:"bridge_fd"`
	BridgeSTP       *string `json:"bridge_stp"`
	BridgePorts     *string `json:"bridge_ports"`

	BondMode           *string `json:"bond_mode"`
	BondXmitHashPolicy *string `json:"bond_xmit_hash_policy"`

	OVSBonds   *string `json:"ovs_bonds"`
	OVSBridge  *string `json:"ovs_bridge"`
	OVSOptions *string `json:"ovs_options"`
	OVSPorts   *string `json:"ovs_ports"`
	OVSTag     int     `json:"ovs_tag"`

	Method   *string `json:"method"`
	Address  *string `json:"address"`
	Netmask  *string `json:"netmask"`
	Gateway  *string `json:"gateway"`
	Comments *string `json:"comments"`

	Method6   *string `json:"method6"`
	Address6  *string `json:"address6"`
	Netmask6  *string `json:"netmask6"`
	Gateway6  *string `json:"gateway6"`
	Comments6 *string `json:"comments6"`
}

type NodeNetworkListContainer struct {
	Data []NodeNetworkListItem `json:"data"`
}

func (c *proxmoximpl) GetNodeNetworks(node string, nictype string) ([]NodeNetworkListItem, error) {

	tokenCookie := http.Cookie{
		Name:  "PVEAuthCookie",
		Value: c.ticket,
	}
	reqopt := grequests.RequestOptions{
		Cookies:            []*http.Cookie{&tokenCookie},
		InsecureSkipVerify: c.insecureSkipVerify,
	}

	if nictype != "" {
		reqopt.Params = map[string]string{
			"type": nictype,
		}
	}

	res, err := grequests.Get(fmt.Sprintf("%s/api2/json/nodes/%s/network", c.serverURL, node), &reqopt)
	if err != nil {
		return []NodeNetworkListItem{}, err
	}

	resp := NodeNetworkListContainer{}
	err = res.JSON(&resp)
	if err != nil {
		return []NodeNetworkListItem{}, err
	} else {
		return resp.Data, nil
	}
}
