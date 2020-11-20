package configfile

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/aalbacetef/ssh-helper/utils"
)

type Fpath string

// Returns a map containing the path to the ssh-helper folder
// as well as the ssh config file.
//
// If the user home directory cannot be loaded, an error will be returned.
//
// @TODO:
// 	- Allow passing in a base directory
//
func DefaultPaths() (map[string]string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("error: could not load user home directory.\n", err)
		return nil, err
	}

	return map[string]string{
		"ssh-helper": path.Join(homedir, ".ssh", "ssh-helper"),
		"config":     path.Join(homedir, ".ssh", "config"),
	}, nil
}

// Generate a config file from a given path.
func Load(fpath string) (*ConfigFile, error) {
	data, _ := utils.LoadFile(fpath)
	c, err := Parse(data)
	c.Fpath = Fpath(fpath)
	return c, err
}

// Generate a ConfigFile format from the default path
func LoadDefault() (*ConfigFile, error) {
	defaults, err := DefaultPaths()
	if err != nil {
		return nil, err
	}

	defaultpath, exists := defaults["config"]
	if !exists {
		return nil, errors.New("config key is not defined in default paths")
	}
	return Load(defaultpath)
}

// Represents an SSH config file and is the struct to
// both read and write to the config file.
type ConfigFile struct {
	Name  string      `json:"name"`
	Fpath Fpath       `json:"fpath"`
	Hosts []HostEntry `json:"hosts"`
}

// Add a host entry to the config file struct.
// Does not write it to file.
func (c *ConfigFile) Add(h HostEntry) error {

	// ensure hostentry has identifier fields stripped of leading and trailing whitespace
	h.HostName = strings.TrimPrefix(h.HostName, " ")
	h.HostName = strings.TrimSuffix(h.HostName, " ")
	h.Host = strings.TrimPrefix(h.Host, " ")
	h.Host = strings.TrimSuffix(h.Host, " ")

	// validate hostentry
	if err := h.Valid(); err != nil {
		return err
	}

	// ensure it doesnt already exist
	for _, v := range c.Hosts {
		if (v.Host == h.Host) && (v.HostName == h.HostName) {
			return errors.New("error: conflicting host found")
		}
	}

	c.Hosts = append(c.Hosts, h)

	return nil
}

// Removes the entry with matching host from the config file.
func (c *ConfigFile) Rm(host string) error {

	// loop and filter the host out
	newhosts := make([]HostEntry, 0)
	for _, v := range c.Hosts {
		if v.Host == host {
			continue
		}

		newhosts = append(newhosts, v)
	}

	if len(newhosts) == len(c.Hosts) {
		return errors.New("host: '" + host + "' not found")
	}

	// update hosts
	c.Hosts = newhosts

	return nil
}
