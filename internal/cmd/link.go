package cmd

import "fmt"

type Link struct {
	Config string `arg:"true" help:"Path to configuration file" type:"existingfile" default:"./backport.yaml"`
}

func (l *Link) Run(app *App) error {
	return fmt.Errorf("not implemented")
}
