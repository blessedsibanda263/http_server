package main

import "net/http"

func getComments(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("here you'll get the comments"))
}

func postComments(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("thank you for posting a comment"))
}

func main() {
	http.HandleFunc("GET /comments", getComments)
	http.HandleFunc("POST /comments", postComments)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
