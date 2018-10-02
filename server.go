package main

import (
	"fmt"
	"github.com/KostyaLukyanchikov/server/handlers"
	"net/http"
)

func main()  {


	http.HandleFunc("/", handlers.Root)
	http.HandleFunc("/date", handlers.Date)
	http.HandleFunc("/date", handlers.Delay)


	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}