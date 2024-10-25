// Code generated by candi v1.18.0.

package repository

import (
	"context"

	"zania-api/internal/modules/data-management/domain"
	shareddomain "zania-api/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
)

// DataManagementRepository abstract interface
type DataManagementRepository interface {
	FetchAll(ctx context.Context, filter *domain.FilterDataManagement) ([]shareddomain.DataManagement, error)
	Count(ctx context.Context, filter *domain.FilterDataManagement) int
	Find(ctx context.Context, filter *domain.FilterDataManagement) (shareddomain.DataManagement, error)
	Save(ctx context.Context, data *shareddomain.DataManagement, updateOptions ...candishared.DBUpdateOptionFunc) error
	Delete(ctx context.Context, filter *domain.FilterDataManagement) (err error)
}