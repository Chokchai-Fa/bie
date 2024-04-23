package main

import (
	"fastmath/routers"
	"fastmath/utils/authentication"
	"fastmath/utils/config"
	"fastmath/utils/dbpostgres"
	"fastmath/utils/email"
	"fastmath/utils/logs"

	"os"
	"strconv"
	"time"

	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/omise/omise-go"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logs.Error("DotEnv|:" + err.Error())
	}

	env := os.Getenv("ENV")
	port := os.Getenv("PORT")

	logs.Info("ENV: " + env)
	logs.Info("PORT: " + port)
}

func connectDBPostgres() (*dbpostgres.DBPG, *sql.DB) {
	port, err := strconv.Atoi(os.Getenv("PG_PORT"))
	if err != nil {
		logs.Error("PostgresPort|:" + err.Error())
		os.Exit(1)
	}

	dbPG := dbpostgres.NewDBPostgres()
	dbPGCli, err := dbPG.Connect(
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_DBNAME"),
		port,
	)
	if err != nil {
		logs.Error("Postgres.Connect|:" + err.Error())
		os.Exit(1)
	}

	// set the maximum life time of a connection
	dbPGCli.SetConnMaxLifetime(5 * time.Minute)

	// set the maximum number of connections in the pool
	dbPGCli.SetMaxOpenConns(50)

	// set the maximum idle connections
	dbPGCli.SetMaxIdleConns(25)

	if err := dbPGCli.Ping(); err != nil {
		logs.Error("Postgres.Ping|:" + err.Error())
		os.Exit(1)
	}

	return dbPG, dbPGCli
}

func main() {
	port := os.Getenv("PORT")
	webUrl := os.Getenv("WEB_URL")

	dbPG, dbPGCli := connectDBPostgres()

	router := gin.New()
	router.Use(gin.Recovery())

	tokenDuration, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_DURATION"))
	if err != nil {
		logs.Error("AccessTokenDurationNumber|:" + err.Error())
		os.Exit(1)
	}

	refreshTokenDuration, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_DURATION"))
	if err != nil {
		logs.Error("RefreshTokenDurationNumber|:" + err.Error())
		os.Exit(1)
	}

	tokenMaker, err := authentication.NewJWTMaker(
		os.Getenv("TOKEN_SYMMETRIC_KEY"),
		tokenDuration,
		refreshTokenDuration,
	)
	if err != nil {
		logs.Error("TokenMaker|:" + err.Error())
		os.Exit(1)
	}

	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		logs.Error("SMTPPortNumber|:" + err.Error())
		os.Exit(1)
	}

	emailServer, err := email.NewEmailServer(
		os.Getenv("SERVER_EMAIL"),
		os.Getenv("SERVER_PASSWORD"),
		os.Getenv("SMTP_SERVER"),
		smtpPort,
	)
	if err != nil {
		logs.Error("MailServer|:" + err.Error())
		os.Exit(1)
	}

	faceBookAuth, err := authentication.NewFacebookAuth(
		os.Getenv("FACEBOOK_APP_ID"),
		os.Getenv("FACEBOOK_SECRET"),
		os.Getenv("FACEEBOOK_CALLBACK_URL"),
	)
	if err != nil {
		logs.Error("FacebookAuth|:" + err.Error())
		os.Exit(1)
	}

	omiseClient, err := omise.NewClient(
		os.Getenv("OMISE_PUBLIC_API_KEY"),
		os.Getenv("OMISE_PRIVATE_API_KEY"),
	)
	if err != nil {
		logs.Error("OmisePay|:" + err.Error())
		os.Exit(1)
	}

	chargeAmount, err := strconv.Atoi(os.Getenv("CHARGE_AMOUNT"))
	if err != nil {
		logs.Error("ChargeAmount|:" + err.Error())
		os.Exit(1)
	}

	appConfig := config.AppConfig{
		DBPG:         dbPG,
		DBPGCli:      dbPGCli,
		WebUrl:       webUrl,
		ChargeAmount: chargeAmount,
		TokenMaker:   tokenMaker,
		EmailServer:  emailServer,
		FaceBookAuth: faceBookAuth,
		OmiseClient:  omiseClient,
	}

	routers.SetupRouter(router, appConfig)

	if err := router.Run(":" + port); err != nil {
		logs.Error(err)
	}
}
