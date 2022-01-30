package handlers

import (
	"fmt"
	"github.com/DeathVenom54/github-deploy-inator/logger"
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err)
	}
	fmt.Println(string(bodyBytes))

	_, err = w.Write([]byte("Hello"))
	if err != nil {
		return
	}
}
