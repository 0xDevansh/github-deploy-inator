package structs

type Listener struct {
	Name          string `json:"name"`
	Endpoint      string `json:"endpoint"`
	Directory     string `json:"directory"`
	Command       string `json:"command"`
	NotifyDiscord bool   `json:"notifyDiscord"`

	Discord struct {
		Token           string `json:"token"`
		ChannelId       string `json:"channelId"`
		NotifyBeforeRun bool   `json:"notifyBeforeRun"`
	}
}
