package healthchecks

import (
	"testing"

	"github.com/streadway/amqp"
)

func TestAMQP(t *testing.T) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	checker := NewAMQPHealth(conn)
	ok, err := checker.Check()
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("Want true")
	}
}
