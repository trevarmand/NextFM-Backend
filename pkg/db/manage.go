package db

// Account
// 816744749573
// KMS key ID
// 260e2020-5ba9-41e2-8d34-a3ae9aec2c15

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
	"github.com/trevarmand/nextfm-backend/pkg/util/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AwsPsqlConnection struct {
	dbConn *sql.DB
}

func GetConnection() *AwsPsqlConnection {
	dbEndpoint := os.Getenv("nextfm_db_endpoint")
	awsRegion := os.Getenv("AWS_REGION")
	dbUser := os.Getenv("nextfm_db_user")
	dbPass := os.Getenv("nextfm_db_pass")
	awsCreds := credentials.NewEnvCredentials()
	authToken, err := rdsutils.BuildAuthToken(dbEndpoint, awsRegion, dbUser, awsCreds)

	dnsStr := fmt.Sprintf("postgres://%s:%s@%s/nextfm", dbUser, url.PathEscape(authToken), dbEndpoint)
	fmt.Println(authToken, err, dnsStr)
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
