package service

import (
	"fmt"
	"time"

	"gitlab.com/zapirus/task/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s *Service) Take(times time.Time, auth string, headers, body []byte) error {
	if err := s.repo.TakeData(times, auth, headers, body); err != nil {
		return fmt.Errorf("error inserting URL: %s", err)
	}
	return nil
}
