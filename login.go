package gplib

import (
	"errors"
	"fmt"
	"github.com/levigross/grequests"
)

type LoginResponse struct {
	CSRFPreventionToken string       `json:"CSRFPreventionToken"`
	Username            string       `json:"username"`
	Ticket              string       `json:"ticket"`
	UserCapabilities    Capabilities `json:"cap"`
}

type LoginResponseContainer struct {
	Data LoginResponse `json:"data"`
}

func (c *proxmoximpl) Login(username string, password string) (LoginResponse, error) {
	reqopt := grequests.RequestOptions{
		Data: map[string]string{
			"username": username,
			"password": password,
		},
		InsecureSkipVerify: c.insecureSkipVerify,
	}

	res, err := grequests.Post(fmt.Sprintf("%s/api2/json/access/ticket", c.serverURL), &reqopt)
	if err != nil {
		return LoginResponse{}, err
	}

	resp := LoginResponseContainer{}
	err = res.JSON(&resp)
	if err != nil {
		return LoginResponse{}, err
	} else if res.StatusCode >= 400 {
		return LoginResponse{}, errors.New(res.RawResponse.Status)
	} else {
		c.csrfPreventionToken = resp.Data.CSRFPreventionToken
		c.ticket = resp.Data.Ticket
		c.username = resp.Data.Username
		return resp.Data, nil
	}
}
