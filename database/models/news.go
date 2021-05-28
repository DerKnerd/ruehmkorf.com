package models

import (
	"database/sql"
	"ruehmkorf.com/database"
	"time"
)

// language=sql
const CreateNewsTable = `
CREATE TABLE "news" (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    title_de text NOT NULL,
    title_en text NOT NULL,
    slug text UNIQUE NOT NULL,
    date date NOT NULL,
    public boolean NOT NULL DEFAULT true,
    gist_de text NULL,
    gist_en text NULL,
    content_de text NULL,
    content_en text NULL
)`

// language=sql
const CreateNewsTagTable = `
CREATE TABLE "news_tag" (
    tag_id uuid,
    news_id uuid,
    PRIMARY KEY (tag_id, news_id)
)
`

const HeroPath = "./data/public/news/hero/"

type News struct {
	Id        string
	TitleDe   string `db:"title_de"`
	TitleEn   string `db:"title_en"`
	Slug      string
	Date      time.Time
	Tags      []Tag
	Public    bool
	GistDe    sql.NullString `db:"gist_de"`
	GistEn    sql.NullString `db:"gist_en"`
	ContentDe sql.NullString `db:"content_de"`
	ContentEn sql.NullString `db:"content_en"`
}

func FindAllNews(offset int, limit int) ([]News, int, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, 0, err
	}

	defer db.Close()
	news := new([]News)

	if err = db.Select(news, "SELECT * FROM \"news\" ORDER BY slug LIMIT $1 OFFSET $2", limit, offset); err != nil {
		return nil, 0, err
	}

	var totalCount int
	if err = db.Get(&totalCount, "SELECT COUNT(*) FROM \"news\""); err != nil {
		return *news, len(*news), err
	}

	return *news, totalCount, nil
}

func FindNewsBySlug(slug string) (*News, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	news := new(News)

	if err = db.Get(news, "SELECT * FROM \"news\" WHERE slug = $1", slug); err != nil {
		return nil, err
	}

	tags, err := FindTagsByNews(*news)
	if err != nil {
		return nil, err
	}

	news.Tags = tags

	return news, nil
}

func CreateNews(news News) (*News, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO \"news\" (title_de, title_en, slug, \"date\", \"public\", gist_de, gist_en, content_de, content_en) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", news.TitleDe, news.TitleEn, news.Slug, news.Date, news.Public, news.GistDe.String, news.GistEn.String, news.ContentDe.String, news.ContentEn.String)

	if err != nil {
		return nil, err
	}

	return FindNewsBySlug(news.Slug)
}

func UpdateNews(news News) (*News, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	_, err = db.Exec("UPDATE \"news\" SET title_de = $1, title_en = $2, \"date\" = $3, \"public\" = $4, gist_de = $5, gist_en = $6, content_de = $7, content_en = $8 WHERE id = $9", news.TitleDe, news.TitleEn, news.Date, news.Public, news.GistDe.String, news.GistEn.String, news.ContentDe.String, news.ContentEn.String, news.Id)

	if err != nil {
		return nil, err
	}

	return FindNewsBySlug(news.Slug)
}

func SetNewsTags(newsId string, tags []Tag) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM \"news_tag\" WHERE news_id = $1", newsId)
	if err != nil {
		return tx.Rollback()
	}

	for _, tag := range tags {
		_, err = tx.Exec("INSERT INTO \"news_tag\" (tag_id, news_id) VALUES ($1, $2)", tag.Id, newsId)

		if err != nil {
			return tx.Rollback()
		}
	}

	return tx.Commit()
}

func DeleteNewsBySlug(slug string) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	news, err := FindNewsBySlug(slug)
	if err != nil {
		return tx.Rollback()
	}

	_, err = tx.Exec("DELETE FROM \"news_tag\" WHERE news_id = $1", news.Id)
	if err != nil {
		return tx.Rollback()
	}

	_, err = tx.Exec("DELETE FROM \"news\" WHERE id = $1", news.Id)
	if err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}
