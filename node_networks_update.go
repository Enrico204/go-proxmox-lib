package gplib

import (
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"net/http"
	"reflect"
	"strings"
)

func (c *proxmoximpl) UpdateNodeNetwork(node string, i NodeNetworkListItem) error {

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
			"iface": i.IFace,
			"node":  node,
			"type":  i.Type,
		},
	}

	val := reflect.ValueOf(i).Elem()
	todelete := make([]string, 0)
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag.Get("json")
		value := ""

		if tag != "address" && tag != "address6" && tag != "autostart" && tag != "bond_mode" && tag != "bond_xmit_hash_policy" &&
			tag != "bridge_ports" && tag != "bridge_vlan_aware" && tag != "comments" && tag != "comments6" && tag != "gateway" &&
			tag != "gateway6" && tag != "netmask" && tag != "netmask6" && tag != "ovs_bonds" && tag != "ovs_bridge" &&
			tag != "ovs_options" && tag != "ovs_tag" && tag != "slaves" {
			continue
		}

		if valueField.IsNil() {
			todelete = append(todelete, tag)
			continue
		}

		if valueField.Kind() == reflect.Int || valueField.Kind() == reflect.Int8 || valueField.Kind() == reflect.Int16 ||
			valueField.Kind() == reflect.Int32 || valueField.Kind() == reflect.Int64 {
			value = fmt.Sprint(valueField.Int())
		} else if valueField.Kind() == reflect.Uint || valueField.Kind() == reflect.Uint8 || valueField.Kind() == reflect.Uint16 ||
			valueField.Kind() == reflect.Uint32 || valueField.Kind() == reflect.Uint64 {
			value = fmt.Sprint(valueField.Uint())
		}

		reqopt.Data[tag] = value
	}
	if len(todelete) > 0 {
		reqopt.Data["delete"] = strings.Join(todelete, ",")
	}

	res, err := grequests.Put(fmt.Sprintf("%s/api2/json/nodes/%s/network/%s", c.serverURL, node, i.IFace), &reqopt)
	if err != nil {
		return err
	} else if res.StatusCode >= 400 {
		return errors.New(res.RawResponse.Status)
	} else {
		return nil
	}
}
