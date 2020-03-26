package configfile

import (
	"strings"
)

func trim(old, val string) string {
	s := strings.Replace(old, val, "", -1)
	s = strings.TrimPrefix(s, " ")
	s = strings.TrimSuffix(s, " ")
	return s
}

func linestartswith(l string) func(p string) bool {
	return func(p string) bool {
		return strings.HasPrefix(l, p)
	}
}

// Parse an SSH config file
func Parse(data string) (*ConfigFile, error) {

	// will be used to read all the hosts
	hosts := make([]HostEntry, 0)
	currEntry := &HostEntry{}

	lines := strings.Split(data, "\n")
	for _, line := range lines {

		// @NOTE: should probably use a regex for whitespace
		line := strings.TrimLeft(line, " \t")
		startswith := linestartswith(line)

		// choose which property to set or if block has been
		// finished parsing
		// @NOTE: consider tokenizing
		if startswith("Host ") {
			currEntry.Host = trim(line, "Host ")
		} else if startswith("User ") {
			currEntry.User = trim(line, "User ")
		} else if startswith("HostName ") {
			currEntry.HostName = trim(line, "HostName ")
		} else if startswith("IdentityFile ") {
			currEntry.IdentityFilePath = Fpath(trim(line, "IdentityFile "))
		} else if line == "" {
			hosts = append(hosts, *currEntry)
			currEntry = &HostEntry{}
		}
	}

	// filter out invalidly parsed hosts
	filteredhosts := make([]HostEntry, 0)
	for _, host := range hosts {
		if err := host.Valid(); err != nil {
			continue
		}

		filteredhosts = append(filteredhosts, host)
	}

	return &ConfigFile{Hosts: filteredhosts}, nil
}
