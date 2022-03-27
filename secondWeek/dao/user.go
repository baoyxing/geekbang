package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

type UserDao struct {
}

var (
	DB sql.DB
)

func initDB() error {
	DB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return errors.Wrap(err, "链接mysql失败")
	} else {
		DB.SetConnMaxLifetime(100)
		DB.SetMaxIdleConns(10)
	}

	return err
}

func (dao *UserDao) ISRegisterWithAccount(account string) (bool, error) {
	initDB()
	querySql := "select * from user where account = " + account
	_, err := DB.Query(querySql)
	if errors.Is(err, sql.ErrNoRows) {
		return true, nil
	} else {
		return false, errors.Wrap(err, "查询用户注册失败")
	}
}
