package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetching     = errors.New("cannot fetching a comment by id")
	ErrNotImplemented = errors.New("not implemented")
)

type Store interface {
	GetComment(ctx context.Context, id string) (Comment, error)
}

type Comment struct {
	ID     uint
	Slug   string
	Body   string
	Author string
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{Store: store}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetching
	}

	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, comment Comment) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context,comment Comment) (Comment,error)  {
	return Comment{},ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context,id string) error {
	return ErrNotImplemented
}
