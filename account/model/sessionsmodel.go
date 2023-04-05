package model

import (
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var _ SessionsModel = (*customSessionsModel)(nil)

var (
	sessionsFieldNames2          = builder.RawFieldNames(&Sessions{})
	sessionsRows2                = strings.Join(sessionsFieldNames, ",")
	sessionsRowsExpectAutoSet2   = strings.Join(stringx.Remove(sessionsFieldNames, "`created_at`"), ",")
	sessionsRowsWithPlaceHolder2 = strings.Join(stringx.Remove(sessionsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	// SessionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSessionsModel.
	SessionsModel interface {
		sessionsModel
	}

	customSessionsModel struct {
		*defaultSessionsModel
	}
)

// NewSessionsModel returns a model for the database table.
func NewSessionsModel(conn sqlx.SqlConn) SessionsModel {
	return &customSessionsModel{
		defaultSessionsModel: newSessionsModel(conn),
	}
}
