package postgres

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func newTestDB(t *testing.T) (*pgxpool.Pool, func()) {
	db, err := pgxpool.Connect(context.Background(), "postgres://postgres:i_omara03@localhost:5434/tests_cws")
	if err != nil {
		t.Fatal(err)
	}

	script, err := ioutil.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(context.Background(), string(script))
	if err != nil {
		t.Fatal(err)
	}

	return db, func() {
		script, err := ioutil.ReadFile("./testdata/tear.sql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = db.Exec(context.Background(), string(script))
		if err != nil {
			t.Fatal(err)
		}

		db.Close()
	}
}
