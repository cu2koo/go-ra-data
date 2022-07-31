# Robo-Advisor Data API

This interface for robo-advisors implements the data processes of a generic robo-advisor for data collection, data preparation, data storage and, to some extent, data security. <br/>
The implementation can be cloned using the ```git clone git@github.com:cu2koo/go-ra-data.git``` command

## Requirements:
The interface has a modular structure so that the requirements are interchangeable except Go.
You will need:
- [Go](https://go.dev/)
- Database for storing user data
- Data provider for collecting market data

The following technologies and providers are used in this implementation:
- [MariaDB](https://mariadb.org/) as a database for the user data
- [FactSet](https://www.factset.com/) as a data provider for the market data

## Configuration
In this example, the [MariaDB is installed via Docker](https://hub.docker.com/_/mariadb/) and then configured. A [Docker installation](https://docs.docker.com/engine/install/) is required for this. This is done in the following steps:
1. Creation of the MariaDB container with the command: ```docker run --detach --network bridge --name mariadb --env MARIADB_ROOT_PASSWORD=secret -p 3306:3306 mariadb:latest```
2. Open the MariaDB console and enter the password defined in the previous command via ```docker exec -it mariadb bash``` and ```mariadb -p```.
3. Running the SQL commands found at [data/mariadb/setup.sql](https://github.com/cu2koo/go-ra-data/blob/main/data/mariadb/setup.sql), for creating the database and tables.

The following files are required for the configuration:
- SSL certificate and key are placed under [config](https://github.com/cu2koo/go-ra-data/tree/main/config) path. Commands for self-signed certificate creation are available as a [shell script](https://github.com/cu2koo/go-ra-data/blob/main/config/certificates.sh).
- MariaDB's configuration is saved as _config.json_ following the pattern of [data/mariadb/config-template.json](https://github.com/cu2koo/go-ra-data/blob/main/data/mariadb/config-template.json).
- FactSet`s configuration is saved as _config.json_ following the pattern of [data/factset/config-template.json](https://github.com/cu2koo/go-ra-data/blob/main/data/factset/config-template.json). To do this, the API key must be created via the [FactSet Developer Portal](https://developer.factset.com/).

## Execution
The application can be compiled and executed via Go. This following command can be used in the project directory for execution: ```go run ./cmd/main.go```.

## Documentation
The documentation is available as [swagger.yaml](https://github.com/cu2koo/go-ra-data/blob/main/doc/swagger.yaml). This file can be opened via the [Swagger Editor](https://editor.swagger.io/), for example.