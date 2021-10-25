package command_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func CreateSqlMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	// db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCreate(t *testing.T) {
	// mockDb, mockSql := CreateSqlMock()
	// queryStatement := regexp.QuoteMeta(`INSERT searchlog SET url=?, response_data=?, status=?, timestamp=?`)
	// searchLog := &searchlog.SearchLog{
	// 	Url:          "https://youtube.com",
	// 	ResponseData: "Dummy Text",
	// 	Status:       200,
	// 	Timestamp:    time.Now().Local(),
	// }

	// // t.Run("should success to create/insert new row", func(t *testing.T) {
	// // 	prepareStatement := mockSql.ExpectPrepare(queryStatement)
	// // 	commandRepo := command.CreateMySqlCommandRepository(mockDb)

	// // 	prepareStatement.ExpectExec().WithArgs(searchLog.Url, searchLog.ResponseData, searchLog.Status, searchLog.Timestamp)

	// // 	err := commandRepo.Create(context.TODO(), searchLog)
	// // 	assert.Nil(t, err)
	// // })
}
