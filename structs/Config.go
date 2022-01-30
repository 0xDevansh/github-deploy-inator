package structs

type Config struct {
	Port      int16      `json:"port"`
	Listeners []Listener `json:"listeners"`
}
