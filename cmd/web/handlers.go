package main

import "net/http"

func (app *application) ExampleEndpoint(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Service `%s` endpoint `ExampleEndpoint` reached.", app.config.serviceName)
}
