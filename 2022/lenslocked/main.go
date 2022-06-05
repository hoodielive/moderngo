package main

import 
(
	"fmt"
	"net/http"
)

func handlerFunc(w http.Response, r *http.Request)
{
	fmt.Fprint(w, "<h1>Welcome to the Egun Awo Site.")
}

func main()
{
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
