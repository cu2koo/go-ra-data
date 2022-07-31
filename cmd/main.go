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
	// Initialization of Logger for logging purpose.
	logger := log.New(os.Stdout, "[SERVER] ", log.LstdFlags)

	// Creation of users for authentication and authorization.
	// Can be outsourced to databases in the future.
	acc := &gin.Accounts{
		"user": "secret",
	}

	// Creation of the data connection implementation for the user service.
	userDB, err := mariadb.New()
	if err != nil {
		logger.Fatalln("MariaDB initialization failed.")
	}

	// Creation of the data connection implementation for the market service.
	factSet, err := factset.New()
	if err != nil {
		logger.Fatalln("FactSet initialization failed.")
	}

	// Creation of a data connection implementation list for the market service.
	// Is required for the use of optionally several data connection implementations.
	marketDataProviders := &[]data.Market{
		factSet,
	}

	// Initialization of the user service implementation with the data connection implementation.
	c, err := service.NewUserService(userDB)
	if err != nil {
		logger.Fatalln("UserService initialization failed.")
	}

	// Initialization of the market service implementation with the list of data connection implementations.
	ms := service.NewMarketService(marketDataProviders)

	// Initialization of the routing with the services and the authorized users.
	r := http.SetupRouter(acc, c, ms)

	// Definition of paths for server key and server certificate.
	certFile := "./config/server.crt"
	keyFile := "./config/server.key"

	// Start of the server using the TLS protocol.
	// For use without TLS protocol, r.Run(":8080") should be executed instead of r.RunTLS(":8080", certFile, keyFile)
	err = r.RunTLS(":8080", certFile, keyFile)
	if err != nil {
		logger.Fatalln("Could not start server.")
	}
}
