package migrations

import (
	"ruehmkorf.com/database"
)

type FileExtension20210531T2303 struct{}

func (migration FileExtension20210531T2303) Execute() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx := db.MustBegin()
	tx.MustExec("ALTER TABLE download ADD file_extension TEXT NULL")
	if err = tx.Commit(); err != nil {
		panic(err)
	}
}

func (migration FileExtension20210531T2303) GetVersion() string {
	return "FileExtension20210531T2303"
}
