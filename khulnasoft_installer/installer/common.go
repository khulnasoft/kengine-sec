package installer

import "github.com/khulnasoft/khulnasoft_installer/agent"

type Installer interface {
	Delete() error
	Install() error
	SaveNewConfig(agent.AgentImage) error
	RollBackConfig()
}
