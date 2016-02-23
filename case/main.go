package main

import (
	"os"

	"github.com/jurgen-kluft/Case/cmd"
	"github.com/jurgen-kluft/writ"
)

func appmain(c *cli.Context) {
	println("Hello friend!")
}

func main() {
	app := cli.NewApp()
	app.Name = "Bras"
	app.Usage = "Bras CMD[add, commit, update, status] {OPTIONS}"
	app.Action = appmain
	app.Version = "1.0.0"

	cmd.Register(app)
	/// add    --simulate --filter {patterns|files}
	/// update --simulate --depot $NAME --null --force --check --modified
	/// status --clean --modified --added --missing --untracked --filter --datetime --verify
	/// sync   --simulate --force
	/// revert --simulate --force {patterns|files}
	/// rename --simulate $SRCFILE $DSTFILE
	/// forget --simulate --added --ignored
	/// purge  --simulate --untrackedcli

	app.Run(os.Args)
}
