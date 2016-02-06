package cmd

import "github.com/codegangsta/cli"

/// Usage: add --simulate --filter {patterns|files}
///
/// Note:
///   only files that are known and not ignored will be added.
///
/// Examples:
///   add all files -> add .
///   add one file  -> add "path\file.ext"
///   add two files -> add "path\file.ext, path\file2.ext"
///   add files, multiple expressions -> add "path\file???.*, path\data???.*"
///   add file(s) under path matching wildcard expression -> add "path\*.*"
///   add file(s) under path (recursively) matching wildcard expression -> add "path*\file???.*"

// Add is used to add files to be tracked
func Add(c *cli.Context) {
	println("Cmd Add")
}
