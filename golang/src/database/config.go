package main
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)

type DBPool struct{
	Connections [4]DatabaseConnection
}

type DatabaseParam struct {
	DB *sql.DB
	IsMaster bool
	IsLive	 bool
	Dns		 string
}

func InitializeDatabase() * DBPool {

	dnslist := []string{
		"root:example@tcp(192.168.101.11:3306)/cybergame?&charset=utf8&interpolateParams=true",
		"root:example@tcp(192.168.101.21:3306)/cybergame?&charset=utf8&interpolateParams=true",
		"root:example@tcp(192.168.101.12:3306)/cybergame?&charset=utf8&interpolateParams=true",
		"root:example@tcp(192.168.101.22:3306)/cybergame?&charset=utf8&interpolateParams=true"}

	pool := &DBPool{}

	IsMaster := true
	for _, dns:= range dnslist {
		db, err := sql.Open("mysql", dsn)

		db.SetMaxOpenConns(10)

		

		param := &DatabaseParam {
			DB: 		db, 
			IsMaster: 	IsMaster,
			IsLive:	 	db.Ping() == nil,
			Dns	:	 	dns,
		}

		append(pool.Connections, )

		IsMaster = !IsMaster


	}

	return &Database{[
		{
			DB: 
		}, 
	]}


}