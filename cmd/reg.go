package cmd

import "github.com/codegangsta/cli"

// Register all commands
func Register(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add file(s) to track",
			Action:  Add,
		},
		{
			Name:    "commit",
			Aliases: []string{"c"},
			Usage:   "commit modified files",
			Action:  Commit,
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "update work directory",
		},
	}
}
