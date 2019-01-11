package gplib

import (
	"os"
	"testing"
)

func TestUsers(t *testing.T) {
	if os.Getenv("WRITE") == "" {
		t.SkipNow()
	}

	err := c.NewUser("testuser@pve", "test@user.tld", 0, "Test", "User", []string{}, "Test.1234")
	if err != nil {
		t.Error(err)
	}

	userlist, err := c.GetUserList()
	found := false
	for i := 0; !found && i < len(userlist); i++ {
		if userlist[i].UserId == "testuser@pve" {
			found = true
		}
	}
	if !found {
		t.Error("testuser@pve not found in user list after creation", userlist)
	}

	err = c.DeleteUser("testuser@pve")
	if err != nil {
		t.Error(err)
	}
}