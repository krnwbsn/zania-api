// Code generated by candi v1.18.0.

package usecase

import (
	"context"

	"zania-api/internal/modules/data-management/domain"
	"zania-api/pkg/shared/repository"
	"zania-api/pkg/shared/usecase/common"

	"github.com/golangid/candi/codebase/factory/dependency"
)

// DataManagementUsecase abstraction
type DataManagementUsecase interface {
	GetAllDataManagement(ctx context.Context, filter *domain.FilterDataManagement) (data domain.ResponseDataManagementList, err error)
}

type dataManagementUsecaseImpl struct {
	deps          dependency.Dependency
	sharedUsecase common.Usecase
	repoSQL       repository.RepoSQL
}

// NewDataManagementUsecase usecase impl constructor
func NewDataManagementUsecase(deps dependency.Dependency) (DataManagementUsecase, func(sharedUsecase common.Usecase)) {
	uc := &dataManagementUsecaseImpl{
		deps:    deps,
		repoSQL: repository.GetSharedRepoSQL(),
	}
	return uc, func(sharedUsecase common.Usecase) {
		uc.sharedUsecase = sharedUsecase
	}
}