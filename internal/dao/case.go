// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"myapp/internal/dao/internal"
)

// internalCaseDao is internal type for wrapping internal DAO implements.
type internalCaseDao = *internal.CaseDao

// caseDao is the data access object for table case.
// You can define custom methods on it to extend its functionality as you wish.
type caseDao struct {
	internalCaseDao
}

var (
	// Case is globally public accessible object for table case operations.
	Case = caseDao{
		internal.NewCaseDao(),
	}
)

// Fill with you ideas below.
