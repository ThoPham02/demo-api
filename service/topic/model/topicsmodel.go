package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TopicsModel = (*customTopicsModel)(nil)

type (
	// TopicsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTopicsModel.
	TopicsModel interface {
		topicsModel
	}

	customTopicsModel struct {
		*defaultTopicsModel
	}
)

// NewTopicsModel returns a model for the database table.
func NewTopicsModel(conn sqlx.SqlConn) TopicsModel {
	return &customTopicsModel{
		defaultTopicsModel: newTopicsModel(conn),
	}
}
