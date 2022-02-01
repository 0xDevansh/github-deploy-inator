package handlers

import (
	"encoding/json"
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
				reply(w)
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
				reply(w)
				return
			}
		}

		m := structs.DiscordNotificationManager{
			Webhook: structs.DiscordWebhookData{
				Url: listener.Discord.Webhook,
			},
		}
		if listener.NotifyDiscord {
			err := m.Setup()
			handleErr(err)

			if listener.Discord.NotifyBeforeRun {
				err := m.SendPreRunNotification(&listener, &webhook)
				handleErr(err)
			}
		}

		// run command
		args := strings.Split(listener.Command, " ")
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = listener.Directory

		out, err := cmd.Output()
		if err != nil {
			if listener.NotifyDiscord {
				err := m.SendErrorMessage(&listener, &err, &webhook)
				handleErr(err)
			}
			handleErr(err)
		} else if listener.NotifyDiscord {
			// send notification
			output := string(out)
			err := m.SendSuccessMessage(&listener, &output, &webhook)
			handleErr(err)
		}

		reply(w)
	}
}

func reply(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func handleErr(err error) {
	if err != nil {
		logger.Error(err, false)
	}

}
