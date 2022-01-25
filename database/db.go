package database

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

func Leveldb() {
	db, err := leveldb.OpenFile("db", nil)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
