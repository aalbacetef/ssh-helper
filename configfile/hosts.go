package configfile

import "errors"

// HostEntry represents an individual record in the ssh
// config file. It requires the Host and HostName fields
// to be defined, but allows the User and IdentityFilePath
// fields to be left empty.
type HostEntry struct {
	Host             string `json:"host"`
	HostName         string `json:"hostname"`
	IdentityFilePath Fpath  `json:"identity-file-path"`
	User             string `json:"user"`
}

// Performs basic validation on the HostEntry provided.
// Requires that both Host and HostName be defined.
func (h HostEntry) Valid() error {
	if h.Host == "" {
		return errors.New("error: host is empty")
	}
	if h.HostName == "" {
		return errors.New("errors: hostname is empty")
	}
	return nil
}
