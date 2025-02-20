package db

type Driver string

const (
	SqliteDriver = "sqlite"
)

type SqliteConfig struct {
	DatabasePath string
}
