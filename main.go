package main

import (
	"os"

	"github.com/alecthomas/kingpin"
)

func main() {
	kingpin.Version("0.0.1")

	app := kingpin.New("Flexvolume - RioFS", "Provides RioFS S3 support for Kubernetes")

	Init(app)
	Attach(app)
	Detach(app)
	Mount(app)
	Unmount(app)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
