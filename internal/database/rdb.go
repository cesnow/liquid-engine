package database

import (
	"errors"
	"fmt"
	"github.com/cesnow/liquid-engine/logger"
	"github.com/cesnow/liquid-engine/settings"
	"github.com/jinzhu/gorm"

	"strings"
)

type RDB struct {
	IDatabase
	client map[string]*gorm.DB
	config *settings.RDBConf
}

func (rdb *RDB) GetClient(dbName string) (*gorm.DB, error) {
	dbName = strings.ToLower(dbName)
	if val, ok := rdb.client[dbName]; ok {
		return val, nil
	}
	return nil, errors.New(fmt.Sprintf("db name `%s` not connected", dbName))
}

func connect(conf *settings.RDBConf, dbName string) (db *gorm.DB, err error) {

	db, err = gorm.Open(
		conf.Dialects,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			conf.User,
			conf.Pass,
			conf.Host,
			dbName,
		))

	if err != nil {
		return nil, err
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.TablePrefix + defaultTableName
	}

	db.LogMode(true)
	// dbClient.SetLogger(logging.GetLogger())
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(conf.MaxIdleConn)
	db.DB().SetMaxOpenConns(conf.MaxOpenConn)

	return db, nil
}

func ConnectWithRDB(conf *settings.RDBConf) (rdb *RDB, err error) {

	rdb = &RDB{
		config: conf,
		client: make(map[string]*gorm.DB, len(conf.DbNames)),
	}

	for _, dbName := range conf.DbNames {
		dbName = strings.ToLower(strings.Trim(dbName, " "))
		dbConn, err := connect(conf, dbName)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Connect to database name `%s` failed, %s", dbName, err))
		}
		logger.SysLog.Infof("[RDS] Connected database name `%s`", dbName)
		rdb.client[dbName] = dbConn
	}

	return rdb, nil

}
