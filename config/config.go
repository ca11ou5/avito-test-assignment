package config

type Config struct {
	ServerAddress    string `env:"SERVER_ADDRESS"`
	PostgresConn     string `env:"POSTGRES_CONN"`
	PostgresJDBCURL  string `env:"POSTGRES_JDBC_URL"`
	PostgresUsername string `env:"POSTGRES_USERNAME"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     string `env:"POSTGRES_PORT"`
	PostgresDatabase string `env:"POSTGRES_DATABASE"`
}
