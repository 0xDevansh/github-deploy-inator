package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/DeathVenom54/github-deploy-inator/logger"
	"github.com/DeathVenom54/github-deploy-inator/structs"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data structs.GithubWebhook
	if err := decoder.Decode(&data); err != nil {
		logger.Error(err)
	}
	fmt.Println(data)

	_, err := w.Write([]byte("Hello"))
	if err != nil {
		return
	}
}
