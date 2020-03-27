package commands

import (
	"fmt"
	"os"

	"github.com/aalbacetef/ssh-helper/configfile"
	"github.com/aalbacetef/ssh-helper/utils"
)

const AddMsg = `

Usage:  ssh-helper add --name NAME
                       --hostname HOSTNAME
                       --identityfile PATH | --newkey
                       [--user USER]

Adds a host to the SSH config. Can use an existing keypair 
or generate one for you.

Options:
  --name            Name to use, this serves as an alias
  --hostname        Hostname to use. Must be DNS-resolvable
  --user            User to log in as
  --identityfile    Path to private key file
  --newkey          Generate a new private/public key pair

Example:
  To add a host with an existing key pair:
      ssh-helper add \
          --name my-test-host \
          --hostname my.test.host \
          --user testuser \
          --identityfile /path/to/private/key

  To add a host with the key pair being generated:
      ssh-helper add \
        --name my-test-host \
        --hostname my.test.host \
        --user testuser \
        --newkey 

  These commands will then allow you to interact with the 
  remote host in various ways:
      - ssh my-test-host
      - git clone my-test-host:account/repo.git (user might need to be git)
      - scp ./readme.txt my-test-host:~/
`

func AddUsage() {
	fmt.Println(AddMsg)
}

// Adds an entry. Assumes input has been validated to exist.
func AddEntry(name, hostname, user, identityfile string) {
	// load the existing configfile (assumes default path)
	// @TODO allow -path to be passed in
	config, err := configfile.LoadDefault()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	// add host to config struct
	hostentry := configfile.HostEntry{
		Host:             name,
		HostName:         hostname,
		User:             user,
		IdentityFilePath: configfile.Fpath(identityfile),
	}
	err = config.Add(hostentry)
	if err != nil {
		fmt.Println("error adding host: ")
		fmt.Println(err)
		return
	}

	// save result!
	defaultpaths, err := configfile.DefaultPaths()
	if err != nil {
		fmt.Println("error fetching default paths: ", err)
		os.Exit(1)
		return
	}

	configpath, exists := defaultpaths["config"]
	if !exists {
		fmt.Println("error: Need to define where config file is in DefaultPaths")
		os.Exit(1)
		return
	}

	err = utils.WriteToFile(
		configpath,
		config.FormattedText(),
	)

	if err != nil {
		fmt.Println("error while saving config file: ", err)
		os.Exit(1)
		return
	}

}
