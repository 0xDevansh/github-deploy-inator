package handlers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/DeathVenom54/github-deploy-inator/structs"
	"net/http"
	"os/exec"
	"strings"
)

func CreateWebhookHandler(config *structs.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var webhook structs.GithubWebhook
		err := decoder.Decode(&webhook)
		handleErr(err)

		// get the correct listener
		var listener *structs.Listener
		for _, l := range config.Listeners {
			if strings.ToLower(l.Repository) == strings.ToLower(webhook.Repository.FullName) {
				listener = &l
				break
			}
		}
		if listener == nil {
			panic(fmt.Errorf("no listener found for webhook from %s", webhook.Repository.FullName))
			return
		}

		// verify signature if secret is provided
		if listener.Secret != "" {
			hash := sha1.New()
			hash.Write([]byte(listener.Secret))
			signature := "sha1=" + hex.EncodeToString(hash.Sum(nil))
			fmt.Println(signature)
			fmt.Println(r.Header.Get("X-Hub-Signature"))
			if signature != r.Header.Get("X-Hub-Signature") {
				panic(fmt.Sprintf("received webhook from %s but signature does not match secret", webhook.Repository.FullName))
			}
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
			handleErr(err)

			if listener.Discord.NotifyBeforeRun {
				err := m.SendPreRunNotification(listener, &webhook)
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
				err := m.SendErrorMessage(listener, &err, &webhook)
				handleErr(err)
			}
			handleErr(err)
		} else if listener.NotifyDiscord {
			// send notification
			output := string(out)
			err := m.SendSuccessMessage(listener, &output, &webhook)
			handleErr(err)
		}

		w.WriteHeader(200)
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
