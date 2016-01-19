package healthchecks

import (
	"database/sql"
	"fmt"
)

type HealthChecker interface {
	Check() (bool, error)
}

type MySQLChecker struct {
	db *sql.DB
}

func MySQLHealth(db *sql.DB) HealthChecker {
	c := MySQLChecker{db}
	return c
}

func (checker MySQLChecker) Check() (bool, error) {
	var r int
	if err := checker.db.QueryRow("select 1 from dual").Scan(&r); err != nil {
		return false, err
	}
	if r != 1 {
		return false, fmt.Errorf("MySQL checker: want 1 from select 1 from dual but got %d", r)
	}
	return true, nil
}
