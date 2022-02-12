package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...

type Config struct {
	Environment string // develop,staging,production

	LibraryServiceHost string
	LibraryServicePort int

	MinioAccessKeyID string
	MinioSecretKey   string
	MinioEndpoint    string
	MinioBucketName  string
	MinioLocation    string
	MinioHost        string

	LogLevel string
	HttpPort string
}

// var (
// 	PartisipantRoles = []string{"respondent", "professional", "consultant"}
// 	StudyTypes       = []string{"one_to_one", "unmoderated", "diary"}
// 	InterviewTypes   = []string{"online", "phone", "in_persion"}
// 	OperationSystems = []string{"linux", "windows", "macos"}
// 	Devices          = []string{"smartphone", "computer", "planshet"}
// )

// const (
// 	DraftStatusID    string = "72ee7350-4a2b-4fc2-945c-34faa56b13ca"
// 	ActiveStatusID   string = "7b75685f-ee20-4fb9-bc5f-7e40c115e708"
// 	FinishedStatusID string = "2e857fbe-9add-4eae-a8c4-fe57fb384347"
// 	CanceledStatusID string = "2a98c22e-cbee-49f9-90ef-f56429d96333"

// 	DefaultID string = "33f50eb1-f913-46d5-9b83-ce316c1fcc9d"

// 	NotApplyStatusID string = "caf9b5d0-3013-4c56-a3b0-b0c1972f721f"

// 	PendingStatusID   string = "01ee5c38-5f56-4e72-9b5a-926e4adce26a"
// 	ConfirmedStatusID string = "e86c9b5b-5c71-48c8-b61f-42eeb51e33c6"
// )

// Load loads environment vars and inflates Config

func Load() Config {
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8090"))

	config.LibraryServiceHost = cast.ToString(getOrReturnDefault("LIBRARY_SERVISE_HOST", "localhost"))
	config.LibraryServicePort = cast.ToInt(getOrReturnDefault("LIBRARY_SERVICE_PORT", 5005))

	config.MinioEndpoint = cast.ToString(getOrReturnDefault("MINIO_ENDPOINT", "test.cdn.urecruit.udevs.io"))
	config.MinioAccessKeyID = cast.ToString(getOrReturnDefault("MINIO_ACCESS_KEY_ID", "2R5YabYDYwesXPDPprWc6DpbczCsXL97"))
	config.MinioSecretKey = cast.ToString(getOrReturnDefault("MINIO_SECRET_KEY_ID", "Ps5Che6XtJ6JmvsFXrXUH3tnhxwnZNYh"))
	config.MinioBucketName = cast.ToString(getOrReturnDefault("MINIO_BACKET_NAME", "photos"))
	config.MinioLocation = cast.ToString(getOrReturnDefault("MINIO_LOCATION", "us-east-1"))
	config.MinioHost = cast.ToString(getOrReturnDefault("MINIO_HOST", "test.cdn.urecruit.udevs.io"))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
