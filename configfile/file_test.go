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
	cfg, err := mockconfigfile()
	if err != nil {
		t.Fatal("could not create mock config file: ", err)
	}

	t.Log("name: ", cfg.Name)
	t.Log("fpath: ", cfg.Fpath)
}

// test the Add function

// test the Rm function
