package initialize

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/phongnd2802/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)


func InitMysql() {
	m := global.Config.MySql

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.User, m.Password, m.Host, m.Port, m.DB)
	//fmt.Println(s)
	db, _ := sql.Open("mysql", s)
	err := db.Ping()
	checkErrorPanic(err, "Connect database failed")
	global.Logger.Info("Connect database success")
	global.Mdb = db
	//fmt.Println(global.Mdb)
	setPool()
}

func checkErrorPanic(err error, errString string) {
	//fmt.Println("Error:::>>")
	if err != nil {
		fmt.Println("Error::>>>")
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func setPool() {
	m := global.Config.MySql

	global.Mdb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns) * time.Second)
	global.Mdb.SetMaxOpenConns(m.MaxOpenConns)
	global.Mdb.SetConnMaxLifetime(time.Duration(m.ConnMaxLife) * time.Second)
}