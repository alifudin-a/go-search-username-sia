package main

import (
	database "github.com/alifudin-a/go-search-username-sia/pkg/database/mysql"
	rd "github.com/alifudin-a/go-search-username-sia/pkg/database/redis"
	"github.com/alifudin-a/go-search-username-sia/pkg/routes"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	database.DBSIA()
	rd.OpenRedisSIA()
	routes.Routes()
}
