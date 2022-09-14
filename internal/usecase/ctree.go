package usecase

import (
	"context"

	"gctree/internal/entity"
	"gctree/internal/usecase/repo"
)

// UserUserCase -.
type CTreeUserCase struct {
	repo *repo.CTreeRepo
}

// New -.
func New(r *repo.CTreeRepo) *CTreeUserCase {
	return &CTreeUserCase{
		repo: r,
	}
}

// Save -.
func (uc *CTreeUserCase) Save(ctx context.Context, t entity.TreeNode) ([]string, int, error) {

	return uc.repo.Save(ctx, t)
}

// List -.
func (uc *CTreeUserCase) List(ctx context.Context, masterID, pid string) (entity.MetaNode, int, error) {

	return uc.repo.List(ctx, masterID, pid)
}

// Detail -.
func (uc *CTreeUserCase) Detail(ctx context.Context, appID string) (entity.TreeNode, int, error) {

	return uc.repo.Detail(ctx, appID)
}

// Delete -.
func (uc *CTreeUserCase) Delete(ctx context.Context, appIDList []string) (int, error) {

	return uc.repo.Delete(ctx, appIDList)
}
