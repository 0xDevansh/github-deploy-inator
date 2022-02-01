package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DeathVenom54/github-deploy-inator/logger"
	"github.com/DeathVenom54/github-deploy-inator/structs"
	"net/http"
	"os/exec"
	"strings"
)

func CreateWebhookHandler(config *structs.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var webhook structs.GithubWebhook
		if err := decoder.Decode(&webhook); err != nil {
			logger.Error(err, false)
			return
		}

		// get the correct listener
		var listener *structs.Listener
		for _, l := range config.Listeners {
			if strings.ToLower(l.Repository) == strings.ToLower(webhook.Repository.FullName) {
				listener = &l
				break
			}
		}
		if listener == nil {
			logger.Error(errors.New(fmt.Sprintf("No listener found for webhook from %s", webhook.Repository.FullName)), false)
			return
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
				if strings.ToLower(webhook.Pusher.Name) == strings.ToLower(pusher) {
					pusherIsAllowed = true
					break
				}
			}
			if !pusherIsAllowed {
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
			if err != nil {
				logger.Error(err, false)
			}

			if listener.Discord.NotifyBeforeRun {
				err := m.SendPreRunNotification(listener, &webhook)
				if err != nil {
					logger.Error(err, false)
				}
			}
		}

		// run command
		args := strings.Split(listener.Command, " ")
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = listener.Directory

		out, err := cmd.Output()
		if err != nil {
			if listener.NotifyDiscord {
				err := m.SendErrorMessage(listener, &err, &webhook)
				if err != nil {
					logger.Error(err, false)
				}
			}
			if err != nil {
				logger.Error(err, false)
				return
			}
		} else if listener.NotifyDiscord {
			// send notification
			output := string(out)
			err := m.SendSuccessMessage(listener, &output, &webhook)
			if err != nil {
				logger.Error(err, false)
			}
		}

		w.WriteHeader(200)
	}
}
