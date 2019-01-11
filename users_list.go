package gplib

import (
	"fmt"
	"github.com/levigross/grequests"
	"net/http"
)

type UserListItem struct {
	UserId string `json:"userid"`
	Email  string `json:"email"`
	Enable int    `json:"enable"`
	Expire int64  `json:"expire"`
}

type UserListResponseContainer struct {
	Data []UserListItem `json:"data"`
}

func (c *proxmoximpl) GetUserList() ([]UserListItem, error) {

	tokenCookie := http.Cookie{
		Name:  "PVEAuthCookie",
		Value: c.ticket,
	}
	reqopt := grequests.RequestOptions{
		Cookies:            []*http.Cookie{&tokenCookie},
		InsecureSkipVerify: c.insecureSkipVerify,
	}

	res, err := grequests.Get(fmt.Sprintf("%s/api2/json/access/users", c.serverURL), &reqopt)
	if err != nil {
		return []UserListItem{}, err
	}

	resp := UserListResponseContainer{}
	err = res.JSON(&resp)
	if err != nil {
		return []UserListItem{}, err
	} else {
		return resp.Data, nil
	}
}
