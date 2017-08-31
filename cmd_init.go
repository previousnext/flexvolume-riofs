package main

import (
	"encoding/json"
	"fmt"

	"github.com/alecthomas/kingpin"
)

type InitCommand struct{}

func Init(app *kingpin.Application) {
	c := &InitCommand{}
	app.Command("init", "Initialize the driver").Action(c.Run)
}

func (c *InitCommand) Run(k *kingpin.ParseContext) error {
	b, err := json.Marshal(map[string]string{
		"status": "Success",
	})
	if err != nil {
		fatal(err)
	}
	fmt.Println(string(b))

	return nil
}
