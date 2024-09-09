package database

import (
	"context"
	"database/sql"
	"fmt"
	"github/devdahcoder/fiber-authentication/internal/config/env"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type DbStruct struct {
	DBUser string
	DBPassword string
	DBName string
	DBHost string
	DBPort string
}

func NewDbStruct(user, password, dbName, host, port string) DbStruct {
	return DbStruct{
		DBUser:     user,
		DBPassword: password,
		DBName:     dbName,
		DBHost:     host,
		DBPort:     port,
	}
}

func (db *DbStruct) GetDBUser() string {
	return db.DBUser
}

func (db *DbStruct) GetDBPassword() string {
	return db.DBPassword
}

func (db *DbStruct) GetDBName() string {
	return db.DBName
}

func (db *DbStruct) GetDBHost() string {
	return db.DBHost
}

func (db *DbStruct) GetDBPort() string {
	return db.DBPort
}

func (db *DbStruct) GetDBConnectionDsn() (string, error) {
	if db == nil || db.DBUser == "" || db.DBPassword == "" || db.DBName == "" || db.DBHost == "" || db.DBPort == "" {
		return "", fmt.Errorf("dbstruct has not been initialized or is missing values")
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db.GetDBUser(), db.GetDBPassword(), db.GetDBHost(), db.GetDBPort(), db.GetDBName()), nil
}

func PingDb(db *sql.DB, dbName string) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	err := db.PingContext(ctx)

	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return
	}
	
	log.Printf("Connected to DB %s successfully\n", dbName)
}

func DatabaseConfig() {
	db := ConnectDB()
	boil.SetDB(db)
}

func ConnectDB() *sql.DB {

	dbValue := NewDbStruct(env.EnvConfig("DB_USER"), env.EnvConfig("DB_PASSWORD"), env.EnvConfig("DB_NAME"), env.EnvConfig("DB_HOST"), env.EnvConfig("DB_PORT"))

	dsn, err := dbValue.GetDBConnectionDsn()

	if err != nil {
		log.Fatal(err)
	}

    db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)
	PingDb(db, dbValue.GetDBName())

	return db
}

