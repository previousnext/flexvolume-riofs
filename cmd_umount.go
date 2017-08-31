package main

import (
	"encoding/json"
	"fmt"

	"github.com/alecthomas/kingpin"
)

type UnmountCommand struct {
	CacheDir string
	Mount    string
}

func Unmount(app *kingpin.Application) {
	c := &UnmountCommand{}
	cmd := app.Command("unmount", "Unmounts a directory").Action(c.Run)
	cmd.Flag("cachedir", "The directory to store cache").Default("/var/cache/riofs").StringVar(&c.CacheDir)
	cmd.Arg("mount", "The directory to unmount").Required().StringVar(&c.Mount)
}

func (c *UnmountCommand) Run(k *kingpin.ParseContext) error {
	// Unmount the filesystem.
	err := shellOut([]string{
		"fusermount",
		"-u",
		c.Mount,
	})
	if err != nil {
		fatal(err)
	}

	// Perform cleanup of the cache directory.
	shellOut([]string{
		"rm",
		"-fR",
		fmt.Sprintf("%s/%s", c.CacheDir, hash(c.Mount)),
	})

	b, err := json.Marshal(map[string]string{
		"status": "Success",
	})
	if err != nil {
		fatal(err)
	}
	fmt.Println(string(b))

	return nil
}
