package migrations

import (
	"ruehmkorf.com/database"
	"ruehmkorf.com/database/models"
)

type CreateProfileTable20210527T0107 struct{}

func (migration CreateProfileTable20210527T0107) Execute() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx := db.MustBegin()
	tx.MustExec(models.CreateProfileTable)
	if err = tx.Commit(); err != nil {
		panic(err)
	}
}

func (migration CreateProfileTable20210527T0107) GetVersion() string {
	return "CreateProfileTable20210527T0107"
}
