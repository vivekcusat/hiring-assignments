package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func serveRandomFile(w http.ResponseWriter, r *http.Request) {
	rnd := rand.Intn(10)
	if rnd < 5 {
		http.ServeFile(w, r, "./dummy.png")
		return
	}
	if rnd < 9 {
		http.ServeFile(w, r, "./dummy.pdf")
		return
	}
	http.ServeFile(w, r, "./corrupt-dummy.pdf")
}
func basicHomePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the Assignment home page!")
    fmt.Println("Endpoint Hit: basicHomePage")
}

func newEndpoint() {
    http.HandleFunc("/health", basicHomePage)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", serveRandomFile)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
