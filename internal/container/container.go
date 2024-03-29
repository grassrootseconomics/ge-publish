package container

import (
	"log/slog"

	"github.com/kamikazechaser/common/logg"
)

type (
	Container struct {
		Logg *slog.Logger
	}
)

func NewContainer() *Container {
	return &Container{
		Logg: logg.NewLogg(logg.LoggOpts{
			FormatType: logg.Human,
			LogLevel:   slog.LevelInfo,
		}),
	}
}

func (c *Container) UseDebugMode() {
	c.Logg = logg.NewLogg(logg.LoggOpts{
		FormatType: logg.Human,
		LogLevel:   slog.LevelDebug,
	})
}
