package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"inventoryGql/graph/generated"
	"inventoryGql/graph/models"
	"inventoryGql/graph/resolvers"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=application dbname=inventoryGql password=password sslmode=disable")
	if err != nil {
		return
	}
	defer db.Close()
	db.AutoMigrate(models.Inventory{}, models.Item{}, models.InventoryItem{})
	db.LogMode(true)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{
			Db: db,
		},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
