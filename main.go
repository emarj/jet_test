package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./db.sqlite")
	if err != nil {
		panic(err)
	}

	accs, err := GetAccountWithEntity(db)
	if err != nil {
		panic(err)
	}

	d, _ := json.MarshalIndent(accs, "", "\t")
	fmt.Println(string(d))

	txs, err := GetOperationsByEntity(db, 1)
	if err != nil {
		panic(err)
	}

	d, _ = json.MarshalIndent(txs, "", "\t")
	fmt.Println(string(d))
}
