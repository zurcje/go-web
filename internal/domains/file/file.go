package file

import (
	"encoding/json"
	"log"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
	AddMock(mock *Mock)
	ClearMock()
}

type Type string

const (
	FileType Type = "file"
)

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{
			FileName: fileName,
		}
	}
	return nil
}

type Mock struct {
	Data []byte
	Err  error
}

type FileStore struct {
	FileName string
	Mock     *Mock
}

func (fs *FileStore) AddMock(mock *Mock) {
	fs.Mock = mock
}

func (fs *FileStore) ClearMock() {
	fs.Mock = nil
}

func (fs *FileStore) Write(data interface{}) error {
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		return json.Unmarshal(fs.Mock.Data, data)
	}
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("failed to write. err:", err)
		return err
	}
	return os.WriteFile(fs.FileName, fileData, 0644)
}

func (fs *FileStore) Read(data interface{}) error {
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		return json.Unmarshal(fs.Mock.Data, data)
	}
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		log.Println("failed to read. err:", err)
		return err
	}
	return json.Unmarshal(file, &data)
}
