package service

import (
	"github.com/krabhinav07/homefixit/internal/lib/job"
	"github.com/krabhinav07/homefixit/internal/repository"
	"github.com/krabhinav07/homefixit/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
