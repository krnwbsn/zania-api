package usecase

import (
	"context"

	"zania-api/internal/modules/data-management/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *dataManagementUsecaseImpl) GetAllDataManagement(ctx context.Context, filter *domain.FilterDataManagement) (result domain.ResponseDataManagementList, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DataManagementUsecase:GetAllDataManagement")
	defer trace.Finish()

	data, err := uc.repoSQL.DataManagementRepo().FetchAll(ctx, filter)
	if err != nil {
		return result, err
	}
	count := uc.repoSQL.DataManagementRepo().Count(ctx, filter)
	result.Meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	result.Data = make([]domain.ResponseDataManagement, len(data))
	for i, detail := range data {
		result.Data[i].Serialize(&detail)
	}

	return
}
