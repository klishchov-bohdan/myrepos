package services

import "github.com/klishchov-bohdan/myrepos/internal/repositories/filesystem"

type Service struct {
	User *filesystem.UserFileRepository
}

func New(UFRepo *filesystem.UserFileRepository) *Service {
	return &Service{
		User: UFRepo,
	}
}
