package justgo

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)
var loadedDb *sqlx.DB
var Storage storage

type DB func() *sqlx.DB
type storage struct {
	db *sqlx.DB
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
		Log.Info("connecting to db")
		s.db = newDB(s.ConfiguredDriver, s.ConnectionString, int(s.ConnectionPool))
	} else {
		Log.WithField("config", s).Warn("Storage is not initialized")
	}
}

func (s *storage) Get() *sqlx.DB{
	if s.db == nil {
		Log.WithField("config", s).Panic("Storage is not loaded")
	}
	return s.db
}

func newDB(driver, connectionString string, poolSize int) *sqlx.DB {

	db, err := sqlx.Open(driver, connectionString)
	if err != nil {
		Log.WithField("error", err).
			WithField("driver", driver).
			Fatal("Error while connecting to database")
	}

	if err = db.Ping(); err != nil {
		Log.WithField("error", err).
			WithField("driver", driver).
			Fatal("ping to the database host failed")
	}

	db.SetMaxOpenConns(poolSize)
	return db
}

func GetDB() *sqlx.DB{
	if loadedDb == nil {
		Log.Panic("db is not loaded")
	}
	return loadedDb
}
