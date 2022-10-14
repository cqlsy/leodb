package leodb

import (
	conf "github.com/cqlsy/leodb/config"
	"github.com/cqlsy/leodb/mog"
	"github.com/cqlsy/leodb/sql"
	"github.com/cqlsy/leoutil"
	"strings"
)

type Db struct {
	MogDb *mog.MongoDb
	Mysql *sql.MysqlDB
}

func InitDataBase(path string) *Db {
	config := new(conf.DBConf)
	leoutil.ParseConf(path, &config)
	if config == nil {
		panic("Please finish config set")
	}
	db := new(Db)
	if strings.ToUpper(config.Protocol) == "MONGODB" {
		mon, err := mog.Connect(config)
		if err != nil {
			panic("connect mongo database err： " + err.Error())
		}
		db.MogDb = mon
	} else if strings.ToUpper(config.Protocol) == "MYSQL" {

	}
	return db
}
