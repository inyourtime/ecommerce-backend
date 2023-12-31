package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Cfg *Env

type Env struct {
	App            *App
	DiscordWebhook *DiscordWebhook
	Db             *Db
	S3             *S3
	Jwt            *Jwt
	Google         *Google
}

func LoadConfig() *Env {
	appEnv := os.Getenv("APP_ENV")

	if appEnv != "prod" && appEnv != "dev" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	Cfg = &Env{
		App: &App{
			ServerPort: os.Getenv("CONFIG_SERVERPORT"),
			ApiPrefix:  os.Getenv("CONFIG_API_PREFIX"),
		},
		DiscordWebhook: &DiscordWebhook{
			ID:    os.Getenv("DISCORD_ID"),
			Token: os.Getenv("DISCORD_TOKEN"),
		},
		Db: &Db{
			Mongo: &Mongo{
				Uri:      os.Getenv("MONGODB_URL"),
				Database: os.Getenv("MONGODB_DATABASE"),
			},
		},
		S3: &S3{
			AccountID:       os.Getenv("OBJECTSTORAGE_ACCOUNTID"),
			AccessKeyID:     os.Getenv("OBJECTSTORAGE_ACCESSKEYID"),
			AccessKeySecret: os.Getenv("OBJECTSTORAGE_SECRETACCESSKEY"),
			Bucket:          os.Getenv("OBJECTSTORAGE_BUCKET"),
		},
		Jwt: &Jwt{
			Secret: os.Getenv("JWT_SECRETKEY"),
		},
		Google: &Google{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		},
	}
	return Cfg
}

type App struct {
	ServerPort string
	ApiPrefix  string
}

type DiscordWebhook struct {
	ID    string
	Token string
}

type Db struct {
	Mongo *Mongo
}

type Mongo struct {
	Uri      string
	Database string
}

type S3 struct {
	AccountID       string
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
}

type Jwt struct {
	Secret string
}

type Google struct {
	ClientID     string
	ClientSecret string
}
