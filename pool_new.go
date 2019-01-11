package gplib

import (
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"net/http"
)

func (c *proxmoximpl) NewPool(poolid string, comment string) error {

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
		Data: map[string]string{
			"poolid": poolid,
			"comment": comment,
		},
	}

	res, err := grequests.Post(fmt.Sprintf("%s/api2/json/pools", c.serverURL), &reqopt)
	if err != nil {
		return err
	} else if res.StatusCode >= 400 {
		return errors.New(res.RawResponse.Status)
	} else {
		return nil
	}
}
