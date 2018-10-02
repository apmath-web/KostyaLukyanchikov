package main

import (
	"github.com/apmath-web/KostyaLukyanchikov/tree/dev/handlers"
	"fmt"
	"net/http"
)

func rootpage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello World!")
}

//func date(w http.ResponseWriter, r *http.Request) {
//	http.HandleFunc("/date",
//	fmt.Fprintln(w, "Date:", time.Now().UTC().Format("2006-01-02T15:04:05-0700"))
//})
func main()  {

	//http.HandleFunc("/wait",
	//	func(w http.ResponseWriter, r *http.Request) {
	//		fmt.Fprintln(w, "Date:", time.Now().UTC().Format("2006-01-02T15:04:05-0700"))
	//	})


	http.HandleFunc("/", rootpage)
	http.HandleFunc("/date", handlers.Date)


	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}