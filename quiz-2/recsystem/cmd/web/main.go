// greeting    greeting
// Welcome to my page this is my main.go
package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MejiaFrancis/3161/3162/quiz-2/recsystem/internal/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// create a new type
type application struct {
	user models.UserModel
	//responses models.ResponseModel
	//options models.OptionsModel
}

func main() {
	// Create a flag for specifiying the port number \
	// when starting the server
	addr := flag.String("port", ":4000", "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("RECSYSTEM_DB_DSN"), "PlstgreSQL DSN")
	flag.Parse()

	// Create an instance of the connection pool
	db, err := openDB(*dsn)
	if err != nil {
		log.Println(err)
		return
	}
	// create an instance of the application type
	app := &application{
		user: models.UserModel{DB: db},
		//responses: models.ResponseModel{DB: db},
		//options: models.OptionsModel{DB: db},
	}

	defer db.Close()
	// acquired a  database connection pool
	log.Println("database connection pool established")
	// create customized server
	log.Printf("Start server on port %s", *addr)
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
		//IdleTimeout:  time.Minute,
		//ReadTimeout:  5 * time.Second,
		//WriteTimeout: 10 * time.Second,
	}

	err = srv.ListenAndServe()
	log.Fatal(err) //should not reach here
}

// Get a database connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	// use a context to check if the DB is reachable
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) //always to this
	defer cancel()                                                          // then this to clean up
	// let's ping the DB
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
