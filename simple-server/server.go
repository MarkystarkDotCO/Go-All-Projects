package main

import (
	"fmt"
	"net/http"
)

func main()  {

	http.HandleFunc("/home", home)


	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello home")
}
