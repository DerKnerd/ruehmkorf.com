package migrations

import (
	"ruehmkorf.com/database"
	"ruehmkorf.com/database/models"
)

type CreateUserTable20210521T2126 struct{}

func (migration CreateUserTable20210521T2126) Execute() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx := db.MustBegin()
	tx.MustExec(models.CreateUserTable)
	if err = tx.Commit(); err != nil {
		panic(err)
	}
}

func (migration CreateUserTable20210521T2126) GetVersion() string {
	return "CreateUserTable20210521T2126"
}
