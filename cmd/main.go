package main

import (
	"github.com/cu2koo/go-ra-data/data/factset"
	"github.com/cu2koo/go-ra-data/data/mariadb"
	"github.com/cu2koo/go-ra-data/http"
	"github.com/cu2koo/go-ra-data/model/data"
	"github.com/cu2koo/go-ra-data/service"
	"github.com/gin-gonic/gin"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger := log.New(os.Stdout, "[SERVER] ", log.LstdFlags)

	acc := &gin.Accounts{
		"user": "secret",
	}

	userDB, err := mariadb.New()
	if err != nil {
		logger.Fatalln("MariaDB initialization failed.")
	}

	factSet, err := factset.New()
	if err != nil {
		logger.Fatalln("FactSet initialization failed.")
	}
	marketDataProviders := &[]data.Market{
		factSet,
	}

	c, err := service.NewUserService(userDB)
	if err != nil {
		logger.Fatalln("UserService initialization failed.")
	}

	ms := service.NewMarketService(marketDataProviders)
	r := http.SetupRouter(acc, c, ms)

	certFile := "./config/server.crt"
	keyFile := "./config/server.key"

	err = r.RunTLS(":8080", certFile, keyFile)
	if err != nil {
		logger.Fatalln("Could not start server.")
	}
}
