package models

import (
	"database/sql"
	"os"
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

var IconPath = os.Getenv("DATA_DIR") + "/public/profile/icon/"
var HeaderPath = os.Getenv("DATA_DIR") + "/public/profile/header/"

type Profile struct {
	Id     string         `json:"id"`
	Name   string         `json:"name"`
	Url    string         `json:"url"`
	Active bool           `json:"active"`
	Icon   string         `json:"-"`
	Header sql.NullString `json:"-"`
}

func FindAllProfiles() ([]Profile, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	profiles := make([]Profile, 0)
	if err = db.Select(&profiles, "SELECT * FROM profile ORDER BY name"); err != nil {
		return nil, err
	}

	return profiles, nil
}

func FindAllActiveProfiles() ([]Profile, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	profiles := new([]Profile)
	if err = db.Select(profiles, "SELECT * FROM profile WHERE active = true ORDER BY name"); err != nil {
		return nil, err
	}

	return *profiles, nil
}

func FindProfileById(id string) (*Profile, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	profile := new(Profile)
	if err = db.Get(profile, "SELECT * FROM profile WHERE id = $1", id); err != nil {
		return nil, err
	}

	return profile, nil
}

func CreateProfile(profile Profile) (string, error) {
	db, err := database.Connect()
	if err != nil {
		return "", err
	}

	defer db.Close()
	_, err = db.Exec("INSERT INTO profile (name, url, active, icon, header) VALUES ($1, $2, $3, $4, $5)", profile.Name, profile.Url, profile.Active, profile.Icon, profile.Header.String)

	var id string
	err = db.Get(&id, "SELECT id FROM profile WHERE name = $1 and url = $2 and active = $3", profile.Name, profile.Url, profile.Active)

	return id, err
}

func GetProfileByName(name string) (*Profile, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	profile := new(Profile)
	if err = db.Get(profile, "SELECT * FROM profile WHERE name = $1", name); err != nil {
		return nil, err
	}

	return profile, nil
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
