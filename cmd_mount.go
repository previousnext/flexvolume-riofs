package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
)

type MountCommand struct {
	Target   string
	Uid      string
	Gid      string
	Fmode    string
	Dmode    string
	CacheDir string
	LogFile  string
	Device   string
	Options  string
}

func Mount(app *kingpin.Application) {
	c := &MountCommand{}
	cmd := app.Command("mount", "Mounts a device onto into a directory").Action(c.Run)
	cmd.Flag("uid", "The uid of the mount").Default("33").StringVar(&c.Uid)
	cmd.Flag("gid", "The uid of the mount").Default("33").StringVar(&c.Gid)
	cmd.Flag("fmode", "The fmode of the file in the mount").Default("511").StringVar(&c.Fmode)
	cmd.Flag("dmode", "The dmode of the directory in the mount").Default("511").StringVar(&c.Dmode)
	cmd.Flag("logfile", "The file to log to").Default("/var/log/syslog").StringVar(&c.LogFile)
	cmd.Flag("cachedir", "The directory to store cache").Default("/var/cache/riofs").StringVar(&c.CacheDir)
	cmd.Arg("target", "The target directory").Required().StringVar(&c.Target)
	cmd.Arg("device", "The device to mount").Required().StringVar(&c.Device)
	cmd.Arg("options", "Raw JSON set of options").Required().StringVar(&c.Options)
}

func (c *MountCommand) Run(k *kingpin.ParseContext) error {
	// Before we perform the mount we need to make sure the directory exists.
	err := os.MkdirAll(c.Target, 0755)
	if err != nil {
		fatal(err)
	}

	// For now we defer to RioFS. Long term we should look at leveraging
	// https://github.com/kahing/goofys and calling it directly via Golang APIs.
	err = shellOut([]string{
		"riofs",
		"-o allow_other",
		"--log-file=" + c.LogFile,
		// We declare a unique cache for each RioFS mount. The easiest way to do
		// this is toget a hash of the "destination" directory, that way we can
		// do cleanup afterwards.
		"--cache-dir=" + fmt.Sprintf("%s/%s", c.CacheDir, hash(c.Target)),
		"--uid=" + c.Uid,
		"--gid=" + c.Gid,
		"--fmode=" + c.Fmode,
		"--dmode=" + c.Dmode,
		c.Device,
		c.Target,
	})
	if err != nil {
		fatal(err)
	}

	b, err := json.Marshal(map[string]string{
		"status": "Success",
	})
	if err != nil {
		fatal(err)
	}
	fmt.Println(string(b))

	return nil
}
