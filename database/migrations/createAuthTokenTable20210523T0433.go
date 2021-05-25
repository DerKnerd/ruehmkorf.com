package migrations

import (
	"ruehmkorf.com/database"
	"ruehmkorf.com/database/models"
)

type CreateAuthTokenTable20210523T0433 struct{}

func (migration CreateAuthTokenTable20210523T0433) Execute() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx := db.MustBegin()
	tx.MustExec(models.CreateAuthTokenTable)
	if err = tx.Commit(); err != nil {
		panic(err)
	}
}

func (migration CreateAuthTokenTable20210523T0433) GetVersion() string {
	return "CreateAuthTokenTable20210523T0433"
}
