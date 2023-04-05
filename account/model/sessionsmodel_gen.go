// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sessionsFieldNames          = builder.RawFieldNames(&Sessions{})
	sessionsRows                = strings.Join(sessionsFieldNames, ",")
	sessionsRowsExpectAutoSet   = strings.Join(stringx.Remove(sessionsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sessionsRowsWithPlaceHolder = strings.Join(stringx.Remove(sessionsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	sessionsModel interface {
		Insert(ctx context.Context, data *Sessions) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Sessions, error)
		Update(ctx context.Context, data *Sessions) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSessionsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Sessions struct {
		Id          int64  `db:"id"`
		UserName    string `db:"user_name"`
		RefeshToken string `db:"refesh_token"`
		UserAgent   string `db:"user_agent"`
		ClientIp    string `db:"client_ip"`
		IsBlocked   bool   `db:"is_blocked"`
		ExpiresAt   int64  `db:"expires_at"`
		CreatedAt   int64  `db:"created_at"`
	}
)

func newSessionsModel(conn sqlx.SqlConn) *defaultSessionsModel {
	return &defaultSessionsModel{
		conn:  conn,
		table: "`sessions`",
	}
}

func (m *defaultSessionsModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSessionsModel) FindOne(ctx context.Context, id int64) (*Sessions, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sessionsRows, m.table)
	var resp Sessions
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSessionsModel) Insert(ctx context.Context, data *Sessions) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, sessionsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserName, data.RefeshToken, data.UserAgent, data.ClientIp, data.IsBlocked, data.ExpiresAt)
	return ret, err
}

func (m *defaultSessionsModel) Update(ctx context.Context, data *Sessions) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sessionsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserName, data.RefeshToken, data.UserAgent, data.ClientIp, data.IsBlocked, data.ExpiresAt, data.Id)
	return err
}

func (m *defaultSessionsModel) tableName() string {
	return m.table
}
