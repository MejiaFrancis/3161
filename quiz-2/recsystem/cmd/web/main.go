// main.go
package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MejiaFrancis/3161/3162/quiz-2/recsystem/internal/models" // this will be change to file from another folder models
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Share data across our handlers

type application struct {
	reservations models.ReservationModel
	login        models.LoginModel
}

func main() {
	// configure our server
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("RCSYSTEM_DB_DSN"), "PostgreSQL DSN (Data Source Name)")
	flag.Parse()

	// get a database connection pool
	db, err := openDB(*dsn)
	if err != nil {
		log.Print(err)
		return
	}

	// share data across our handlers
	app := &application{
		reservations: models.ReservationModel{DB: db},
		login:        models.LoginModel{DB: db},
	}
	// cleanup the connection pool
	defer db.Close()
	// acquired a database connection pool
	log.Println("database connection pool established")
	// create and start a custom web server
	log.Printf("starting server on %s", *addr)
	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

// The openDB() function returns a database connection pool or error
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	// create a context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// test the DB connection
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
