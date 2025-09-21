package config

import "log/slog"

func Logger() *slog.Logger {
	return slog.Default()
}
