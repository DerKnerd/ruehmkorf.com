package migrations

import "ruehmkorf.com/database"

type Migration interface {
	Execute()
	GetVersion() string
}

// language=sql
var migrationsTable = `
CREATE TABLE IF NOT EXISTS migrations (
    Version varchar(255) PRIMARY KEY
)
`

var Migrations = []Migration{
	new(CreateUserTable20210521T2126),
	new(CreateAuthTokenTable20210523T0433),
	new(CreateNewsAndTagTable202105252356),
	new(CreateProfileTable20210527T0107),
	new(CreateDownloadTable20210528T2344),
	new(CreateSettingsTable20210529T0311),
}

func createMigrationsTable() error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec(migrationsTable)
	if err != nil {
		return err
	}

	return nil
}

func saveMigration(version string) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec("INSERT INTO Migrations (version) VALUES ($1)", version)

	return err
}

func isMigrated(version string) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		return true, err
	}

	defer db.Close()

	var count int
	err = db.Get(&count, "SELECT COUNT(*) FROM migrations WHERE version = $1", version)
	if err != nil {
		return true, err
	}

	if count == 1 {
		return true, nil
	}

	return false, err
}

func Migrate() error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	_ = db.MustExec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	err = createMigrationsTable()
	if err != nil {
		return err
	}

	for _, migration := range Migrations {
		version := migration.GetVersion()
		migrated, err := isMigrated(version)
		if err != nil {
			return err
		}

		if !migrated {
			migration.Execute()
			err = saveMigration(version)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
