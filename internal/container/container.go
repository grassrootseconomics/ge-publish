package container

import (
	"log/slog"

	"github.com/kamikazechaser/common/logg"
)

type (
	LoggOpts struct {
		JSON  bool
		Debug bool
	}
	Container struct {
		Logg *slog.Logger
	}
)

var loggOpts = logg.LoggOpts{
	FormatType: logg.Human,
	LogLevel:   slog.LevelInfo,
}

func NewContainer() *Container {
	return &Container{
		Logg: logg.NewLogg(loggOpts),
	}
}

func (c *Container) OverrideLogger(o LoggOpts) {
	if o.JSON {
		loggOpts.FormatType = logg.JSON
	}

	if o.Debug {
		loggOpts.LogLevel = slog.LevelDebug
	}

	c.Logg = logg.NewLogg(loggOpts)
}
