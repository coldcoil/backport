package cmd

import "github.com/alecthomas/kong"

// CLI represents the commands, flags, and args possible for backport
type CLI struct {
	Link    Link             `cmd:"true"`
	Version kong.VersionFlag `name:"version" help:"Print version information and quit"`
}
