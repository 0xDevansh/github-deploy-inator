package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/DeathVenom54/github-deploy-inator/logger"
	"github.com/DeathVenom54/github-deploy-inator/structs"
	"net/http"
	"os/exec"
	"strings"
)

func CreateWebhookHandler(listener structs.Listener) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var webhook structs.GithubWebhook
		if err := decoder.Decode(&webhook); err != nil {
			logger.Error(err, false)
		}

		// run filters
		if listener.Branch != "" {
			branch := webhook.Ref[11:]
			if listener.Branch != branch {
				return
			}
		}
		if len(listener.AllowedPushers) > 0 {
			pusherIsAllowed := false
			for _, pusher := range listener.AllowedPushers {
				if webhook.Pusher.Name == pusher {
					pusherIsAllowed = true
					break
				}
			}
			if !pusherIsAllowed {
				return
			}
		}

		args := strings.Split(listener.Command, " ")
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = listener.Directory

		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))

		_, err = w.Write([]byte("Hello"))
		if err != nil {
			return
		}
	}
}
