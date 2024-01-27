package domain

type DBSettings struct {
	Host               string
	User               string
	Password           string
	DBName             string
	Catalog            string
	SSL                string
	Schema             string
	TimeZone           string
	Port               string
	DisableAutoMigrate bool
}
