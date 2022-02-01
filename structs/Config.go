package structs

type Config struct {
	Port      string     `json:"port"`     // example: ":8080"
	Endpoint  string     `json:"endpoint"` // endpoint where http server will listen for hooks
	Listeners []Listener `json:"listeners"`
}

type Listener struct {
	// required properties
	Name       string `json:"name"`       // a unique name for the webhook
	Repository string `json:"repository"` // the repository name in the format "username/repo-name"
	Directory  string `json:"directory"`  // the directory in which the command will be run
	Command    string `json:"command"`    // the command to run
	// for your deployment, it is suggested to put the various commands
	// in your scripts like in node.js or a .sh file, and execute it.

	// additional filters
	Branch         string   `json:"branch"`         // execute only if push is on this branch
	AllowedPushers []string `json:"allowedPushers"` // execute only if pusher is one of these (username)

	// notification options
	NotifyDiscord bool `json:"notifyDiscord"`
	Discord       struct {
		Webhook         string `json:"webhook"`
		NotifyBeforeRun bool   `json:"notifyBeforeRun"`
		SendOutput      bool   `json:"sendOutput"`
	}
}
