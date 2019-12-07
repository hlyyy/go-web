//数据库的连接

package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func (db *Database) Init() {
	newdb,err := gorm.Open("mysql","root:huanglingyun0130@/testdb?charset=utf8&parseTime=True&loc=local")
	if err != nil {
		fmt.Println(err)
	}
	DB = &Database{
		Self: newdb,
	}
}

func (db *Database) Close() {
	DB.Self.Close()
}


