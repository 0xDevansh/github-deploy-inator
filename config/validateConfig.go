package config

import (
	"errors"
	"fmt"
	"github.com/DeathVenom54/github-deploy-inator/structs"
	"regexp"
)

func ValidateConfig(config *structs.Config) error {
	// validate port and endpoint
	err := shouldMatchRegex("port", config.Port, `^:\d+$`)
	if err != nil {
		return err
	}
	err = shouldMatchRegex("endpoint", config.Endpoint, `^\/[\w-\/]*$`)
	if err != nil {
		return err
	}

	// validate listeners
	if len(config.Listeners) == 0 {
		return errors.New("no listeners assigned, please add at least one")
	}

	for i, listener := range config.Listeners {
		// required fields
		err := shouldExist("name", listener.Name, i)
		if err != nil {
			return err
		}
		err = shouldMatchRegex(fmt.Sprintf("listeners[%d].repository", i), listener.Repository, `[\w-]+\/[\w-]+`)
		if err != nil {
			return err
		}
		err = shouldExist("directory", listener.Directory, i)
		if err != nil {
			return err
		}
		err = shouldExist("command", listener.Command, i)
		if err != nil {
			return err
		}

		// discord
		if listener.NotifyDiscord {
			if listener.Discord.Webhook == "" {
				match, err := regexp.MatchString(structs.DiscordWebhookRegex, listener.Discord.Webhook)
				if err != nil {
					return err
				}
				if !match {
					return errors.New("please provide a valid Discord webhook url")
				}
				return errors.New(fmt.Sprintf("Discord.Webhook for listeners[%d] must be provided when NotifyDiscord is true\n", i))
			}
		}
	}
	return nil
}

func shouldMatchRegex(field, value, regex string) error {
	match, err := regexp.MatchString(regex, value)
	if err != nil {
		return err
	}
	if !match {
		return errors.New(fmt.Sprintf("invalid %s in config.json, should be in format %s", field, regex))
	}
	return nil
}

func shouldExist(paramName string, paramValue string, listenerIndex int) error {
	if paramValue == "" {
		return errors.New(fmt.Sprintf("Invalid %s for listeners[%d]: \"%s\"", paramName, listenerIndex, paramValue))
	}

	return nil
}
