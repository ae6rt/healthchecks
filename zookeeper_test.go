package healthchecks

import (
	"testing"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func TestZK(t *testing.T) {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		t.Fatal(err)
	}

	checker := NewZookeeperChecker(c)
	if err != nil {
		t.Fatal(err)
	}

	_, err = checker.Check()
	if err != nil {
		t.Fatal(err)
	}
}
