package utils

type StowConfigToml struct {
	Name         string
	Dependencies []string
	Scripts      []string
}
type ConfigToml struct {
	Stow     []StowConfigToml
	Flatpaks []string
}
