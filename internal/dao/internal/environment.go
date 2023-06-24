// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EnvironmentDao is the data access object for table environment.
type EnvironmentDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns EnvironmentColumns // columns contains all the column names of Table for convenient usage.
}

// EnvironmentColumns defines and stores column names for table environment.
type EnvironmentColumns struct {
	Id         string //
	Name       string // 名称
	Describe   string // 描述
	Uri        string // 环境地址
	State      string // 状态
	Remarks    string // 备注
	CreateAt   string // 创建时间
	CreateUser string // 创建人
	UpdateAt   string // 更新时间
	UpdateUser string // 更新人
}

// environmentColumns holds the columns for table environment.
var environmentColumns = EnvironmentColumns{
	Id:         "id",
	Name:       "name",
	Describe:   "describe",
	Uri:        "uri",
	State:      "state",
	Remarks:    "remarks",
	CreateAt:   "create_at",
	CreateUser: "create_user",
	UpdateAt:   "update_at",
	UpdateUser: "update_user",
}

// NewEnvironmentDao creates and returns a new DAO object for table data access.
func NewEnvironmentDao() *EnvironmentDao {
	return &EnvironmentDao{
		group:   "default",
		table:   "environment",
		columns: environmentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EnvironmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EnvironmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EnvironmentDao) Columns() EnvironmentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EnvironmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EnvironmentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EnvironmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
