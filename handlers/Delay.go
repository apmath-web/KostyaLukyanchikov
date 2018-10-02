package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Delay(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("delay")
	if param != "" {
		DelayTime, err := strconv.Atoi(param)
		if err == nil {
			time.Sleep(time.Duration(DelayTime) * time.Millisecond)
			fmt.Fprintln(w, "Was delayed for "+param+" ms")
		}
	}
}
