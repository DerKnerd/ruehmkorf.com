package migrations

import (
	"ruehmkorf.com/database"
)

type MissingForeignKeys20210530T0330 struct{}

func (migration MissingForeignKeys20210530T0330) Execute() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx := db.MustBegin()
	tx.MustExec("ALTER TABLE news_tag ADD CONSTRAINT news_tag_tag_id_fk FOREIGN KEY (tag_id) REFERENCES tag ON DELETE CASCADE")
	tx.MustExec("ALTER TABLE news_tag ADD CONSTRAINT news_tag_news_id_fk FOREIGN KEY (news_id) REFERENCES news ON DELETE CASCADE")
	if err = tx.Commit(); err != nil {
		panic(err)
	}
}

func (migration MissingForeignKeys20210530T0330) GetVersion() string {
	return "MissingForeignKeys20210530T0330"
}
