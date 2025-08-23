package database

import (
	"context"
	"os"

	"github.com/DerKnerd/gorp"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

var dbMap *gorp.DbMap

func GetDbMap() *gorp.DbMap {
	return dbMap
}

func SetupDatabase() {
	if dbMap == nil {
		pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			panic(err)
		}

		conn := stdlib.OpenDBFromPool(pool)

		dialect := gorp.PostgresDialect{}

		dbMap = &gorp.DbMap{Db: conn, Dialect: dialect}

		AddTableWithName[User]("user")
		AddTableWithName[Token]("token")
		AddTableWithName[Profile]("profile")
		AddTableWithName[SpotMapping]("spot_mapping")

		err = dbMap.CreateTablesIfNotExists()
		if err != nil {
			panic(err)
		}

		_, err = conn.Exec(`
alter table token
	drop constraint if exists token_user_fkey;
alter table token
	add constraint token_user_fkey foreign key (user_id) references "user" (id);
`)

		if err != nil {
			panic(err)
		}
	}
}
