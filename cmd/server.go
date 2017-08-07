package main

import "net/http"
import "ferp/pkg/app"

const port = ":8080"

func main() {
	mux := http.NewServeMux()
	app.Router(mux)
	http.ListenAndServe(port, mux)
}
