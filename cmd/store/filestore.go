package store

import (
	"encoding/json"
	"os"
	"path"

	"github.com/cazeaux/go-iracing/cmd/config"
)

type FileStore struct {
	Path string
}

func (f *FileStore) Writer(u *config.User, data *Data) error {

	filename := path.Join(f.Path, u.ID)

	fd, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()

	text, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = fd.Write(text)
	if err != nil {
		return err
	}

	return nil
}

func (f *FileStore) Reader(u *config.User) (*Data, error) {

	filename := path.Join(f.Path, u.ID)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, nil
	}
	
	text, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data Data

	err = json.Unmarshal(text, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
