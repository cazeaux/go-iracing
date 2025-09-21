package notification

import (
	"github.com/cazeaux/go-iracing/cmd/config"
	"github.com/cazeaux/go-iracing/cmd/results"
)

type Notifier interface {
	Notify(r *results.Results, u *config.User) error
}
