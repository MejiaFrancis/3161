// greeting    greeting
// Welcome to my page this is my main.go
package main

import (
	"context"
	"crypto/tls"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MejiaFrancis/3161/3162/test-1/recsystem/internal/models"

	"github.com/alexedwards/scs/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// create a new type
type application struct {
	errorLog          *log.Logger
	infoLog           *log.Logger
	user              models.UserModel
	equipments        models.EquipmentModel
	role              models.RoleModel
	reservations      models.ReservationModel
	logs              models.LogModel
	feedback          models.FeedbackModel
	equipmentusagelog models.EquipmentUsageLogModel
	equipment_types   models.EquipmentTypeModel
	announcements     models.AnnouncementModel
	sessionManager    *scs.SessionManager
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
	//create instances of errorLog & infoLog
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// setup a new session manager
	sessionManager := scs.New()
	sessionManager.Lifetime = 1 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode

	// create an instance of the application type
	app := &application{
		errorLog:          errorLog,
		infoLog:           infoLog,
		user:              models.UserModel{DB: db},
		equipments:        models.EquipmentModel{DB: db},
		role:              models.RoleModel{DB: db},
		reservations:      models.ReservationModel{DB: db},
		logs:              models.LogModel{DB: db},
		feedback:          models.FeedbackModel{DB: db},
		equipmentusagelog: models.EquipmentUsageLogModel{DB: db},
		equipment_types:   models.EquipmentTypeModel{DB: db},
		announcements:     models.AnnouncementModel{DB: db},
		sessionManager:    sessionManager,

		//responses: models.ResponseModel{DB: db},
		//options: models.OptionsModel{DB: db},
	}

	defer db.Close()
	// acquired a  database connection pool
	log.Println("database connection pool established")
	// create customized server
	log.Printf("Start server on port %s", *addr)
	// configure TLS
	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}
	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig:    tlsConfig,
	}

	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	log.Fatal(err) //should not reach here
}

// Get a database connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	// use a context to check if the DB is reachable
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //always to this
	defer cancel()                                                          // then this to clean up
	// let's ping the DB
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
