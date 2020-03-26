package commands

import (
	"fmt"

	"gitlab.com/aalbacetef/ssh-helper/configfile"
)

// These are the supported output formats for commands.
type OutputMode string

const (
	FORMATTED_TEXT OutputMode = "formatted-text"
	JSON           OutputMode = "json"
)

// Prints the list of all configs available (for now same as listStored).
func ListAll(mode OutputMode) {
	config, err := configfile.LoadDefault()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	// default stringification is JSON (indented)
	// @TODO maybe consider returning string instead?
	if mode == FORMATTED_TEXT {
		fmt.Println(config.FormattedText())
	} else if mode == JSON {
		fmt.Println(config.String())
	}
}

const ListMsg = `
Usage:  ssh-helper list [--json]

Options:
  --json    Sets the output mode to JSON. Default is formatted text.
`

func ListUsage() {
	fmt.Println(ListMsg)
}
