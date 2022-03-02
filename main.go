package main
/* Created By Andri L. Vicko, OCP, OCE, OCM */
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", ShowIndex)
	fmt.Println("Connected to port 9001")
	log.Fatal(http.ListenAndServe(":9001", router))
}

func ShowIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
