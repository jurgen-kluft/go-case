package cmd

import "github.com/codegangsta/cli"

/// Usage: commit --simulate --missing --depot $NAME

// [CommandLine.Option('m', "missing", DefaultValue = false, HelpText = "Commit 'latest' files that have missing-chunks")]
// public bool Missing { get; set; }
//
// [CommandLine.Option('p', "push", DefaultValue = true, HelpText = "Push commit to server")]
// public bool Push { get; set; }
//
// [CommandLine.Option('d', "depot", DefaultValue = "current", HelpText = "The depot key")]
// public string Depot { get; set; }
//
// [CommandLine.Option('w', "wait", DefaultValue = false, HelpText = "Push commit to server and wait")]
// public bool Wait { get; set; }
//
// [CommandLine.Option('s', "simulate", HelpText = "Run the command in simulation")]
// public bool Simulation { get; set; }
//
// [CommandLine.Option('b', "debug", HelpText = "Do not remove the commit and log files")]
// public bool Debug { get; set; }
//
// [CommandLine.Option('h', "help", HelpText = "Show usage information of a command", Required = false)]

// Commit will process any 'modified' file, update the meta-information and add
// the file content to the data back-end
func Commit(c *cli.Context) {
	println("Cmd Commit")
}
