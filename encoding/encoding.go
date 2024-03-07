package encoding

import (
	"encoding/json"
	"fmt"
	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
	"os"
)

// JSONData тип для перекодирования из JSON в YAML
// JSONData type for converting from JSON to YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
// YAMLData type for converting from YAML to JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
// MyEncoder interface for YAMLData and JSONData structures
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
// Encoding converts a file from JSON to YAML
func (j *JSONData) Encoding() error {
	jsonFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonFile, &j.DockerCompose); err != nil {
		return err
	}
	yamlData, err := yaml.Marshal(&j.DockerCompose)
	f, err := os.Create(j.FileOutput)
	if err != nil {
		fmt.Printf("ошибка при создании файла: %s", err.Error())
		return err
	}
	defer f.Close()

	_, err = f.Write(yamlData)
	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s", err.Error())
		return err
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
// Encoding converts a file from YAML to JSON
func (y *YAMLData) Encoding() error {
	yamlFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(yamlFile, &y.DockerCompose); err != nil {
		return err
	}
	yamlData, err := json.Marshal(&y.DockerCompose)
	f, err := os.Create(y.FileOutput)
	if err != nil {
		fmt.Printf("ошибка при создании файла: %s", err.Error())
		return err
	}
	defer f.Close()

	_, err = f.Write(yamlData)
	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s", err.Error())
		return err
	}
	return nil
}
