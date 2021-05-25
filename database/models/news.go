package models

import (
	"database/sql"
	"ruehmkorf.com/database"
	"time"
)

// language=sql
var CreateNewsTable = `
CREATE TABLE "news" (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    title_de text NOT NULL,
    title_en text NOT NULL,
    slug text UNIQUE NOT NULL,
    date date NOT NULL,
    hero_image text NOT NULL,
    public boolean NOT NULL DEFAULT true,
    gist_de text NULL,
    gist_en text NULL,
    content_de text NULL,
    content_en text NULL
)`

// language=sql
var CreateNewsTagTable = `
CREATE TABLE "news_tag" (
    tag_id uuid,
    news_id uuid,
    PRIMARY KEY (tag_id, news_id)
)
`

type News struct {
	Id        string
	TitleDe   string `db:"title_de"`
	TitleEn   string `db:"title_en"`
	Slug      string
	Date      time.Time
	Tags      []Tag
	HeroImage string `db:"hero_image"`
	Public    bool
	GistDe    sql.NullString `db:"gist_de"`
	GistEn    sql.NullString `db:"gist_en"`
	ContentDe sql.NullString `db:"content_de"`
	ContentEn sql.NullString `db:"content_en"`
}

func FindAllNews(offset int, count int) ([]News, int, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, 0, err
	}

	defer db.Close()
	news := new([]News)

	if err = db.Select(news, "SELECT * FROM \"news\" LIMIT $1 OFFSET $2", count, offset); err != nil {
		return nil, 0, err
	}

	var totalCount int
	if err = db.Get(&totalCount, "SELECT COUNT(*) FROM \"news\""); err != nil {
		return *news, len(*news), err
	}

	return *news, totalCount, nil
}
