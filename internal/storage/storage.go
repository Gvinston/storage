package storage

import (
	"fmt"
	"github.com/Gvinston/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	Files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		Files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	file, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	s.Files[file.ID] = file

	return file, nil
}

func (s *Storage) GetById(fileID uuid.UUID) (*file.File, error) {
	f, ok := s.Files[fileID]
	if ok != true {
		return nil, fmt.Errorf("file %v not found", fileID)
	}

	return f, nil
}
