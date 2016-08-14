////////////////////////////////////////////////////////////////////////
// Outputting a Web Log
// In splitting up the routes file, I also had an ulterio motive. As you will see shortly, it now becomes very easy to decorate my http handlers with additional functonality.
// Let's start with the ability to log out web requests like most modern web services do. In Go, there is no web logging package or functionality in the standard library, so we have to create it.
// We'll do that by creating a file called logger.go and add the following code
////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		fmt.Println("ddddddd")

		inner.ServeHTTP(w, r)

		fmt.Printf("dddd")
		log.Printf(
			"%\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

////////////////////////////////////////////////////////////////////////
// This is a very standard idiom in Go. Effectively we are going to pass our handler to this function, which wil then wrap the passed handler with logging and timing functionality.
// Next, we will need to utilize the logger decorator in our routes.
////////////////////////////////////////////////////////////////////////
