package repository

import "github.com/krabhinav07/homefixit/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
