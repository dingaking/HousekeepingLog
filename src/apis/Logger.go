package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {

	logFile, err := os.OpenFile("log.txt", os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	/*defer logFile.Close()*/

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)
		log.SetOutput(logFile)
		log.Printf(
			"%\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)

	})
}
