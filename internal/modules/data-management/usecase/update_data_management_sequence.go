package usecase

import (
	"context"

	"zania-api/internal/modules/data-management/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *dataManagementUsecaseImpl) UpdateDataManagementSequence(ctx context.Context, sequences []domain.DataManagementSequence) error {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DataManagementUsecase:UpdateDataManagementSequence")
	defer trace.Finish()

	err := uc.repoSQL.DataManagementRepo().UpdateSequence(ctx, sequences)
	if err != nil {
		return err
	}

	return nil
}
