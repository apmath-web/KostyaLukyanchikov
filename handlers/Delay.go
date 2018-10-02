package handlers

import (
	"fmt"
	"net/http"
)

func Delay(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("delay")
	if param != "" {
		fmt.Fprintln(w, "Was delayed for"+param+"ms")
	}
}
