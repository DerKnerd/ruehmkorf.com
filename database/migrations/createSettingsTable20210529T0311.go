package migrations

import (
	"ruehmkorf.com/database"
	"ruehmkorf.com/database/models"
)

type CreateSettingsTable20210529T0311 struct{}

func (migration CreateSettingsTable20210529T0311) Execute() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx := db.MustBegin()
	tx.MustExec(models.CreateSettingsTable)
	if err = tx.Commit(); err != nil {
		panic(err)
	}
}

func (migration CreateSettingsTable20210529T0311) GetVersion() string {
	return "CreateSettingsTable20210529T0311"
}
