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

func TestDirExists(t *testing.T) {
	memFs := afero.NewMemMapFs()

	memFs.Mkdir("/dir", 0777)

	res, err := ewords.IsDirExists(utils.StrPtr("/dir"), memFs)
	utils.AssertNil(t, err)
	utils.AssertEqualsBool(t, true, res)
}

func TestDirNotExists(t *testing.T) {
	memFs := afero.NewMemMapFs()

	res, err := ewords.IsDirExists(utils.StrPtr("/dir"), memFs)
	utils.AssertNil(t, err)
	utils.AssertEqualsBool(t, false, res)
}

func TestFindFirstUnexisted(t *testing.T) {

	memFs := afero.NewMemMapFs()

	memFs.MkdirAll("/data", 0777)
	memFs.Mkdir("/", 0777)

	p1 := "/"
	p2 := "/data/dir1"
	p3 := "/data/dir1/dir2"
	p4 := "."

	res, err := ewords.FindFirstUnexisted(p1, memFs)
	utils.AssertNotNill(t, err)
	utils.AssertEqualsString(t, "path: / should not be passed here", err.Error())
	utils.AssertEmptyString(t, res)

	res, err = ewords.FindFirstUnexisted(p2, memFs)
	utils.AssertNil(t, err)
	utils.AssertEqualsString(t, "/data/dir1", res)

	res, err = ewords.FindFirstUnexisted(p3, memFs)
	utils.AssertNil(t, err)
	utils.AssertEqualsString(t, "/data/dir1", res)

	_, err = ewords.FindFirstUnexisted(p4, memFs)
	utils.AssertNotNill(t, err)
	utils.AssertEqualsString(t, "path: . should be abs", err.Error())
}

func TestDirCanBeCreatedDirExists(t *testing.T) {
	memFs := afero.NewMemMapFs()

	dir1 := "/dir"
	memFs.Mkdir(dir1, 0777)

	res, err := ewords.DirCanBeCreated(&dir1, memFs)
	utils.AssertNil(t, err)
	utils.AssertEqualsInt(t, ewords.DirExists, res)

	dir2 := "/data/dir2"
	dir2Unnorm := "/data/dir/../.././data/dir2"
	memFs.MkdirAll(dir2, 0777)
	res, err = ewords.DirCanBeCreated(&dir2Unnorm, memFs)
	utils.AssertNil(t, err)
	utils.AssertEqualsInt(t, ewords.DirExists, res)

	dir3 := "/"
	res, err = ewords.DirCanBeCreated(&dir3, memFs)
	utils.AssertNil(t, err)
	utils.AssertEqualsInt(t, ewords.DirExists, res)
}

func TestDictCanBeCreatedDirNotExists(t *testing.T) {
	memFs := afero.NewMemMapFs()
	aMemFs := &afero.Afero{Fs: memFs}
	root := "/"
	memFs.Mkdir(root, 0777)

	dir1 := "/data/dir"
	res, err := ewords.DirCanBeCreated(&dir1, memFs)
	utils.AssertNil(t, err)
	utils.AssertEqualsInt(t, ewords.CanCreate, res)
	ex, err := aMemFs.DirExists(dir1)
	utils.AssertNil(t, err)
	utils.AssertEqualsBool(t, false, ex)

	memFs = afero.NewMemMapFs()
	aMemFs = &afero.Afero{Fs: memFs}
	root2 := "/volume"
	memFs.MkdirAll(root2, 0777)
	dir2 := "/volume/vm1/dir"
	res, err = ewords.DirCanBeCreated(&dir2, memFs)
	utils.AssertNil(t, err)
	utils.AssertEqualsInt(t, ewords.CanCreate, res)
	ex, err = aMemFs.DirExists(root2)
	utils.AssertNil(t, err)
	utils.AssertEqualsBool(t, true, ex)
	ex, err = aMemFs.DirExists(dir2)
	utils.AssertNil(t, err)
	utils.AssertEqualsBool(t, false, ex)
	ex, err = aMemFs.DirExists(path.Dir(dir2))
	utils.AssertNil(t, err)
	utils.AssertEqualsBool(t, false, ex)
}
