package gplib

import (
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"net/http"
	"strings"
)

func (c *proxmoximpl) NewUser(userid string, email string, expire int, firstname string, lastname string, groups []string, password string) error {

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
			"userid": userid,
			"email": email,
			"expire": fmt.Sprint(expire),
			"firstname": firstname,
			"lastname": lastname,
			"groups": strings.Join(groups, ","),
			"password": password,
		},
	}

	res, err := grequests.Post(fmt.Sprintf("%s/api2/json/access/users", c.serverURL), &reqopt)
	if err != nil {
		return err
	} else if res.StatusCode >= 400 {
		return errors.New(res.RawResponse.Status)
	} else {
		return nil
	}
}
