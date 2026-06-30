package auth

import (
	"context"
)

type Service struct {
	Repo *Repository
}
type ServiceDeps struct {
	Repo *Repository
}

func NewService(deps ServiceDeps) *Service {
	return &Service{
		Repo: deps.Repo,
	}

}

func (s *Service) List(ctx context.Context, query ListTodoReq) (*ListTodoRes, error) {
	res, err := s.Repo.QueryTodo(ctx, query)
	if err != nil {
		return nil, err
	}
	return res, nil
}
