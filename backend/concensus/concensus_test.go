package concensus

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)


func Test_DbFileCreation(t *testing.T) {
	dir, err := ioutil.TempDir("", "rqlite-test-")
	defer os.RemoveAll(dir)

	dbPath := path.Join(dir, "test_db")

	db, err := sql.Open("sqlite3" , dbPath)
	if err != nil {
		t.Fatalf("failed to open new database: %s", err.Error())
	}
	if db == nil {
		t.Fatal("database is nil")
	}

	err = db.Close()
	if err != nil {
		t.Fatalf("failed to close database: %s", err.Error())
	}
}

func Test_TableCreation(t *testing.T) {
	db, path := mustCreateDatabase()
	defer db.Close()
	defer os.Remove(path)

	r, err := db.Exec("CREATE TABLE foo (id INTEGER NOT NULL PRIMARY KEY, name TEXT)")
	if err != nil {
		t.Fatalf("failed to create table: %s", err.Error())
	}
	if exp, got := `{"Locker":{}}`, asJSON(r); exp != got {
		t.Fatalf("unexpected results for query, expected %s, got %s", exp, got)
	}

	q, err := db.Query("SELECT * FROM foo")
	if err != nil {
		t.Fatalf("failed to query empty table: %s", err.Error())
	}
	if exp, got := `{}`, asJSON(q); exp != got {
		t.Fatalf("unexpected results for query, expected %s, got %s", exp, got)
	}
}

func mustCreateDatabase() (*sql.DB, string) {
	var err error
	f := mustTempFile()
	db, err := sql.Open("sqlite3", f)
	if err != nil {
		panic("failed to open database")
	}

	return db, f
}

func mustTempFile() string {
	tmpfile, err := ioutil.TempFile("", "rqlite-db-test")
	if err != nil {
		panic(err.Error())
	}
	tmpfile.Close()
	return tmpfile.Name()
}

func asJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Sprintf("failed to JSON marshal value: %s", err.Error()))
	}
	return string(b)
}