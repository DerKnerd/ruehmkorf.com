package models

// language=sql
var CreateTagTable = `
CREATE TABLE "tag" (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    tag text UNIQUE NOT NULL 
)
`

type Tag struct {
	Id  string
	Tag string
}
