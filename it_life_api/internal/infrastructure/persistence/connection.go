package persistence

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connection is a singleton instance to contain connection.
var connection *gorm.DB

// getDBConnection creates a DB connection and return.
func getDBConnection() *gorm.DB {
	if connection == nil {
		var err error

		clearDefaultConfig()

		connection, err = gorm.Open(postgres.Open(getDBConfig()), &gorm.Config{})
		if err != nil {
			panic("failed to connect to the database")
		}
	}

	return connection
}

// TODO: Close connection.
// closeDBConnection closes a DB connection.
// func closeDBConnection() {
// 	if connection != nil {
// 		connection, err := connection.DB()
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer connection.Close()
// 	}
// }

// clearDefaultConfig clears default connection config in OS environment.
func clearDefaultConfig() {
	// Unset these because login with service file isn't supported by golang postgres library.
	_ = os.Unsetenv("PGSERVICE")
	_ = os.Unsetenv("PGSERVICEFILE")
	_ = os.Unsetenv("PGREALM")
}

// getDBConfig gets DB config.
func getDBConfig() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := getSSLMode()
	dbTimeZone := getTimeZone()

	if len(dbHost) == 0 || len(dbPort) == 0 || len(dbUser) == 0 || len(dbName) == 0 || len(dbPassword) == 0 {
		panic("database config is not specified")
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s %s %s", dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode, dbTimeZone)
}

// getSSLMode gets SSL mode.
func getSSLMode() string {
	if os.Getenv("GO_ENV") == "local" {
		return "sslmode=disable"
	}
	return ""
}

// getTimeZone gets timezone.
func getTimeZone() string {
	return "TimeZone=Asia/Tokyo"
}
