package service

import (
    "net/http"
	"fmt"
)

func unknown() http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "501 Not Implemented")
    }
}