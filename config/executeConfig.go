package config

import "fmt"

func ExecuteConfig() error {
	config, err := ReadConfig()
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", config)
	return nil
}
