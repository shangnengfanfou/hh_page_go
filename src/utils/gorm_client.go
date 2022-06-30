package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
)

/**
 * @Description:全局 DB
 */

var (
	DB *gorm.DB
)

/**
 * @Description: 初始化数据库
 * @return *gorm.DB
 */

func SqlClient(conf *ini.File) *gorm.DB {
	var err error

	// 数据库的类型
	dbType := conf.Section("").Key("db_type").String()

	// Mysql配置信息
	mysqlName := conf.Section("mysql").Key("db_name").String()
	mysqlUser := conf.Section("mysql").Key("db_user").String()
	mysqlPwd := conf.Section("mysql").Key("db_pwd").String()
	mysqlHost := conf.Section("mysql").Key("db_host").String()
	mysqlPort := conf.Section("mysql").Key("db_port").String()
	mysqlCharset := conf.Section("mysql").Key("db_charset").String()

	var dataSource string
	switch dbType {
	case "mysql":
		dataSource = mysqlUser + ":" + mysqlPwd + "@tcp(" + mysqlHost + ":" +
			mysqlPort + ")/" + mysqlName + "?charset=" + mysqlCharset

		DB, err = gorm.Open(dbType, dataSource)
	}

	if err != nil {
		DB.Close()
		// log.Fatal(consts.CONN_DATABASE_ERROR, err)
	}

	// 设置连接池，空闲连接
	DB.DB().SetMaxIdleConns(50)
	// 打开链接
	DB.DB().SetMaxOpenConns(100)

	// 表明禁用后缀加s
	DB.SingularTable(true)

	return DB
}
