package cmd

import (
	"github.com/alecthomas/kong"
)

// App is passed to the Run function for commands
type App struct {
	Kong *kong.Context
}
