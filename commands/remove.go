package commands

import (
	"fmt"
	"os"

	"gitlab.com/aalbacetef/ssh-helper/configfile"
	"gitlab.com/aalbacetef/ssh-helper/utils"
)

const RmMsg = `
Usage:  ssh-helper remove --host HOST

Options:
  --host    Host to be removed from the SSH config.
`

func RmUsage() {
	fmt.Println(RmMsg)
}

func Remove(host string) {
	// load default config
	config, err := configfile.LoadDefault()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	// remove host entry
	err = config.Rm(host)
	if err != nil {
		fmt.Println("error removing host: ", err)
		return
	}

	// save result!
	defaultpaths, err := configfile.DefaultPaths()
	if err != nil {
		fmt.Println("error fetching default paths: ", err)
		os.Exit(1)
		return
	}

	// get default config path
	configpath, exists := defaultpaths["config"]
	if !exists {
		fmt.Println("error: Need to define where config file is in DefaultPaths")
		os.Exit(1)
		return
	}

	// update file
	err = utils.WriteToFile(
		configpath,
		config.FormattedText(),
	)

	if err != nil {
		fmt.Println("error while saving config file: ", err)
		os.Exit(1)
		return
	}

	return
}
