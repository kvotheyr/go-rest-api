package db

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	mysqlmigration "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go-rest-api/pkg/env"
)

var (
	db                *sqlx.DB
	dbServiceName     = env.GetString("MYSQL_SERVICE_NAME", "mysql_db")
	dbMaxOpenConns    = env.GetInt("MYSQL_SERVICE_MAX_OPEN_CONNS", 10)
	dbMaxIdleConns    = env.GetInt("MYSQL_SERVICE_MAX_IDLE_CONNS", 2)
	dbConnMaxLifetime = env.GetInt("MYSQL_SERVICE_CONN_MAX_LIFETIME_IN_MIN", 5)
	dbConnMaxIdleTime = env.GetInt("MYSQL_SERVICE_CONN_MAX_IDLE_TIME_IN_SEC", 2)
)

func GetDB() *sqlx.DB {
	return db
}

func connectionString() (string, error) {
	config := mysql.Config{
		User:                 "test",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", dbServiceName, 3306),
		DBName:               "rest-db",
		Collation:            "utf8_general_ci",
		MultiStatements:      true,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	return config.FormatDSN(), nil
}

func Init() error {
	dsn, err := connectionString()
	if err != nil {
		return errors.Wrap(err, "Failed to build dsn")
	}

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return errors.Wrap(err, "Failed to initialize database")
	}

	db.SetMaxOpenConns(dbMaxOpenConns)
	db.SetMaxIdleConns(dbMaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(dbConnMaxLifetime) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(dbConnMaxIdleTime) * time.Second)

	migratioDirName := "file:///app/migrations"

	driver, err := mysqlmigration.WithInstance(db.DB, &mysqlmigration.Config{})
	if err != nil {
		return errors.Wrap(err, "Failed to initialize migrations driver")
	}

	m, err := migrate.NewWithDatabaseInstance(
		migratioDirName,
		"mysql",
		driver,
	)

	if err != nil {
		return errors.Wrap(err, "Failed to access migrations directory")
	}

	err = m.Up()
	if err == migrate.ErrNoChange {
		return nil
	}

	return errors.Wrap(err, "Failed to run migrations")
}
