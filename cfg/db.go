package cfg

import "fmt"

const (
	db_user = "at-cs"
	db_pass = "changeme"
	db_host = "localhost"
)

var (
	DBConnStr = fmt.Sprintf("postgres://%s:%s@%s/at-cs?sslmode=disable", db_user, db_pass, db_host)
)
