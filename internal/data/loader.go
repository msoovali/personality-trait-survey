package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/msoovali/personality-trait-survey/internal/domain"
)

const (
	conf_dir  = "conf"
	data_json = "/data.json"
)

type Data struct {
	Traits    []domain.Trait    `json:"traits"`
	Questions []domain.Question `json:"questions"`
}

func LoadData() *Data {
	filePath := conf_dir + data_json
	byteValue := loadFileAndReadContent(filePath)
	if byteValue == nil {
		return nil
	}
	var data Data
	err := json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Printf("Couldn't unmarshal file (%s) contents: %v\n", filePath, err)
		return nil
	}

	return &data
}

func loadFileAndReadContent(filePath string) []byte {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil
	}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Couldn't open file(%s)! %v\n", filePath, err)
		return nil
	}
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file(%s) content! %v\n", filePath, err)
		return nil
	}

	return byteValue
}
