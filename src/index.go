package main

import (
   "fmt"
   "log"
   "net/http"
)

func handleReq(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "404 not found.", http.StatusNotFound)
	}

	if req.Method != "GET" {
		http.Error(res, "Method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(res, "This is working")
}


func main() {
	const port string = ":3000";

	http.HandleFunc("/hello", handleReq)

	fmt.Printf("Server starting at port no." + port)
	if err:=http.ListenAndServe(port, nil); err != nil{
		log.Fatal(err);
	}
}