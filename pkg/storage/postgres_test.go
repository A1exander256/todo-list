package storage

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
)

type cfgTest struct {
	host     string
	port     string
	username string
	password string
	dbname   string
	sslmode  string
}

func newCfgTest() *cfgTest {
	return &cfgTest{
		host:     "localhost",
		port:     "5432",
		username: "postgres",
		password: "postgres",
		dbname:   "tado",
		sslmode:  "disable",
	}
}

func TestNewPostgresDB(t *testing.T) {
	cfg := newCfgTest()
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.host,
		cfg.port,
		cfg.username,
		cfg.password,
		cfg.dbname,
		cfg.sslmode)
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		t.Fatal("error connecting to postgres :", err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal("postgres connection check error :", err)
	}
}
