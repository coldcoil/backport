package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/coldcoil/backport/internal/cmd"
	"github.com/prometheus/common/version"
	"go.uber.org/zap"
)

func main() {
	err := run()
	if err != nil {
		zap.L().With(zap.Error(err)).Fatal("unhandled error")
	}
}

func run() error {
	l, err := zap.NewDevelopment()
	if err != nil {
		return fmt.Errorf("creating logger: %w", err)
	}
	zap.ReplaceGlobals(l)
	cli := &cmd.CLI{}
	ktx := kong.Parse(cli,
		kong.Name("backport"),
		kong.Description("A network service tunnel"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": version.Version,
		})
	err = ktx.Run(&cmd.App{
		Kong: ktx,
	})
	if err != nil {
		return err
	}
	return nil
}
