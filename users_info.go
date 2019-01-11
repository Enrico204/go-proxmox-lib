package gplib

import (
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"net/http"
)

type UserInfo struct {
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	EMail     string   `json:"comment"`
	Groups    []string `json:"groups"`
	Expire    int64    `json:"expire"`
	Enable    int      `json:"enable"`
	Comment   string   `json:"comment"`
	Keys      string   `json:"keys"`
}

type UserInfoResponse struct {
	Data UserInfo `json:"data"`
}

func (c *proxmoximpl) GetUserInfo(userid string) (UserInfo, error) {

	tokenCookie := http.Cookie{
		Name:  "PVEAuthCookie",
		Value: c.ticket,
	}
	reqopt := grequests.RequestOptions{
		Cookies:            []*http.Cookie{&tokenCookie},
		InsecureSkipVerify: c.insecureSkipVerify,
	}

	res, err := grequests.Get(fmt.Sprintf("%s/api2/json/access/users/%s", c.serverURL, userid), &reqopt)
	if err != nil {
		return UserInfo{}, err
	}

	resp := UserInfoResponse{}
	err = res.JSON(&resp)
	if err != nil {
		return UserInfo{}, err
	} else if res.StatusCode >= 400 {
		return UserInfo{}, errors.New(res.RawResponse.Status)
	} else {
		return resp.Data, nil
	}
}
