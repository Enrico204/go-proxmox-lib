package gplib

import (
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"net/http"
)

type VersionInfo struct {
	RepoID   string `json:"repoid"`
	Release  string `json:"release"`
	Version  string `json:"version"`
	Keyboard string `json:"keyboard"`
}

type VersionInfoResponse struct {
	Data VersionInfo `json:"data"`
}

func (c *proxmoximpl) GetVersion() (VersionInfo, error) {

	tokenCookie := http.Cookie{
		Name:  "PVEAuthCookie",
		Value: c.ticket,
	}
	reqopt := grequests.RequestOptions{
		Cookies:            []*http.Cookie{&tokenCookie},
		InsecureSkipVerify: c.insecureSkipVerify,
	}

	res, err := grequests.Get(fmt.Sprintf("%s/api2/json/version", c.serverURL), &reqopt)
	if err != nil {
		return VersionInfo{}, err
	}

	resp := VersionInfoResponse{}
	err = res.JSON(&resp)
	if err != nil {
		return VersionInfo{}, err
	} else if res.StatusCode >= 400 {
		return VersionInfo{}, errors.New(res.RawResponse.Status)
	} else {
		return resp.Data, nil
	}
}
