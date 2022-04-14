package ewords

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func ReadDir(dirPath *string) (map[string]string, error) {
	entries, err := os.ReadDir(*dirPath)
	if err != nil {
		return nil, fmt.Errorf("error while reading dir: %v - %v", dirPath, err)
	}
	res := map[string]string{}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		filepath := path.Join(*dirPath, entry.Name())
		data, err := ReadFile(&filepath)
		if err != nil {
			return nil, fmt.Errorf("error while proccessing file: %v in the directory: %v - %v", entry, dirPath, err)
		}
		res[entry.Name()] = data
	}
	return res, nil
}

func ListFiles(dirPath *string) ([]string, error) {
	entries, err := os.ReadDir(*dirPath)
	if err != nil {
		return nil, fmt.Errorf("error while reading dir: %v - %v", dirPath, err)
	}
	res := []string{}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		res = append(res, path.Join(*dirPath, entry.Name()))
	}
	return res, nil
}

func ReadFile(path *string) (string, error) {
	b := strings.Builder{}
	data, err := os.ReadFile(*path)
	if err != nil {
		et := fmt.Sprintf("Error while reading file: %v", err)
		log.Print(et)
		return "", fmt.Errorf(et)
	}
	b.Write(data)
	return b.String(), nil
}

func WriteFile(path string, data *string) error {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		et := fmt.Sprintf("File %v - exists", path)
		log.Print(et)
		return fmt.Errorf(et)
	}
	err = os.WriteFile(path, []byte(*data), 0666)
	if err != nil {
		et := fmt.Sprintf("Error while writing to the file %v - %v", path, err)
		log.Print(et)
		return fmt.Errorf(et)
	}
	return nil
}
