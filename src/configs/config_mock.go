package configs

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "ecommerce-backend"

func InitConfigMock() {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	if err := godotenv.Load(string(rootPath) + `/.env`); err != nil {
		log.Fatal("Error loading .env file")
	}

	Cfg = &Env{
		DiscordWebhook: &DiscordWebhook{
			ID:    "1159508574125428776",
			Token: "peWVu2CWq7ZnLyZ4KCab_MGsw9-EmX6X4nA43FHrffk_ITAXqGDOA0Cwct3d8agenfsO",
		},
		Jwt: &Jwt{
			Secret: "secretkub",
		},
		S3: &S3{
			AccountID:       os.Getenv("OBJECTSTORAGE_ACCOUNTID"),
			AccessKeyID:     os.Getenv("OBJECTSTORAGE_ACCESSKEYID"),
			AccessKeySecret: os.Getenv("OBJECTSTORAGE_SECRETACCESSKEY"),
			Bucket:          "god-test",
		},
	}
}
