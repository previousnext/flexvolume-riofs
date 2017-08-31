package main

import (
	"encoding/json"
	"fmt"

	"github.com/alecthomas/kingpin"
)

type DetachCommand struct {
	Device string
}

func Detach(app *kingpin.Application) {
	c := &DetachCommand{}
	cmd := app.Command("detach", "Attach the volume to the host").Action(c.Run)
	cmd.Arg("device", "Device to detach").Required().StringVar(&c.Device)
}

func (c *DetachCommand) Run(k *kingpin.ParseContext) error {
	b, err := json.Marshal(map[string]string{
		"status": "Success",
	})
	if err != nil {
		fatal(err)
	}
	fmt.Println(string(b))

	return nil
}
