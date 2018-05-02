package Utility

const (
	YapaDir = ".yapa"

	ConfigJSON = "config.json"

	TodoJSON = "todo.json"

	ServersJSON = "servers.json"
)

var (
	DefaultYapaDir = UserHomeDir() + "/" + YapaDir

	DefaultYapaConfigPath = DefaultYapaDir + "/" + ConfigJSON

	DefaultYapaTodoJSONPath = DefaultYapaDir + "/" + TodoJSON

	DefaultYapaServerConfigPath = DefaultYapaDir + "/" + ServersJSON
)

type YapaConfig struct {
	Username string `json:"username"`
	System   string `json:"system"`
}
