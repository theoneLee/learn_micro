package db

import (
	"database/sql"
	"user-srv/basic/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/util/log"
)

func initMysql() {
	var err error

	// 创建连接
	mysqlDB, err = sql.Open("mysql", config.GetMysqlConfig().GetURL())
	//todo 将url拆分成
	//fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true&loc=Asia%2FShanghai",
	//	config.GetMysqlConfig().GetUser(),
	//	config.GetMysqlConfig().GetPassword(),
	//	config.GetMysqlConfig().GetHost(),
	//	config.GetMysqlConfig().GetDbName(),
	//)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// 最大连接数
	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())

	// 最大闲置数
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	// 激活链接
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
}
