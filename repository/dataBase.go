package repository

import (
	"fmt"
	"log"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbAlias       = "default"
	mysqlUser     = "root"      //"testapp"
	mysqlPassword = "Mirror123" //"abc123"
	mysqlHost     = "localhost"
	mysqlPort     = 3306
	mysqlDatabase = "dbTest"
	mysqlCharset  = "utf8"
)

var (
	mysqlCon = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		mysqlUser,
		mysqlPassword,
		mysqlHost,
		mysqlPort,
		mysqlDatabase,
		mysqlCharset,
	)
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase(dbAlias, "mysql", mysqlCon)
	if err != nil {
		panic(err)
	}
}

//GetSession - conexion de Base de Datos
func GetSession() orm.Ormer {

	force := true   // Drop table and re-create.
	verbose := true // Print log
	if err := orm.RunSyncdb(dbAlias, force, verbose); err != nil {
		log.Println(err)
	}

	session := orm.NewOrm()
	session.Using(dbAlias)

	return session
}
