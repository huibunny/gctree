// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"gctree/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// CTree -.
	CTree interface {
		Save(context.Context, entity.TreeNode) ([]string, int, error)
		List(context.Context, string, string) (entity.MetaNode, int, error)
		Detail(context.Context, string) (entity.TreeNode, int, error)
		Delete(context.Context, []string) (int, error)
	}

	// AppRepo -.
	AppRepo interface {
		Login(context.Context) error
	}
)
