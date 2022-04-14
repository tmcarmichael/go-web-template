// Entry point to front end of application.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type appConfig = struct {
	serviceName string
	port        int
	env         string
	api         string
	util        struct {
		key    string
		secret string
	}
	db struct {
		connectionString string
	}
}

type application struct {
	config   appConfig
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       60 * time.Second,
		ReadTimeout:       20 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	app.infoLog.Printf("Starting HTTP service `%s` in `%s` mode on port: %d", app.config.serviceName, app.config.env, app.config.port)
	return srv.ListenAndServe()
}

func main() {
	var cfg appConfig

	// Define flags with default values.
	flag.IntVar(&cfg.port, "port", 8080, "Port for service to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment")
	flag.StringVar(&cfg.api, "api", "http://localhost:8000", "URL to API")
	flag.StringVar(&cfg.serviceName, "serviceName", "go-web-template", "Service name")
	flag.Parse()

	// Set auth config of cfg, keys from env etc.
	cfg.util.key = os.Getenv("UTIL_KEY")
	cfg.util.secret = os.Getenv("UTIL_SECRET")

	// Set up base std/err logging.
	infoLog := log.New(os.Stdout, "INFO  ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR  ", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
	}

	err := app.serve()
	if err != nil {
		app.errorLog.Println(err)
		os.Exit(1)
	}
}
