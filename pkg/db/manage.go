package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/trevarmand/nextfm-backend/pkg/util/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AwsPsqlConnection struct {
	dbConn *sql.DB
}

func GetConnection() *AwsPsqlConnection {
	dbEndpoint := os.Getenv("nextfm_db_endpoint")
	dbUser := os.Getenv("nextfm_db_user")
	dbPass := os.Getenv("nextfm_db_pass")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", dbEndpoint, dbUser, dbPass, "nextfm")

	db, connectErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	log.LogError("util:sql:AwsPsql:GetConnection", "failed to open DB connection", connectErr)

	dbConn, err := db.DB()
	log.LogError("util:sql:AwsPsql:GetConnection", "Failed to initialize DB object from connection", err)

	pingFailure := dbConn.Ping()
	log.LogError("util:sql:AwsPsql:GetConnection", "Failed to ping db to verify connection", pingFailure)

	return &AwsPsqlConnection{
		dbConn: dbConn,
	}
}
