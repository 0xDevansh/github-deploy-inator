package structs

type Config struct {
	Port      string     `json:"port"` // example: ":8080"
	Listeners []Listener `json:"listeners"`
}

type Listener struct {
	// required properties
	Endpoint  string `json:"endpoint"`  // endpoint where http server will listen for hooks
	Directory string `json:"directory"` // the directory in which the command will be run
	Command   string `json:"command"`   // the command to run
	// for your deployment, it is suggested to put the various commands
	// in your scripts like in node.js or a .sh file, and execute it.

	//additional filters
	Branch  string   `json:"branch"`  // execute only if push is on this branch
	Authors []string `json:"authors"` // execute only if head commit author is one of these

	NotifyDiscord bool `json:"notifyDiscord"`

	Discord struct {
		Webhook         string `json:"webhook"`
		NotifyBeforeRun bool   `json:"notifyBeforeRun"`
		SendOutput      bool   `json:"sendOutput"`
	}
}
