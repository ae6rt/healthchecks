package healthchecks

import (
	"testing"

	"gopkg.in/redis.v3"
)

func TestRedis(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	checker := NewRedisChecker(client)
	_, err := checker.Check()
	if err != nil {
		t.Fatal(err)
	}
}
