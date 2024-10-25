// Code generated by candi v1.18.0.

package repository

import (
	"context"
	
	"time"
	"strings"

	"zania-api/internal/modules/data-management/domain"
	shareddomain "zania-api/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"

	"zania-api/pkg/shared"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type dataManagementRepoSQL struct {
	readDB, writeDB *gorm.DB
	updateTools     *candishared.DBUpdateTools
}

// NewDataManagementRepoSQL mongo repo constructor
func NewDataManagementRepoSQL(readDB, writeDB *gorm.DB) DataManagementRepository {
	return &dataManagementRepoSQL{
		readDB: readDB, writeDB: writeDB,
		updateTools: &candishared.DBUpdateTools{
			KeyExtractorFunc: candishared.DBUpdateGORMExtractorKey, IgnoredFields: []string{"id"},
		},
	}
}

func (r *dataManagementRepoSQL) FetchAll(ctx context.Context, filter *domain.FilterDataManagement) (data []shareddomain.DataManagement, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DataManagementRepoSQL:FetchAll")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()

	if filter.OrderBy == "" {
		filter.OrderBy = "updated_at"
	}

	db := r.setFilterDataManagement(shared.SetSpanToGorm(ctx, r.readDB), filter).Order(clause.OrderByColumn{
		Column: clause.Column{Name: filter.OrderBy},
		Desc:   strings.ToUpper(filter.Sort) == "DESC",
	})
	if filter.Limit > 0 || !filter.ShowAll {
		db = db.Limit(filter.Limit).Offset(filter.CalculateOffset())
	}
	err = db.Find(&data).Error
	return
}

func (r *dataManagementRepoSQL) Count(ctx context.Context, filter *domain.FilterDataManagement) (count int) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DataManagementRepoSQL:Count")
	defer trace.Finish()

	var total int64
	r.setFilterDataManagement(shared.SetSpanToGorm(ctx, r.readDB), filter).Model(&shareddomain.DataManagement{}).Count(&total)
	count = int(total)
	
	trace.Log("count", count)
	return
}

func (r *dataManagementRepoSQL) Find(ctx context.Context, filter *domain.FilterDataManagement) (result shareddomain.DataManagement, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DataManagementRepoSQL:Find")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()

	err = r.setFilterDataManagement(shared.SetSpanToGorm(ctx, r.readDB), filter).First(&result).Error
	return
}

func (r *dataManagementRepoSQL) Save(ctx context.Context, data *shareddomain.DataManagement, updateOptions ...candishared.DBUpdateOptionFunc) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DataManagementRepoSQL:Save")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()

	db := r.writeDB
	if tx, ok := candishared.GetValueFromContext(ctx, candishared.ContextKeySQLTransaction).(*gorm.DB); ok {
		db = tx
	}
	data.UpdatedAt = time.Now()
	if data.CreatedAt.IsZero() {
		data.CreatedAt = time.Now()
	}
	if data.ID == 0 {
		err = shared.SetSpanToGorm(ctx, db).Omit(clause.Associations).Create(data).Error
	} else {
		err = shared.SetSpanToGorm(ctx, db).Model(data).Omit(clause.Associations).Updates(r.updateTools.ToMap(data, updateOptions...)).Error
	}
	return
}

func (r *dataManagementRepoSQL) Delete(ctx context.Context, filter *domain.FilterDataManagement) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DataManagementRepoSQL:Delete")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()

	db := r.writeDB
	if tx, ok := candishared.GetValueFromContext(ctx, candishared.ContextKeySQLTransaction).(*gorm.DB); ok {
		db = tx
	}
	err = r.setFilterDataManagement(shared.SetSpanToGorm(ctx, db), filter).Delete(&shareddomain.DataManagement{}).Error
	return
}

func (r *dataManagementRepoSQL) setFilterDataManagement(db *gorm.DB, filter *domain.FilterDataManagement) *gorm.DB {
	
	if filter.ID != nil {
		db = db.Where("id = ?", *filter.ID)
	}
	if filter.Search != "" {
		db = db.Where("(field ILIKE '%%' || ? || '%%')", filter.Search)
	}

	for _, preload := range filter.Preloads {
		db = db.Preload(preload)
	}

	return db
}