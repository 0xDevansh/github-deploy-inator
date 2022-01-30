package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)

	_, err := w.Write([]byte("Hello"))
	if err != nil {
		return
	}
}
