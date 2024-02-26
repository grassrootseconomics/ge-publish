package command

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	toolsPath      = ".getools"
	privateKeyPath = ".ge_pk"
)

func (c *Command) RegisterWritePrivateKeyCommand(logg *slog.Logger) *cli.Command {
	return &cli.Command{
		Name:  "save-private-key",
		Usage: "Save a private key that can be autoloaded by ge-publish",
		Action: func(cCtx *cli.Context) error {
			return nil
		},
	}
}

func (c *Command) PrivateKeyFileLocation(logg *slog.Logger) string {
	location, err := os.UserHomeDir()
	if err != nil {
		logg.Error("homedir not found", "error", err)
		os.Exit(1)
	}

	return fmt.Sprintf("%s/%s/%s", location, toolsPath, privateKeyPath)
}
