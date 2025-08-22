package database

type User struct {
	Id          int    `json:"id" db:"id,primarykey,autoincrement"`
	Username    string `json:"username" db:"username,unique,notnull"`
	Password    string `json:"-" db:"password,notnull"`
	TotpEnabled bool   `json:"totpEnabled" db:"totp_enabled,default:false"`
	TotpSecret  string `json:"-" db:"totp_secret"`
}

type Token struct {
	Id     int    `json:"id" db:"id,primarykey,autoincrement"`
	UserId int    `json:"userId" db:"user_id,notnull"`
	Token  string `json:"token" db:"token,notnull"`
}
