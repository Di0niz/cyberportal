package github.com/Di0niz/cyber-backend/database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// определяем список баз данных для подключения
type DBPool struct {
	Connections []*DatabaseParam
}

// параметр подключения к базе данных
type DatabaseParam struct {
	DB       *sql.DB
	IsMaster bool
	IsLive   bool
	Dns      string
}

func InitializeDatabase() *DBPool {

	dnslist := []string{
		"root:example@tcp(192.168.101.11:3306)/cybergame?&charset=utf8&interpolateParams=true",
		"root:example@tcp(192.168.101.21:3306)/cybergame?&charset=utf8&interpolateParams=true",
		"root:example@tcp(192.168.101.12:3306)/cybergame?&charset=utf8&interpolateParams=true",
		"root:example@tcp(192.168.101.22:3306)/cybergame?&charset=utf8&interpolateParams=true"}

	pool := &DBPool{}

	IsMaster := true
	for _, dns := range dnslist {

		db, err := sql.Open("mysql", dns)
		if err != nil {
			db.SetMaxOpenConns(20)
		}
		param := &DatabaseParam{
			DB:       db,
			IsMaster: IsMaster,
			IsLive:   db.Ping() == nil,
			Dns:      dns,
		}

		pool.Connections = append(pool.Connections, param)

		IsMaster = !IsMaster
	}

	return pool
}
