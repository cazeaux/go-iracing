package store

import (
	"context"
	"log/slog"
	"time"

	"github.com/cazeaux/go-iracing/cmd/config"
)

type Data struct {
	LastResultTimestamp time.Time `json:"lastResultTimestamp"`
	LastSubSessionID    int       `json:"lastSubSessionID"`
}

type Message struct {
	Data *Data
	User *config.User
}

type DataStoreReaderWriter interface {
	Writer(u *config.User, data *Data) error
	Reader(u *config.User) (*Data, error)
}

type DataStore struct {
	WriterChannel chan Message
	Store         DataStoreReaderWriter
	Logger        *slog.Logger
}

func NewDataStore(l *slog.Logger, w chan Message, s DataStoreReaderWriter) *DataStore {
	return &DataStore{
		WriterChannel: w,
		Store:         s,
		Logger:        l,
	}
}

func WriterRoutine(ctx context.Context, d *DataStore) {
	for {
		select {
		case <-ctx.Done():
			return
		case message := <-d.WriterChannel:
			d.Logger.Info("writing data cache", "user", message.User.ID)
			err := d.Store.Writer(message.User, message.Data)
			if err != nil {
				d.Logger.Error("error on writer for user", "user", message.User.ID, "error", err)
			}
		}
	}
}

func Get(d *DataStore, u *config.User) (*Data, error) {
	data, err := d.Store.Reader(u)
	if err != nil {
		return nil, err
	}

	if data == nil {
		d.Logger.Info("user data cache does not exit", "user", u.ID)
		data = &Data{
			LastResultTimestamp: time.Now(),
			LastSubSessionID:    0,
		}
	}

	return data, nil
}
