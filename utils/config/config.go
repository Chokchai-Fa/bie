package config

import (
	"database/sql"
	"fastmath/utils/authentication"
	"fastmath/utils/dbpostgres"
	"fastmath/utils/email"

	"github.com/omise/omise-go"
	"golang.org/x/oauth2"
)

type AppConfig struct {
	DBPG         *dbpostgres.DBPG
	DBPGCli      *sql.DB
	WebUrl       string
	ChargeAmount int
	TokenMaker   authentication.Maker
	EmailServer  email.EmailServer
	FaceBookAuth *oauth2.Config
	OmiseClient  *omise.Client
}
