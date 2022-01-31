package structs

type Config struct {
	Port      string     `json:"port"` // example: ":8080"
	Listeners []Listener `json:"listeners"`
}

type Listener struct {
	Name          string `json:"name"`
	Endpoint      string `json:"endpoint"`
	Directory     string `json:"directory"`
	Command       string `json:"command"`
	NotifyDiscord bool   `json:"notifyDiscord"`

	Discord struct {
		Webhook         string `json:"webhook"`
		NotifyBeforeRun bool   `json:"notifyBeforeRun"`
	}
}
