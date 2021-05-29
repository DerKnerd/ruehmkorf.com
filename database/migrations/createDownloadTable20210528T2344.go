package migrations

import (
	"ruehmkorf.com/database"
	"ruehmkorf.com/database/models"
)

type CreateDownloadTable20210528T2344 struct{}

func (migration CreateDownloadTable20210528T2344) Execute() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx := db.MustBegin()
	tx.MustExec(models.CreateDownloadTable)
	if err = tx.Commit(); err != nil {
		panic(err)
	}
}

func (migration CreateDownloadTable20210528T2344) GetVersion() string {
	return "CreateDownloadTable20210528T2344"
}
