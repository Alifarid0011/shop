package main

import (
	"log"

	"github.com/Alifarid0011/shop/src/api"
	"github.com/Alifarid0011/shop/src/config"
	"github.com/Alifarid0011/shop/src/data/cache"
	"github.com/Alifarid0011/shop/src/data/db"
)

// @securityDefinitions.Apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	err := cache.InitRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.CloseRedis()
	err = db.InitDb(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDb()
	api.InitServer()
}
