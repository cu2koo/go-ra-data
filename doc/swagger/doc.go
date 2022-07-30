// Package swagger Robo-Advisor Data API
//
// This interface combines interfaces to user databases and market data providers.
//
// This is intended for an implementation of a robo advisor and implements various data processes for
// data collection, data preparation, data storage and data security.
//
// To use it, the user must identify himself via BasicAuth.
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - basicAuth:
//
//     SecurityDefinitions:
//     basicAuth:
//          type: basic
//
// swagger:meta
package swagger
