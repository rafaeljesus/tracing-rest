package migrations

import (
	"database/sql"
)

func Up_20160918155111(txn *sql.Tx) {
	_, err := txn.Exec(
		"create table events (" +
			"id serial primary key," +
			"name text not null," +
			"status text," +
			"payload jsonb," +
			"created_at timestamptz not null," +
			"updated_at timestamptz not null," +
			"deleted_at timestamptz);")
	if err != nil {
		panic(err)
	}
}

func Down_20160918155111(txn *sql.Tx) {
	_, err := txn.Exec("drop table events;")
	if err != nil {
		panic(err)
	}
}
