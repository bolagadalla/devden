package models

type TemplateConfig struct {
	Id                   string
	Name                 string
	Description          string
	URI                  string // can be file location or url
	IsCloud              bool
	PreCreationCommands  []string
	PostCreationCommands []string
}
