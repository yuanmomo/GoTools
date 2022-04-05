package Sqlite

import (
	"database/sql"
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var (
	SQLiteCli *gorm.DB
)

func InitSqlite(dbname string, enableLog bool) {
	var err error
	config := &gorm.Config{}
	if enableLog {
		config.Logger = logger.Default.LogMode(logger.Info)
	} else {
		config.Logger = logger.Discard
	}
	SQLiteCli, err = gorm.Open(sqlite.Open(dbname), config)
	if err != nil {
		logrus.Error(err)
	}
	sqlDB, err := SQLiteCli.DB()
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func AutoMigrate(objs ...interface{}) error {
	return SQLiteCli.AutoMigrate(objs...)
}

func Create(obj interface{}) error {
	return SQLiteCli.Create(obj).Error
}

func SQLExec(sql string, args ...interface{}) (error, int64) {
	tx := SQLiteCli.Exec(sql, args...)
	return tx.Error, tx.RowsAffected
}

func SQLQuery(sqlstr string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := SQLiteCli.Raw(sqlstr, args...).Rows()
	if err != nil {
		fmt.Println("DbQuery", err)
		return nil, err
	}
	return SQLQueryRows(rows)
}

func SQLQueryRows(rows *sql.Rows) ([]map[string]interface{}, error) {
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	if err != nil {
		return nil, err
	}
	return tableData, nil
}
