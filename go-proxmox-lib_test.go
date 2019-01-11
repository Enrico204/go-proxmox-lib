package gplib

import (
	"fmt"
	"os"
	"testing"
)

var c Proxmox

func TestMain(m *testing.M) {
	c = New(os.Getenv("HOST"), true)
	_, err := c.Login(os.Getenv("USER"), os.Getenv("PASS"))
	if err != nil {
		fmt.Println("Login error: ", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}