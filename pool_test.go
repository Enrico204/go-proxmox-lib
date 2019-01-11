package gplib

import (
	"os"
	"testing"
)

func TestPool(t *testing.T) {
	if os.Getenv("WRITE") == "" {
		t.SkipNow()
	}

	err := c.NewPool("pool-test", "")
	if err != nil {
		t.Error(err)
	}

	poollist, err := c.GetPoolList()
	found := false
	for i := 0; !found && i < len(poollist); i++ {
		if poollist[i].PoolId == "pool-test" {
			found = true
		}
	}
	if !found {
		t.Error("pool-test not found in pool list after creation", poollist)
	}

	err = c.DeletePool("pool-test")
	if err != nil {
		t.Error(err)
	}
}