package models

import (
	"database/sql"
	"ruehmkorf.com/database"
)

//language=sql
const CreateProfileTable = `
CREATE TABLE profile (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    url text NOT NULL,
    active boolean NOT NULL DEFAULT false,
    icon text NOT NULL,
    header text NULL
)
`

type Profile struct {
	Id     string
	Name   string
	Url    string
	Active bool
	Icon   string
	Header sql.NullString
}

func FindAllProfiles() ([]Profile, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	profiles := new([]Profile)
	if err = db.Select(profiles, "SELECT * FROM profile ORDER BY name"); err != nil {
		return nil, err
	}

	return *profiles, nil
}

func CreateProfile(profile Profile) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec("INSERT INTO profile (name, url, active, icon, header) VALUES ($1, $2, $3, $4, $5)", profile.Name, profile.Url, profile.Active, profile.Icon, profile.Header.String)

	return err
}

func UpdateProfile(profile Profile) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec("UPDATE profile SET name = $1, url = $2, active = $3, icon = $4, header = $5 WHERE id = $6", profile.Name, profile.Url, profile.Active, profile.Icon, profile.Header.String, profile.Id)

	return err
}

func DeleteProfile(id string) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec("DELETE FROM profile WHERE id = $1", id)

	return err
}
