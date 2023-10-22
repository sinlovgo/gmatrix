package subcommand_new

import (
	"github.com/bar-counter/slog"
	"github.com/sinlovgo/gmatrix/command"
	"github.com/sinlovgo/gmatrix/constant"
	"github.com/sinlovgo/gmatrix/internal/urfave_cli"
	"github.com/urfave/cli/v2"
)

const commandName = "new"

var commandEntry *NewCommand

type NewCommand struct {
	isDebug bool

	PlatformConfig *constant.PlatformConfig
}

func (n *NewCommand) Exec() error {
	slog.Debugf("-> Exec subCommand [ %s ]", commandName)

	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "lib",
			Usage: "Use a library template",
			Value: false,
		},
		&cli.StringFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "Set the resulting package name, defaults to the directory name",
		},
	}
}

func withEntry(c *cli.Context) (*NewCommand, error) {
	slog.Debugf("-> withEntry subCommand [ %s ]", commandName)

	if c.Bool("lib") {
		slog.Info("new lib mode")
	}
	globalEntry := command.CmdGlobalEntry()
	return &NewCommand{
		isDebug: globalEntry.Verbose,

		// todo: if not use platform config, remove this
		PlatformConfig: constant.BindPlatformConfig(c),
	}, nil
}

func action(c *cli.Context) error {
	slog.Debugf("-> Sub Command action [ %s ] start", commandName)
	entry, err := withEntry(c)
	if err != nil {
		return err
	}
	commandEntry = entry
	return commandEntry.Exec()
}

func Command() []*cli.Command {
	urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())
	return []*cli.Command{
		{
			Name:   commandName,
			Usage:  "",
			Action: action,

			// todo: if not use platform config, remove this
			//Flags: flag(),
			Flags: urfave_cli.UrfaveCliAppendCliFlag(flag(), constant.PlatformFlags()),
		},
	}
}
