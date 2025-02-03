package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"golang_starter_template/pkg/utils"
	"strings"
)

type Database struct {
	db *sql.DB
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}

func (d *Database) Close() {
	err := d.GetDB().Close()
	if err != nil {
		utils.LoggerError.Printf(utils.Error+"Error closing database connection caused by: %s %s\n"+err.Error(), utils.Reset)
		return
	}
}

func Connect() (*Database, error) {
	dbDriver := utils.GetEnv("DB_DRIVER")
	dbUser := utils.GetEnv("DB_USER")
	dbPassword := utils.GetEnv("DB_PASSWORD")
	dbHost := utils.GetEnv("DB_HOST")
	dbPort := utils.GetEnv("DB_PORT")
	dbName := utils.GetEnv("DB_NAME")

	var (
		connString string
		driverName string
	)

	switch dbDriver {
	case "mysql":
		driverName = "mysql"
		connString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	case "postgres":
		driverName = "postgres"
		connString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	case "sqlite":
		driverName = "sqlite3"
		connString = fmt.Sprintf("./pkg/config/%s.db", strings.ToLower(utils.GetEnv("APP_NAME")))
	default:
		return nil, errors.New("unsupported database driver")
	}

	db, err := sql.Open(driverName, connString)
	if err != nil {
		return nil, errors.New("error connecting to database: " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		utils.LoggerError.Printf(utils.Error+"Error pinging database caused by: %s %s\n", err.Error(), utils.Reset)
		err = db.Close()
		if err != nil {
			utils.LoggerError.Println(utils.Error+"Error closing database connection caused by: %v %v", err.Error(), utils.Reset)
		}
		return nil, err
	}

	utils.LoggerInfo.Println(utils.Info + "Database connection established successfully!" + utils.Reset)
	return &Database{db: db}, nil
}

func Migrate(db *sql.DB) error {
	err := utils.Environment()
	if err != nil {
		return errors.New("error loading environment variables")
	}

	dbDriver := utils.GetEnv("DB_DRIVER")
	migrationPath := "file://pkg/config/migrations"

	var driver database.Driver

	switch dbDriver {
	case "mysql":
		driver, err = mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			return errors.New("error creating mysql driver: " + err.Error())
		}
	case "postgres":
		driver, err = postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return errors.New("error creating postgres driver: " + err.Error())
		}
	case "sqlite":
		driver, err = sqlite3.WithInstance(db, &sqlite3.Config{})
		if err != nil {
			return errors.New("error creating sqlite driver: " + err.Error())
		}
	default:
		return errors.New("unsupported database driver")
	}

	m, err := migrate.NewWithDatabaseInstance(migrationPath, dbDriver, driver)
	if err != nil {
		return errors.New("error creating migration instance: " + err.Error())
	}

	if err = m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.New("error rolling back migrations: " + err.Error())
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.New("error applying migrations: " + err.Error())
	}

	utils.LoggerInfo.Println(utils.Info + "Database migration successful!" + utils.Reset)
	return nil
}
