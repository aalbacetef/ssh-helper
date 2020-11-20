package configfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// generate a mock configfile struct using a /tmp/ based path

func mockconfigfile() (ConfigFile, error) {
	cfg := ConfigFile{}

	// create a temporary file
	tmpdir := os.TempDir()
	tmpfile, err := ioutil.TempFile(tmpdir, "ssh-helper-tests-")
	if err != nil {
		return cfg, fmt.Errorf("error: could not create temporary file: %v", err)
	}

	// write an empty string to it
	if _, err := tmpfile.Write([]byte("")); err != nil {
		return cfg, fmt.Errorf("error: could not write to tmp file: %v", err)
	}

	// save file path to cfg struct
	st, err := tmpfile.Stat()
	if err != nil {
		return cfg, fmt.Errorf("error: could not stat tmp file: %v", err)
	}
	cfg.Fpath = Fpath(filepath.Join(tmpdir, st.Name()))
	cfg.Name = "mock-file"
	cfg.Hosts = []HostEntry{}
	return cfg, nil
}

func TestConfigFile(t *testing.T) {

	t.Run("can-add-host", func(tt *testing.T) {
		cfg, err := mockconfigfile()
		if err != nil {
			t.Fatal("could not create mock config file: ", err)
		}

		mockhost := HostEntry{
			Host:     "local-test",
			HostName: "local.test",
		}
		if err := cfg.Add(mockhost); err != nil {
			tt.Fatal("could not add mockhost: ", err)
		}

		const expectedhosts = 1
		if len(cfg.Hosts) != expectedhosts {
			tt.Fatalf("expected %d hosts, got %d", expectedhosts, len(cfg.Hosts))
		}

		if cfg.Hosts[0].Host != mockhost.Host {
			tt.Fatalf("expected host %s, got %s", cfg.Hosts[0].Host, mockhost.Host)
		}

		if cfg.Hosts[0].HostName != mockhost.HostName {
			tt.Fatalf("expected hostname %s, got %s", cfg.Hosts[0].HostName, mockhost.HostName)
		}
	})

	t.Run("can-rm-host", func(tt *testing.T) {
		cfg, err := mockconfigfile()
		if err != nil {
			t.Fatal("could not create mock config file: ", err)
		}

		mockhost := HostEntry{
			Host:     "local-test",
			HostName: "local.test",
		}
		// @NOTE previous test would've caught errors
		_ = cfg.Add(mockhost)
		if err := cfg.Rm(mockhost.Host); err != nil {
			tt.Fatalf("could not rm host: %v", err)
		}

		if len(cfg.Hosts) != 0 {
			t.Log(cfg.Hosts)
			tt.Fatalf("expected empty hosts list, got %d", len(cfg.Hosts))
		}
	})
}
