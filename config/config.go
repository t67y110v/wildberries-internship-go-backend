package config

import "os"

func ConfigurationSetUp() {
	//конфиги бд
	os.Setenv("DB_USERNAME", "DB_USERNAME")
	os.Setenv("DB_PASSWORD", "DB_PASSWORD")
	os.Setenv("DB_HOST", "5432")
	os.Setenv("DB_NAME", "Orders")
	os.Setenv("DB_POOL_MAXCONN", "5")
	os.Setenv("DB_POOL_MAXCONN_LIFETIME", "300")

}
