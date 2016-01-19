package healthchecks

type HealthChecker interface {
	Check() (bool, error)
}
