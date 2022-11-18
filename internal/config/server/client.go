package server

// ClientList object client list
type ClientList struct {
	Rest Rest
}

// Rest object rest
type Rest struct {
	MasterData Client `yaml:"masterdata"`
}

// Client object client
type Client struct {
	URL          string `yaml:"url"`
	ClientID     string `yaml:"clientID"`
	ClientSecret string `yaml:"clientSecret"`
	ChannelID    string `yaml:"channelID"`
}

// Client object client
type ClientTimeout struct {
	URL     string `yaml:"url"`
	Timeout string `yaml:"timeout"`
}
