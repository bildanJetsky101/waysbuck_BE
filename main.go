package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"waysbuck/database"
	mysql "waysbuck/pkg"
	"waysbuck/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	mysql.DatabaseInit()

	database.RunMigration()
	r := mux.NewRouter()

	routes.RouterInit(r.PathPrefix("/api/v1").Subrouter())

	// Setup allowed Header, Method, and Origin for CORS on this below code ...
	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = "5000"
	fmt.Println("server running localhost:" + port)

	// Embed the setup allowed in 2 parameter on this below code ...
	http.ListenAndServe("localhost:"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
