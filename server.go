package main

import (
  "github.com/scorbettUM/headspace/client_application/api_services/account_service/middlewares"
  // "github.com/scorbettUM/headspace/client_application/api_services/account_service/routes"
  // "github.com/mnmtanish/go-graphiql"
  "github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"net/http"
  "os"
  "fmt"
)

func StartServer(){

  ACCOUNT_SERVICE_PORT := os.Getenv("ACCOUNT_SERVICE_PORT")

  r := mux.NewRouter()
  // loginRoutes := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)

  // loginRoutes.Handle("/home", routes.Home)
  // loginRoutes.Handle("/profile", routes.Profile)
  // loginRoutes.Handle("/settings", routes.Settings)

  r.PathPrefix("/api/account").Handler(negroni.New(
    negroni.Wrap(middlewares.AuthMiddleware(middlewares.CorsHandler(accountRoutes)))
  ))

  r.PathPrefix("/api/incident").Handler(negroni.New(
    negroni.Wrap(middlewares.AuthMiddleware(middlewares.CorsHandler(incidentRoutes)))
  ))

  r.PathPrefix("/api/report").Handler(negroni.New(
    negroni.Wrap(middlewares.AuthMiddleware(middlewares.CorsHandler(reportRoutes)))
  ))

	fmt.Println(http.ListenAndServe(ACCOUNT_SERVICE_PORT, middlewares.CorsHandler(r)))
}
