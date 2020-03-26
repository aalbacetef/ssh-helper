package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"gitlab.com/aalbacetef/ssh-helper/utils"
)

const DefaultBackupDir = "~/.ssh-backups"

const BackupMsg = `
Usage: ssh-helper backup [--path] --tag

Options:
  --path    Use this path as the backup directory. If it does not exist
            it will be created.
  --tag     A tag to use to name the backup.
`

func BackupUsage() {
	fmt.Println(BackupMsg)
}

func Backup(tag, bpath string) {
	basepath := path.Clean(DefaultBackupDir)
	if bpath != "" {
		basepath = path.Clean(bpath)
	}

	// @NOTE : there are probably better naming schemes
	fname := fmt.Sprint(tag, "-", time.Now().Unix(), ".tar")
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("error getting home directory: ", err)
		return
	}

	// @NOTE : this only matters on linux?
	basepath = strings.Replace(
		basepath,
		"~",
		homedir,
		1,
	)

	// ensure path exists
	err = utils.MkDirP(basepath)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	// @NOTE : this will only run on Linux...
	cmd := exec.Command(
		"tar",
		"-cf",
		path.Join(basepath, fname),
		path.Join(homedir, ".ssh"),
	)

	err = cmd.Run()
	if err != nil {
		fmt.Println("error tar-ing file: ", err)
		return
	}

	fmt.Println("Done.")
}
