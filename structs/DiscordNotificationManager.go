package structs

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
	"regexp"
	"time"
)

const DiscordWebhookRegex = `^https:\/\/discord.com\/api\/webhooks\/(?P<id>\d+)\/(?P<token>\w+)$`

type DiscordNotificationManager struct {
	Session         *discordgo.Session
	Webhook         DiscordWebhookData
	NotifyBeforeRun bool
	SendOutput      bool
}

type DiscordWebhookData struct {
	Url   string
	Id    string
	Token string
}

func (m *DiscordNotificationManager) Setup() error {
	// parse ID and Token
	if m.Webhook.Url == "" {
		return errors.New("no webhook.url found")
	}
	r, err := regexp.Compile(DiscordWebhookRegex)
	if err != nil {
		return err
	}

	matches := r.FindStringSubmatch(m.Webhook.Url)
	if len(matches) < 3 {
		return errors.New("the provided webhook url is invalid")
	}
	m.Webhook.Id = matches[1]
	m.Webhook.Token = matches[2]

	// create session
	m.Session, err = discordgo.New()
	if err != nil {
		return err
	}

	return nil
}

func (m *DiscordNotificationManager) SendPreRunNotification(listener *Listener, ghWebhook *GithubWebhook) error {
	t := time.Now().UTC().Unix()

	preRunEmbed := embed.NewEmbed().
		SetColor(blurple).
		SetTitle(fmt.Sprintf("Deploying %s...", listener.Name)).
		AddField("Repository", fmt.Sprintf("[%s](%s)", ghWebhook.Repository.FullName, ghWebhook.Repository.URL)).
		AddField("Pusher", fmt.Sprintf("[%s](%s)", ghWebhook.Pusher.Name, "https://github.com/"+ghWebhook.Pusher.Name)).
		AddField("Branch", fmt.Sprintf("[%s](%s)", ghWebhook.Ref[11:], ghWebhook.Repository.URL+"/tree/"+ghWebhook.Ref[11:])).
		AddField("Command", listener.Command).
		AddField("Time", fmt.Sprintf("<t:%d:f>", t)).
		MessageEmbed

	webhookParams := discordgo.WebhookParams{Embeds: []*discordgo.MessageEmbed{preRunEmbed}}

	_, err := m.Session.WebhookExecute(m.Webhook.Id, m.Webhook.Token, false, &webhookParams)
	if err != nil {
		return err
	}
	return nil
}

func (m *DiscordNotificationManager) SendSuccessMessage(listener *Listener, output *string, ghWebhook *GithubWebhook) error {
	t := time.Now().UTC().Unix()

	successEmbed := embed.NewEmbed().
		SetTitle(fmt.Sprintf("Succesfully deployed %s", listener.Name)).
		SetColor(green)

	if !listener.Discord.NotifyBeforeRun {
		successEmbed = successEmbed.
			AddField("Repository", fmt.Sprintf("[%s](%s)", ghWebhook.Repository.FullName, ghWebhook.Repository.URL)).
			AddField("Pusher", fmt.Sprintf("[%s](%s)", ghWebhook.Pusher.Name, "https://github.com/"+ghWebhook.Pusher.Name)).
			AddField("Branch", fmt.Sprintf("[%s](%s)", ghWebhook.Ref[11:], ghWebhook.Repository.URL+"/tree/"+ghWebhook.Ref[11:])).
			AddField("Command", listener.Command)
	}
	if listener.Discord.SendOutput {
		successEmbed = successEmbed.
			AddField("Output", fmt.Sprintf("```\n%s\n```", *output))
	}
	successEmbed = successEmbed.AddField("Time", fmt.Sprintf("<t:%d:f>", t))

	webhookParams := discordgo.WebhookParams{Embeds: []*discordgo.MessageEmbed{successEmbed.MessageEmbed}}
	_, err := m.Session.WebhookExecute(m.Webhook.Id, m.Webhook.Token, false, &webhookParams)
	if err != nil {
		return err
	}

	return nil
}

func (m *DiscordNotificationManager) SendErrorMessage(listener *Listener, error *error, ghWebhook *GithubWebhook) error {
	t := time.Now().UTC().Unix()

	errorEmbed := embed.NewEmbed().
		SetTitle(fmt.Sprintf("There was an error while deploying %s", listener.Name)).
		SetColor(red).
		AddField("Error", (*error).Error())

	if !listener.Discord.NotifyBeforeRun {
		errorEmbed = errorEmbed.
			AddField("Repository", ghWebhook.Repository.FullName).
			AddField("Pusher", ghWebhook.Pusher.Name).
			AddField("Branch", ghWebhook.Ref[11:]).
			AddField("Command", listener.Command)
	}
	errorEmbed = errorEmbed.AddField("Time", fmt.Sprintf("<t:%d:f>", t))

	webhookParams := discordgo.WebhookParams{Embeds: []*discordgo.MessageEmbed{errorEmbed.MessageEmbed}}
	_, err := m.Session.WebhookExecute(m.Webhook.Id, m.Webhook.Token, false, &webhookParams)
	if err != nil {
		return err
	}

	return nil
}
