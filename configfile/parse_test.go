package configfile

import "testing"

// @TODO add test for filtering out invalid hosts
func TestParse(t *testing.T) {
	const mockhostfile = `
		Host local-test
	    HostName 192.168.0.1
	    User test-user
	    IdentityFile ~/.ssh/local-test/local-test
 
	`

	mockhostentry := HostEntry{
		Host:             "local-test",
		HostName:         "192.168.0.1",
		User:             "test-user",
		IdentityFilePath: "~/.ssh/local-test/local-test",
	}

	t.Run("parse-host-file", func(tt *testing.T) {
		cfg, err := Parse(mockhostfile)
		if err != nil {
			tt.Fatalf("error whilst parsing mock host file: %v", err)
		}

		const expected = 1
		if len(cfg.Hosts) != expected {
			tt.Fatalf("expected %d host(s), found %d", expected, len(cfg.Hosts))
		}

		if cfg.Hosts[0].Host != mockhostentry.Host {
			tt.Fatalf("expected host %s, got %s", mockhostentry.Host, cfg.Hosts[0].Host)
		}

		if cfg.Hosts[0].HostName != mockhostentry.HostName {
			tt.Fatalf("expected hostname %s, got %s", mockhostentry.HostName, cfg.Hosts[0].HostName)
		}

		if cfg.Hosts[0].User != mockhostentry.User {
			tt.Fatalf("expected user %s, got %s", mockhostentry.User, cfg.Hosts[0].User)
		}
		if cfg.Hosts[0].IdentityFilePath != mockhostentry.IdentityFilePath {
			tt.Fatalf("expected identity file path %s, got %s", mockhostentry.IdentityFilePath, cfg.Hosts[0].IdentityFilePath)
		}
	})

}
