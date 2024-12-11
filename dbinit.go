package main

import (
	"context"
	"database/sql"
	_ "embed"
	"ethUpdateNotifier/dbutil"

	_ "github.com/glebarez/go-sqlite"
)

//go:embed schema.sql
var ddl string

func init() {
	dbctx = context.Background()

	db, err := sql.Open("sqlite", "versions.db")
	checkErr(err)

	_, err = db.ExecContext(dbctx, ddl)
	checkErr(err)

	queries = dbutil.New(db)
}
