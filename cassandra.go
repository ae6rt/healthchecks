package healthchecks

import (
	"fmt"

	"github.com/gocql/gocql"
)

type CassandraChecker struct {
	cluster *gocql.ClusterConfig
}

func (checker CassandraChecker) Check() (bool, error) {
	session, _ := checker.cluster.CreateSession()
	defer session.Close()

	var text string

	if err := session.Query("SELECT now() FROM system.local").Scan(&text); err != nil {
		return false, err
	}

	fmt.Println(text)
	return true, nil
}

func NewCassandraChecker(c *gocql.ClusterConfig) HealthChecker {
	return CassandraChecker{c}
}
