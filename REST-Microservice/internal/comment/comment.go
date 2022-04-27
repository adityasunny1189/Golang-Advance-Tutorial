package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - a representation of the comment
// structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface definces all of of
// the methods that our service need in order
// to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - is the struct on which all our
// logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to the new
// service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving the comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctc context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctc context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
