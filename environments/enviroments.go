package environments

import "os"

type Env struct{}

func NewEnv() *Env {
	return &Env{}
}

func (e *Env) HostName() string {
	return os.Getenv("HOSTNAME")
}

func (e *Env) Port() string {
	return os.Getenv("PORT")
}

func (e *Env) DbUser() string {
	return os.Getenv("DATABASE_USER")
}

func (e *Env) DbPass() string {
	return os.Getenv("DATABASE_PASS")
}

func (e *Env) DbHost() string {
	return os.Getenv("DATABASE_HOST")
}

func (e *Env) DbPort() string {
	return os.Getenv("DATABASE_PORT")
}

func (e *Env) DbSchema() string {
	return os.Getenv("DATABASE_SCHEMA")
}

func (e *Env) DbCharset() string {
	return os.Getenv("DATABASE_CHARSET")
}
