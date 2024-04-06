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
