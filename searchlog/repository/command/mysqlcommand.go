package command

import (
	"context"
	"database/sql"
	"time"

	"github.com/andrefebrianto/Search-Movie-Service/searchlog"
)

type MysqlCommandRepository struct {
	db *sql.DB
}

func CreateCassandraCommandRepository(mysqlClient *sql.DB) MysqlCommandRepository {
	return MysqlCommandRepository{db: mysqlClient}
}

func (repo MysqlCommandRepository) Create(ctx context.Context, searchLog *searchlog.SearchLog) error {
	queryStatement := `INSERT  searchlog SET keyword=? , page=? , timestamp=?`
	statement, err := repo.db.PrepareContext(ctx, queryStatement)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.ExecContext(ctx, searchLog.KeyWord, searchLog.Page, time.Now().Local())
	if err != nil {
		return err
	}

	return nil
}
