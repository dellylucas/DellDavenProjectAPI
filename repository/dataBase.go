package repository

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbAlias       = "default"
	mysqlUser     = "root"
	mysqlPassword = "Mirror123"
	mysqlHost     = "localhost"
	mysqlPort     = 3306
	mysqlDatabase = "dbprojectdeda"
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
	/*force := true   // Drop table and re-create.
	verbose := true // Print log
	_= orm.RunSyncdb(dbAlias, force, verbose)
	*/
	session := orm.NewOrm()
	session.Using(dbAlias)

	return session
}
