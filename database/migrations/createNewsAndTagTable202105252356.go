package migrations

import (
	"ruehmkorf.com/database"
	"ruehmkorf.com/database/models"
)

type CreateNewsAndTagTable202105252356 struct{}

func (migration CreateNewsAndTagTable202105252356) Execute() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx := db.MustBegin()
	tx.MustExec(models.CreateNewsTable)
	tx.MustExec(models.CreateTagTable)
	tx.MustExec(models.CreateNewsTagTable)
	if err = tx.Commit(); err != nil {
		panic(err)
	}
}

func (migration CreateNewsAndTagTable202105252356) GetVersion() string {
	return "CreateNewsAndTagTable202105252356"
}
