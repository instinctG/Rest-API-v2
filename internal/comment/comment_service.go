package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetching       = errors.New("cannot fetching a comment by id")
	ErrNotImplemented = errors.New("not implemented")
)

type Store interface {
	GetComment(ctx context.Context, id string) (Comment, error)
	PostComment(ctx context.Context, comment Comment) (Comment, error)
	UpdateComment(ctx context.Context, id string, comment Comment) (Comment, error)
	DeleteComment(ctx context.Context, id string) error
}

type Comment struct {
	ID     string
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
	fmt.Println("Retrieving a comment")

	comment, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetching
	}
	return comment, nil
}

func (s *Service) UpdateComment(ctx context.Context, ID string, updatedCmt Comment) (Comment, error) {

	cmt, err := s.Store.UpdateComment(ctx, ID, updatedCmt)
	if err != nil {
		fmt.Println("error updating comment")
		return Comment{}, err
	}

	return cmt, nil

}

func (s *Service) PostComment(ctx context.Context, comment Comment) (Comment, error) {
	cmt, err := s.Store.PostComment(ctx, comment)
	if err != nil {
		return Comment{}, fmt.Errorf("error to post comment: %w", err)
	}
	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {

	err := s.Store.DeleteComment(ctx, id)
	if err != nil {
		return ErrNotImplemented
	}
	return nil
}
