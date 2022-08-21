package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"waysbuck/database"
	mysql "waysbuck/pkg"
	"waysbuck/routes"

	"github.com/gorilla/mux"
)

func main() {

	mysql.DatabaseInit()

	database.RunMigration()
	r := mux.NewRouter()

	routes.RouterInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
