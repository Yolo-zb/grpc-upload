package server

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
)

type ImageStore struct {
	mutex sync.RWMutex
}

func (store *ImageStore) Save(FileData bytes.Buffer, FilePath string, FileExt string) (string, error) {
	imageID, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("cannot generate image id: %w", err)
	}

	imagePath := fmt.Sprintf("%s%s%s", FilePath, imageID, FileExt)

	file, err := os.Create(imagePath)
	if err != nil {
		return "", fmt.Errorf("cannot create image file: %w", err)
	}

	_, err = FileData.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("cannot write image to file: %w", err)
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()

	return imagePath, nil
}
