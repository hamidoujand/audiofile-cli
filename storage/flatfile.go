package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/hamidoujand/audiofile/models"

	"github.com/google/uuid"
)

type FlatFile struct {
	Name string
}

func (f FlatFile) GetByID(id string) (*models.Audio, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	metadataFilePath := filepath.Join(dirname, "audiofile", id, "metadata.json")
	if _, err := os.Stat(metadataFilePath); errors.Is(err, os.ErrNotExist) {
		_ = os.Mkdir(metadataFilePath, os.ModePerm)
	}
	file, err := os.ReadFile(metadataFilePath)
	if err != nil {
		return nil, err
	}
	data := models.Audio{}
	err = json.Unmarshal([]byte(file), &data)
	return &data, err
}

func (f FlatFile) SaveMetadata(audio *models.Audio) error {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	audioDirPath := filepath.Join(dirname, "audiofile", audio.Id)
	metadataFilePath := filepath.Join(audioDirPath, "metadata.json")
	file, err := os.Create(metadataFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := audio.JSON()
	if err != nil {
		fmt.Println("Err: ", err)
		return err
	}
	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func (f FlatFile) Upload(bytes []byte, filename string) (string, string, error) {
	// generate guid
	id := uuid.New()
	// copy file to configured storage path by tag name or id
	dirname, err := os.UserHomeDir()
	if err != nil {
		return id.String(), "", err
	}
	audioDirPath := filepath.Join(dirname, "audiofile", id.String())
	if err := os.MkdirAll(audioDirPath, os.ModePerm); err != nil {
		return id.String(), "", err
	}
	audioFilePath := filepath.Join(audioDirPath, filename)
	err = os.WriteFile(audioFilePath, bytes, 0644)
	if err != nil {
		return id.String(), "", err
	}
	return id.String(), audioFilePath, nil
}

func (f FlatFile) List() ([]*models.Audio, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	metadataFilePath := filepath.Join(dirname, "audiofile")
	if _, err := os.Stat(metadataFilePath); errors.Is(err, os.ErrNotExist) {
		_ = os.Mkdir(metadataFilePath, os.ModePerm)
	}
	files, err := os.ReadDir(metadataFilePath)
	if err != nil {
		return nil, err
	}
	audioFiles := []*models.Audio{}
	for _, file := range files {
		if file.IsDir() {
			name, err := f.GetByID(file.Name())
			if err != nil {
				return nil, err
			}
			audioFiles = append(audioFiles, name)
		}
	}
	return audioFiles, nil
}

func (f FlatFile) Delete(id string) error {
	fmt.Println("Deleting")
	return nil
}
