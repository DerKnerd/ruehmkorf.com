package models

import "ruehmkorf.com/database"

//language=sql
const CreateSettingsTable = `
CREATE TABLE settings (
    "key" text primary key,
    value text
)
`

type Setting struct {
	Key   string
	Value string
}

func FindSettingByKey(key string) string {
	db, err := database.Connect()
	if err != nil {
		return ""
	}

	defer db.Close()
	setting := new(Setting)

	if err = db.Get(setting, "SELECT * FROM settings WHERE key = $1", key); err != nil {
		return ""
	}

	return setting.Value
}

func UpdateSetting(key string, value string) {
	db, err := database.Connect()
	if err != nil {
		return
	}

	defer db.Close()

	var count int
	if err = db.Get(&count, "SELECT count(*) FROM settings WHERE key = $1", key); err != nil {
		return
	}

	if count > 0 {
		if _, err = db.Exec("UPDATE settings SET value = $1 WHERE key = $2", value, key); err != nil {
			return
		}
	} else {
		if _, err = db.Exec("INSERT INTO settings (value, \"key\") VALUES ($1, $2)", value, key); err != nil {
			return
		}
	}
}
