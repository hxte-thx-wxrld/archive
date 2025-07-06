package htwarchive

type ServerConfig struct {
	PostgresURL string
	S3Endpoint  string
}

func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		PostgresURL: "postgres://pguser:pgpassword@db:5432/pguser?search_path=public",
		S3Endpoint:  "http://127.0.0.1:8333",
	}
}
