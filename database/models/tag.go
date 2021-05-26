package models

import (
	"github.com/jmoiron/sqlx"
	"ruehmkorf.com/database"
)

// language=sql
const CreateTagTable = `
CREATE TABLE "tag" (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    tag text UNIQUE NOT NULL 
)
`

type Tag struct {
	Id  string
	Tag string
}

func getTagsByTagList(tagList []string) ([]Tag, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	tags := new([]Tag)

	query, args, err := sqlx.In("SELECT * FROM \"tag\" WHERE tag IN (?)", tagList)
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)
	err = db.Select(tags, query, args...)

	return *tags, err
}

func containsTag(tag string, tags []Tag) bool {
	for _, item := range tags {
		if tag == item.Tag {
			return true
		}
	}

	return false
}

func CreateTags(splitTags []string) ([]Tag, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	tagList, err := getTagsByTagList(splitTags)
	if err != nil {
		return nil, err
	}

	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}

	for _, tag := range splitTags {
		if containsTag(tag, tagList) {
			continue
		}

		_, err = tx.Exec("INSERT INTO \"tag\" (tag) VALUES ($1)", tag)
		if err != nil {
			return nil, tx.Rollback()
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return getTagsByTagList(splitTags)
}

func FindTagsByNews(news News) ([]Tag, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	tags := new([]Tag)

	err = db.Select(tags, "SELECT t.* FROM \"tag\" t JOIN \"news_tag\" n ON n.news_id = $1 AND n.tag_id = t.id", news.Id)
	if err != nil {
		return nil, err
	}

	return *tags, err
}
