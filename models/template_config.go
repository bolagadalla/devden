package models

type TemplateConfig struct {
	Id                   string
	Name                 string
	Description          string
	URI                  string // can be file location or url
	CurrentLocation      string
	IsCloud              bool
	PreCreationCommands  []string
	PostCreationCommands []string
}
