package main

import (
	"database/sql"

	j "github.com/emarj/go-jet_test/gen/table"

	jet "github.com/go-jet/jet/v2/sqlite"
)

func GetAccountWithEntity(db *sql.DB) ([]Account, error) {

	Owner := j.Account.AS("owner")

	stmt := jet.SELECT(j.Account.AllColumns,
		Owner.AllColumns,
	).FROM(
		j.Account.INNER_JOIN(Owner, Owner.ID.EQ(j.Account.OwnerID)),
	)

	accounts := []Account{}

	err := stmt.Query(db, &accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func GetOperationsByEntity(db *sql.DB, eID int) ([]Operation, error) {

	From := j.Account.AS("from")
	To := j.Account.AS("to")

	FromEntity := j.Entity.AS("from_owner")
	ToEntity := j.Entity.AS("to.owner")

	stmt := jet.SELECT(
		j.Operation.AllColumns,
		From.AllColumns,
		To.AllColumns,
		FromEntity.AllColumns,
		ToEntity.AllColumns,
	).FROM(
		j.Operation.INNER_JOIN(
			From,
			From.ID.EQ(j.Operation.FromID),
		).INNER_JOIN(
			To,
			To.ID.EQ(j.Operation.ToID),
		).INNER_JOIN(
			FromEntity,
			FromEntity.ID.EQ(From.OwnerID),
		).INNER_JOIN(
			ToEntity,
			ToEntity.ID.EQ(To.OwnerID),
		),
	).WHERE(
		FromEntity.ID.EQ(jet.Int(int64(eID))).OR(ToEntity.ID.EQ(jet.Int(int64(eID)))),
	)

	//fPrintln(stDebugSql())

	operations := []Operation{}

	err := stmt.Query(db, &operations)
	if err != nil {
		return nil, err
	}
	return operations, nil
}
