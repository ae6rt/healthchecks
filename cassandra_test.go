package healthchecks

import (
	"testing"

	"github.com/gocql/gocql"
)

func TestCassandra(t *testing.T) {
	cluster := gocql.NewCluster("127.0.0.1")
	// http://stackoverflow.com/questions/34846153/golang-gocql-failed-to-connect-to-127-0-0-19042-not-enough-bytes-to-read-header
	cluster.ProtoVersion = 4

	checker := NewCassandraChecker(cluster)
	if _, err := checker.Check(); err != nil {
		t.Fatal(err)
	}
}
