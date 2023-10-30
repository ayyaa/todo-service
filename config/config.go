package config

import (
	"fmt"

	"os"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

type Config struct {
	DB  Database
	AWS AWSCredential
}

type Database struct {
	ConnectionString string
}

type AWSCredential struct {
	AccesKey   string
	KeyID      string
	Region     string
	BucketName string
	AWSSession *s3.S3
	AwsFlag    bool
}

// Load all config for the system
func LoadConfig() Config {

	// Load config value from file .env for local configuration
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("failed to load local config")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	accesKey := os.Getenv("AWS_ACCESS_KEY")
	keyID := os.Getenv("AWS_KEY_ID")
	bucketName := os.Getenv("AWS_BUCKET_NAME")
	region := os.Getenv("AWS_REGION")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbName, port)
	cfg := Config{
		DB: Database{
			ConnectionString: dsn,
		},
		AWS: AWSCredential{
			AccesKey:   accesKey,
			KeyID:      keyID,
			BucketName: bucketName,
			Region:     region,
			AwsFlag:    true,
		},
	}

	return cfg
}
