package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/frostyslav/cloudmobility-hackathon/app/webserver"
	"github.com/sirupsen/logrus"
)

func init() {
	l := logrus.New()
	l.Formatter = &logrus.TextFormatter{FullTimestamp: true, ForceColors: true}
	l.Level = logrus.DebugLevel
	l.Out = os.Stdout

	log.SetOutput(l.Writer())
}

func main() {

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//})

	port := 8080
	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), webserver.Serve))
}
