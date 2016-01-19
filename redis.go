package healthchecks

import "gopkg.in/redis.v3"

type RedisChecker struct {
	client *redis.Client
}

func NewRedisChecker(client *redis.Client) HealthChecker {
	return RedisChecker{client}
}

func (checker RedisChecker) Check() (bool, error) {
	if _, err := checker.client.Ping().Result(); err != nil {
		return false, err
	}
	return true, nil
}
