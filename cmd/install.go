package cmd

import (
	"sort"

	"github.com/LongMarch7/gfile/cmd/bench"
	"github.com/LongMarch7/gfile/cmd/receive"
	"github.com/LongMarch7/gfile/cmd/send"
	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
)

// Install all the commands
func Install(app *cli.App) {
	app.Commands = []cli.Command{
		send.New(),
		receive.New(),
		bench.New(),
	}
	log.Trace("Installed commands")

	sort.Sort(cli.CommandsByName(app.Commands))
}
