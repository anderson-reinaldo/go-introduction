package context

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Server() {
	fmt.Println("Server is running in port :8082")
	http.HandleFunc("/", home)
	http.ListenAndServe(":8082", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("Iniciou minha request")
	defer log.Println("Finalizou minha request")

	select {
	case <-time.After(5 * time.Second):
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Request finalizada"))
		return

	case <-ctx.Done():
		fmt.Fprint(w, "Request cancelada")
		http.Error(w, "Request cancelada", http.StatusRequestTimeout)
		return
	}
}
