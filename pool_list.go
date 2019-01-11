package gplib

import (
	"fmt"
	"github.com/levigross/grequests"
	"net/http"
)

type PoolListItem struct {
	PoolId  string `json:"poolid"`
	Comment string `json:"comment"`
}

type PoolListItemContainer struct {
	Data []PoolListItem `json:"data"`
}

func (c *proxmoximpl) GetPoolList() ([]PoolListItem, error) {

	tokenCookie := http.Cookie{
		Name:  "PVEAuthCookie",
		Value: c.ticket,
	}
	reqopt := grequests.RequestOptions{
		Cookies:            []*http.Cookie{&tokenCookie},
		InsecureSkipVerify: c.insecureSkipVerify,
	}

	res, err := grequests.Get(fmt.Sprintf("%s/api2/json/pools", c.serverURL), &reqopt)
	if err != nil {
		return []PoolListItem{}, err
	}

	resp := PoolListItemContainer{}
	err = res.JSON(&resp)
	if err != nil {
		return []PoolListItem{}, err
	} else {
		return resp.Data, nil
	}
}
