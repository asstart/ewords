//go:build !integration
// +build !integration

package ewords_test

import (
	"fmt"
	"io"
	"path"
	"testing"

	"github.com/asstart/english-words/app/ewords"
	"github.com/asstart/english-words/app/utils"

	"github.com/spf13/afero"
)

func TestReadEmptyDir(t *testing.T) {
	var memFs afero.Fs = afero.NewMemMapFs()
	dp := "dir"

	memFs.Mkdir(dp, 0777)

	res, err := ewords.ReadDir(&dp, memFs)

	utils.AssertNil(t, err)
	utils.AssertEqualsInt(t, 0, len(res))
}

func TestReadExistedDir(t *testing.T) {
	memFs := afero.NewMemMapFs()

	dp := "dir"

	fn1 := "file1"
	fc1 := "File 1 content"
	fn2 := "file2"
	fc2 := "File 2 content"

	memFs.Mkdir(dp, 0777)
	f1, _ := memFs.Create(path.Join(dp, fn1))
	f1.WriteString(fc1)
	f2, _ := memFs.Create(path.Join(dp, fn2))
	f2.WriteString(fc2)

	res, err := ewords.ReadDir(&dp, memFs)
	utils.AssertNil(t, err)
	utils.AssertEqualsInt(t, 2, len(res))

	v, ok := res[fn1]

	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, fc1, v)

	v, ok = res[fn2]

	utils.AssertEqualsBool(t, true, ok)
	utils.AssertEqualsString(t, fc2, v)
}

func TestReadNotDir(t *testing.T) {
	memFs := afero.NewMemMapFs()

	path := "some_file"

	memFs.Create(path)

	res, err := ewords.ReadDir(&path, memFs)

	utils.AssertNotNill(t, err)
	utils.AssertNil(t, res)

	experr := fmt.Sprintf("error while reading dir: %s - readdir %v: not a dir", path, path)

	utils.AssertEqualsString(t, experr, err.Error())
}

func TestListFilesEmptyDir(t *testing.T) {
	memFs := afero.NewMemMapFs()

	dp := "dir"

	memFs.Mkdir(dp, 0777)

	res, err := ewords.ListFiles(&dp, memFs)

	utils.AssertNil(t, err)
	utils.AssertNotNill(t, res)
	utils.AssertEqualsInt(t, 0, len(res))
}

func TestListFiles(t *testing.T) {
	memFs := afero.NewMemMapFs()

	dp := "dir"
	f1 := path.Join(dp, "file1")
	f2 := path.Join(dp, "file2")

	memFs.Mkdir(dp, 0777)
	memFs.Create(f1)
	memFs.Create(f2)

	res, err := ewords.ListFiles(&dp, memFs)
	utils.AssertNil(t, err)
	utils.AssertNotNill(t, res)
	utils.AssertEqualsInt(t, 2, len(res))
}

func TestListFilesNotDir(t *testing.T) {
	memFs := afero.NewMemMapFs()

	path := "some_file"

	memFs.Create(path)

	res, err := ewords.ListFiles(&path, memFs)

	utils.AssertNotNill(t, err)
	utils.AssertNil(t, res)

	experr := fmt.Sprintf("error while reading dir: %s - readdir %v: not a dir", path, path)

	utils.AssertEqualsString(t, experr, err.Error())
}

func TestReadFile(t *testing.T) {
	memFs := afero.NewMemMapFs()

	fp := "some_file"
	fc := "file content"

	file, _ := memFs.Create(fp)
	file.WriteString(fc)

	res, err := ewords.ReadFile(&fp, memFs)
	utils.AssertNil(t, err)
	utils.AssertEqualsString(t, fc, res)
}

func TestReadEmptyFile(t *testing.T) {
	memFs := afero.NewMemMapFs()

	fp := "some_file"
	fc := ""

	file, _ := memFs.Create(fp)
	file.WriteString(fc)

	res, err := ewords.ReadFile(&fp, memFs)
	utils.AssertNil(t, err)
	utils.AssertEqualsString(t, fc, res)
}

func TestReadUnexistedFile(t *testing.T) {
	memFs := afero.NewMemMapFs()

	fp := "some_file"

	res, err := ewords.ReadFile(&fp, memFs)
	utils.AssertNotNill(t, err)
	utils.AssertEmptyString(t, res)

	errtext := fmt.Sprintf("Error while reading file %v - open %v: file does not exist", fp, fp)

	utils.AssertEqualsString(t, errtext, err.Error())
}

func TestWriteFile(t *testing.T) {
	memFs := afero.NewMemMapFs()

	fp := "some_file"
	fc := "file content"

	err := ewords.WriteFile(fp, &fc, memFs)

	utils.AssertNil(t, err)

	f, err := memFs.Open(fp)
	utils.AssertNil(t, err)
	defer f.Close()

	var res string
	var buffer []byte = make([]byte, 1)
	for {
		_, err = f.Read(buffer)
		if err == io.EOF {
			break
		}
		res += string(buffer)
	}

	utils.AssertEqualsString(t, fc, res)
}

func TestWriteFileToNotExistedDir(t *testing.T) {
	memFs := afero.NewMemMapFs()

	fp := path.Join("some_dir", "some_file")
	fc := "file content"

	err := ewords.WriteFile(fp, &fc, memFs)

	utils.AssertNil(t, err)

	f, err := memFs.Open(fp)

	utils.AssertNil(t, err)
	defer f.Close()

	var res string
	var buffer []byte = make([]byte, 1)
	for {
		_, err = f.Read(buffer)
		if err == io.EOF {
			break
		}
		res += string(buffer)
	}

	utils.AssertEqualsString(t, fc, res)
}
