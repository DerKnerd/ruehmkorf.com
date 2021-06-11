package migrations

import (
	"ruehmkorf.com/database"
)

type FixDateColumnInDownloadTable20210611T1911 struct{}

func (migration FixDateColumnInDownloadTable20210611T1911) Execute() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx := db.MustBegin()
	tx.MustExec("ALTER TABLE download DROP COLUMN date")
	tx.MustExec("ALTER TABLE download ADD date date DEFAULT CURRENT_TIMESTAMP")
	if err = tx.Commit(); err != nil {
		panic(err)
	}
}

func (migration FixDateColumnInDownloadTable20210611T1911) GetVersion() string {
	return "CreateDownloadTable20210528T2344"
}
