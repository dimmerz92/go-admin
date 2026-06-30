package configs

type DatabaseDriver = string

type DatabaseConfig struct {
	Driver DatabaseDriver
	DSN    string
}
