package models

import (
	"database/sql"
	"ruehmkorf.com/database"
	"time"
)

// language=sql
const CreateDownloadTable = `
CREATE TABLE download (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name_de text NOT NULL,
    name_en text NOT NULL,
    slug text NOT NULL UNIQUE,
    date time NOT NULL,
    self_destruct boolean DEFAULT false,
    self_destruct_days int NULL,
    public boolean NOT NULL DEFAULT false,
    description_de text NULL,
    description_en text NULL,
    type text NOT NULL 
)
`

type Download struct {
	Id               string
	NameDe           string `db:"name_de"`
	NameEn           string `db:"name_en"`
	Slug             string
	Date             time.Time
	SelfDestruct     bool          `db:"self_destruct"`
	SelfDestructDays sql.NullInt32 `db:"self_destruct_days"`
	Public           bool
	DescriptionDe    sql.NullString `db:"description_de"`
	DescriptionEn    sql.NullString `db:"description_en"`
	Type             string
}

const DownloadFilePath = "./data/public/download/file/"
const DownloadPreviewImagePath = "./data/public/download/preview/"

func FindAllDownloads(offset int, limit int) ([]Download, int, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, 0, err
	}

	defer db.Close()
	news := new([]Download)

	if err = db.Select(news, "SELECT * FROM \"download\" ORDER BY slug LIMIT $1 OFFSET $2", limit, offset); err != nil {
		return nil, 0, err
	}

	var totalCount int
	if err = db.Get(&totalCount, "SELECT COUNT(*) FROM \"download\""); err != nil {
		return *news, len(*news), err
	}

	return *news, totalCount, nil
}

func FindAllDownloadsToSelfDestruct() ([]Download, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	news := new([]Download)

	if err = db.Select(news, "SELECT * FROM \"download\" WHERE \"date\" + interval '1 day' *  self_destruct_days < CURRENT_TIMESTAMP AND self_destruct = true ORDER BY slug"); err != nil {
		return nil, err
	}

	return *news, nil
}

func FindDownloadBySlug(slug string) (*Download, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	download := new(Download)

	if err = db.Get(download, "SELECT * FROM \"download\" WHERE slug = $1", slug); err != nil {
		return nil, err
	}

	return download, nil
}

func CreateDownload(download Download) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO download (name_de, name_en, slug, \"date\", self_destruct, self_destruct_days, \"public\", description_de, description_en, type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", download.NameDe, download.NameEn, download.Slug, download.Date, download.SelfDestruct, download.SelfDestructDays.Int32, download.Public, download.DescriptionDe.String, download.DescriptionDe.String, download.Type)

	return err
}

func UpdateDownload(download Download) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("UPDATE download SET name_de = $1, name_en = $2, slug = $3, \"date\" = $4, self_destruct = $5, self_destruct_days = $6, \"public\" = $7, description_de = $8, description_en = $9, type = $10 WHERE id = $11", download.NameDe, download.NameEn, download.Slug, download.Date, download.SelfDestruct, download.SelfDestructDays.Int32, download.Public, download.DescriptionDe.String, download.DescriptionDe.String, download.Type, download.Id)

	return err
}

func DeleteDownloadBySlug(slug string) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec("DELETE FROM download WHERE slug = $1", slug)

	return err
}
