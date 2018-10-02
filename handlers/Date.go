package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func date(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Date:", time.Now().UTC().Format("2006-01-02T15:04:05-0700"))
}