package leodb

import (
	"github.com/cqlsy/leoconfig"
	"github.com/cqlsy/leodb/config"
	"github.com/cqlsy/leodb/mog"
	"github.com/cqlsy/leodb/sql"
	"strings"
)

type Db struct {
	MogDb *mog.MongoDb
	Mysql *sql.MysqlDB
}

func InitDataBase(path string) *Db {
	Conf := new(config.Info)
	leoconfig.ParseConf(path, &Conf)
	if Conf == nil {
		panic("Please finish config set")
	}
	db := new(Db)
	if strings.ToUpper(Conf.Db.Protocol) == "MONGODB" {
		mon, err := mog.Connect(Conf)
		if err != nil {
			panic("connect mongo database err： " + err.Error())
		}
		db.MogDb = mon
	} else if strings.ToUpper(Conf.Db.Protocol) == "MYSQL" {

	}
	return db
}
