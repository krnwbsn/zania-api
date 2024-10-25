// Code generated by candi v1.18.0. DO NOT EDIT.

package repository

import (
	"sync"
	"github.com/golangid/candi/codebase/factory/dependency"
)

var (
	once sync.Once
)

func SetSharedRepository(deps dependency.Dependency) {
	once.Do(func() {
		setSharedRepoSQL(deps.GetSQLDatabase().ReadDB(), deps.GetSQLDatabase().WriteDB())
	})
}
