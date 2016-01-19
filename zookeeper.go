package healthchecks

import "github.com/samuel/go-zookeeper/zk"

type ZookeeperChecker struct {
	c *zk.Conn
}

func (checker ZookeeperChecker) Check() (bool, error) {
	_, _, _, err := checker.c.ChildrenW("/")
	if err != nil {
		return false, err
	}

	return true, nil
}

func NewZookeeperChecker(conn *zk.Conn) ZookeeperChecker {
	return ZookeeperChecker{conn}
}
