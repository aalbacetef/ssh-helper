package main

import (
	"flag"
	"fmt"
	"os"

	"gitlab.com/aalbacetef/ssh-helper/commands"
	"gitlab.com/aalbacetef/ssh-helper/utils"
)

func main() {
	// Command: Add
	// Flags:
	//	-name
	//	-hostname
	//	-user
	//	-identityfile
	//	-genkey
	//
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addCmd.Usage = commands.AddUsage
	hostname := addCmd.String("hostname", "", "IP address or resolvable domain name.")
	identityfile := addCmd.String(
		"identityfile",
		"",
		"The `path` to IdentityFile to use in case you want to manually set it. "+
			"Recommended to let ssh-helper manage that for you.",
	)
	name := addCmd.String("name", "", "Host to add.")
	newkey := addCmd.Bool("newkey", false, "Have ssh-helper generate an SSH key pair for you.")
	user := addCmd.String("user", "", "User to add.")

	// Command: Backup
	// Flags:
	//  -backup
	//
	backupCmd := flag.NewFlagSet("build", flag.ExitOnError)
	backupCmd.Usage = commands.BackupUsage
	backupTag := backupCmd.String("tag", "", "A tag to use to name the backup.")
	backupPath := backupCmd.String("path", "", "Use this path as the backup directory.")

	// Command: List
	// Flags:
	//	-all
	//  -stored
	//	-json
	//
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listCmd.Usage = commands.ListUsage
	listUsejson := listCmd.Bool("json", false, "Use JSON as output format. Defaults to text.")

	// set usage
	flag.Usage = Usage

	// not enough arguments
	if len(os.Args) < 2 {
		Usage()
		return
	}

	// handle command
	switch os.Args[1] {
	case "add":
		_ = addCmd.Parse(os.Args[2:])
	case "backup":
		_ = backupCmd.Parse(os.Args[2:])
	case "list":
		_ = listCmd.Parse(os.Args[2:])
	default:
		flag.Parse()
	}

	// Set output mode to Formatted Text by default
	mode := commands.FORMATTED_TEXT

	// Command: Add
	if addCmd.Parsed() {
		// new config
		if *name == "" {
			fmt.Println("error: -name is required.")
			addCmd.Usage()

			return
		}

		// hostname must exist
		if *hostname == "" {
			fmt.Println("error: -hostname is required.")
			addCmd.Usage()

			return
		}

		// either a key is generated or is provided
		if *identityfile == "" && !(*newkey) {
			fmt.Println("error: -newkey or -identityfile must be specified.")
			addCmd.Usage()

			return
		}

		// generate keypair
		if *newkey {
			keypath, err := utils.NewKey(*name)
			if err != nil {
				fmt.Println("error: could not create keypair.\n", err)
				return
			}

			fmt.Println("Generated key at: ", keypath)

			// point identityfile at keypath now
			identityfile = &keypath
		}

		// add entry to config file
		commands.AddEntry(*name, *hostname, *user, *identityfile)

		return
	}

	if backupCmd.Parsed() {
		if *backupTag == "" {
			fmt.Println("error: --tag must be supplied")
			backupCmd.Usage()
			return
		}

		commands.Backup(*backupTag, *backupPath)

		return
	}

	// Command: List
	if listCmd.Parsed() {
		// set output mode to JSON
		if *listUsejson {
			mode = commands.JSON
		}

		commands.ListAll(mode)

		return
	}

	// should not have reached here
	fmt.Println("unrecognized command: ", os.Args[1])
	flag.Usage()
}

const HeaderMsg = `
SSH-Helper
-----------

Usage: ssh-helper COMMAND [OPTIONS]


A tool to manage your ssh configs. By default 
uses ~/.ssh/ssh-helper to manage all configs.

Commands:
  add       Add an ssh config, can generate a key automatically
  backup    Backs up the current ~/.ssh directory
  config    Print the current configs of ssh-helper
  list      List all available hosts
  remove    Remove a host from the config. This operation will not delete the key unless asked to.


Run 'ssh-helper COMMAND --help' for more information.
`

func Usage() {
	fmt.Println(HeaderMsg)
}
