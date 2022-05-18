package ewords

import (
	"errors"
	"fmt"
	"log"
	"path"
	"path/filepath"

	// "strconv"
	"strings"

	"github.com/spf13/afero"
)

const (
	CanCreate = iota
	DirExists
	Error
)

func IsDirExists(dir *string, fs afero.Fs) (bool, error) {
	return afero.DirExists(fs, *dir)
}

func DirCanBeCreated(dir *string, fs afero.Fs) (int, error) {
	rp, err := RealPath(*dir)
	if err != nil {
		return Error, err
	}

	exists, err := IsDirExists(&rp, fs)
	if err != nil {
		return Error, err
	}
	if exists {
		return DirExists, nil
	}

	deleteTarget, err := FindFirstUnexisted(rp, fs)
	if err != nil {
		return Error, err
	}

	err = fs.MkdirAll(rp, 0777)
	if err != nil {
		return Error, err
	}
	err = fs.RemoveAll(deleteTarget)
	if err != nil {
		return Error, err
	}
	return CanCreate, nil
}

func FindFirstUnexisted(dir string, fs afero.Fs) (string, error) {
	fromroot := strings.HasPrefix(dir, "/")
	if !fromroot {
		return "", fmt.Errorf("path: %v should be abs", dir)
	}
	if dir == "/" {
		return "", fmt.Errorf("path: %v should not be passed here", dir)
	}
	var deleteTarget string
	curr := dir
	stop := false
	for !stop {
		parent := path.Dir(curr)
		ex, err := IsDirExists(&parent, fs)
		if err != nil {
			return "", err
		}
		if ex {
			deleteTarget = curr
			stop = true
		} else {
			curr = parent
		}
	}
	return deleteTarget, nil
}

func RealPath(path string) (string, error) {
	cleaned := filepath.Clean(path)
	if filepath.IsAbs(cleaned) {
		return cleaned, nil
	}
	abs, err := filepath.Abs(cleaned)
	if err != nil {
		return "", err
	}
	return abs, nil
}

func ReadDir(dirPath *string, fs afero.Fs) (map[string]string, error) {
	afs := &afero.Afero{Fs: fs}
	entries, err := afs.ReadDir(*dirPath)
	if err != nil {
		return nil, fmt.Errorf("error while reading dir: %v - %v", *dirPath, err)
	}
	res := map[string]string{}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		filepath := path.Join(*dirPath, entry.Name())
		data, err := ReadFile(&filepath, fs)
		if err != nil {
			return nil, fmt.Errorf("error while proccessing file: %v in the directory: %v - %v", entry, dirPath, err)
		}
		res[entry.Name()] = data
	}
	return res, nil
}

func ListFiles(dirPath *string, fs afero.Fs) ([]string, error) {
	afs := &afero.Afero{Fs: fs}
	entries, err := afs.ReadDir(*dirPath)
	if err != nil {
		return nil, fmt.Errorf("error while reading dir: %v - %v", *dirPath, err)
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

func ReadFile(path *string, fs afero.Fs) (string, error) {
	afs := &afero.Afero{Fs: fs}
	b := strings.Builder{}
	data, err := afs.ReadFile(*path)
	if err != nil {
		et := fmt.Sprintf("Error while reading file %v - %v", *path, err)
		log.Print(et)
		return "", fmt.Errorf(et)
	}
	b.Write(data)
	return b.String(), nil
}

func WriteFile(fpath string, data *string, fs afero.Fs) error {
	afs := &afero.Afero{Fs: fs}
	_, err := afs.Stat(fpath)
	if errors.Is(err, afero.ErrFileExists) {
		et := fmt.Sprintf("File %v - exists", fpath)
		log.Print(et)
		return fmt.Errorf(et)
	}

	dir := path.Dir(fpath)
	de, err := afs.DirExists(dir)
	if err != nil {
		return fmt.Errorf("Error while checking dir: %v existing - %v", dir, err)
	}
	if !de {
		afs.MkdirAll(dir, 0777)
	}

	err = afs.WriteFile(fpath, []byte(*data), 0666)
	if err != nil {
		et := fmt.Sprintf("Error while writing to the file %v - %v", fpath, err)
		log.Print(et)
		return fmt.Errorf(et)
	}
	return nil
}
