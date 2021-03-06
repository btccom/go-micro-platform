package debug

import (
	"fmt"

	"github.com/btccom/go-micro-platform/v2/plugin"
)

var (
	defaultManager = plugin.NewManager()
)

// Plugins lists the debug plugins
func Plugins() []plugin.Plugin {
	return defaultManager.Plugins()
}

// Register registers a debug plugin
func Register(pl plugin.Plugin) error {
	if plugin.IsRegistered(pl) {
		return fmt.Errorf("%s registered globally", pl.String())
	}
	return defaultManager.Register(pl)
}
