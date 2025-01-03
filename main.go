package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Trynax/rssagg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)


type apiConfig struct {
	DB *database.Queries
}

func main() {
	

	godotenv.Load()

	portString:=os.Getenv("PORT")

	if portString==""{
	   log.Fatal("PORT is not found in the enviroment")
	}

	dbUrl:=os.Getenv("DB_URL")

	if portString==""{
	   log.Fatal("DB_URL is not found in the enviroment")
	} 

	conn, err:=sql.Open("postgres", dbUrl)

	if err!= nil {
       log.Fatal("can't connect to database: ",err)
    }
	
	apiCfg := apiConfig{
		DB: database.New(conn),
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()


	v1router.Get("/healthz", handlerReadiness)
	v1router.Get("/err", handlerErr)
	v1router.Post("/users", apiCfg.handlerCreateUser )

	router.Mount("/v1", v1router)
	


	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}

	log.Printf("server starting on port %v", portString)
	err= srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Port:", portString)
}