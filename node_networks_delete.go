package gplib

import (
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"net/http"
)

func (c *proxmoximpl) DeleteNodeNetwork(node string, iface string) error {

	tokenCookie := http.Cookie{
		Name:  "PVEAuthCookie",
		Value: c.ticket,
	}
	reqopt := grequests.RequestOptions{
		Cookies:            []*http.Cookie{&tokenCookie},
		InsecureSkipVerify: c.insecureSkipVerify,
		Headers: map[string]string {
			"CSRFPreventionToken": c.csrfPreventionToken,
		},
	}

	res, err := grequests.Delete(fmt.Sprintf("%s/api2/json/nodes/%s/network/%s", c.serverURL, node, iface), &reqopt)
	if err != nil {
		return err
	} else if res.StatusCode >= 400 {
		return errors.New(res.RawResponse.Status)
	} else {
		return nil
	}
}
