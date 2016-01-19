package healthchecks

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMySQL(t *testing.T) {
	db, err := sql.Open("mysql", "root:r00t@tcp(127.0.0.1:3306)/")
	if err != nil {
		t.Fatal(err)
	}
	checker := MySQLHealth(db)
	ok, err := checker.Check()
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("Want true")
	}
}
