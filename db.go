package justgo

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)
var defaultStorage storage

type DB struct {
	*sqlx.DB
}
type storage struct {
	db DB
	ConfiguredDriver,
	ConnectionString string
	ConnectionPool int64
}

func (s *storage) Load(){
	enableDB := false
	s.ConfiguredDriver = Config.GetStringOrDefault(ConfigKey.DB_DRIVER, "disable")
	s.ConnectionString = Config.GetStringOrDefault(ConfigKey.DB_CONNECTION_STRING, "")
	s.ConnectionPool = Config.GetIntOrDefault(ConfigKey.DB_CONNECTION_POOL, 100)
	for _, driver := range sql.Drivers() {
		if s.ConfiguredDriver == driver {
			enableDB = true
			break
		}
	}

	if enableDB && s.ConnectionString != ""{
		log.Info("connecting to DB")
		s.db = DB{newDB(s.ConfiguredDriver, s.ConnectionString, int(s.ConnectionPool))}
	} else {
		log.WithField("config", s).Warn("defaultStorage is not initialized")
	}
}

func newDB(driver, connectionString string, poolSize int) *sqlx.DB {

	db, err := sqlx.Open(driver, connectionString)
	if err != nil {
		log.WithField("error", err).
			WithField("driver", driver).
			Fatal("Error while connecting to database")
	}

	if err = db.Ping(); err != nil {
		log.WithField("error", err).
			WithField("driver", driver).
			Fatal("ping to the database host failed")
	}

	db.SetMaxOpenConns(poolSize)
	return db
}

func GetDB() *DB {
	return &defaultStorage.db
}
